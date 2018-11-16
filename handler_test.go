package main

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/OSSystems/auditmq/config"
	"github.com/OSSystems/auditmq/mocks"
	"github.com/OSSystems/auditmq/pkg"
	"github.com/OSSystems/auditmq/storage"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
	mockCtrl    *gomock.Controller
	message     *mocks.MockDelivery
	storageMock *mocks.MockStorage
	handler     *Handler
}

func (s *HandlerTestSuite) SetupTest() {
	nackRetryTime = 0

	s.mockCtrl = gomock.NewController(s.T())
	s.message = mocks.NewMockDelivery(s.mockCtrl)
	s.storageMock = mocks.NewMockStorage(s.mockCtrl)
	s.handler = &Handler{
		storage: s.storageMock,
	}
}

func (s *HandlerTestSuite) prepareMessage() {
	messageBytes, err := json.Marshal(pkg.Payload{
		Service: "api-server",
		Data: map[string]interface{}{
			"value1": 10,
		},
	})

	if err != nil {
		s.FailNow("Prepare message error!")
	}

	s.message.EXPECT().Body().Return(messageBytes)
}

func (s *HandlerTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

func (s *HandlerTestSuite) TestHandlerWithInvalidMessage() {
	s.message.EXPECT().Body().Return([]byte("{Invalid}"))
	s.message.EXPECT().Nack(false, true)
	s.handler.Handle(s.message)
}

func (s *HandlerTestSuite) TestHandlerWithoutServices() {
	s.prepareMessage()
	s.message.EXPECT().Nack(false, true)
	s.storageMock.EXPECT().GetService("api-server").Return(nil, errors.New("Unknown error"))
	s.handler.Handle(s.message)
}

func (s *HandlerTestSuite) TestHandlerWithServices() {
	s.prepareMessage()
	s.message.EXPECT().Ack(false)

	invalidReplica := storage.NewService("old-server")

	validService := storage.NewService("api-server")
	validService.AddField("value1", config.ServiceData{
		Type:    "int",
		Owner:   "",
		Samples: 5,
	}, nil)

	validReplica := storage.NewService("valid-replica")
	validReplica.AddField("value1", config.ServiceData{
		Type:    "int",
		Owner:   "",
		Samples: 5,
	}, validService)
	validReplica.Push("value1", 10)

	s.storageMock.EXPECT().GetService("api-server").Return(validService, nil).Times(2)
	s.storageMock.EXPECT().Services().Return(storage.Services{
		"api-server":    validService,
		"old-server":    invalidReplica,
		"valid-replica": validReplica,
	})
	s.handler.Handle(s.message)

	validServiceField, err := validService.GetField("value1")
	s.NoError(err)
	s.Equal([]interface{}{
		float64(10),
	}, validServiceField.ToSlice())

	validReplicaField, err := validReplica.GetField("value1")
	s.NoError(err)
	s.Equal([]interface{}{
		float64(10),
		float64(10),
	}, validReplicaField.ToSlice())
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
