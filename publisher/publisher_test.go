package publisher

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/NeowayLabs/wabbit"
	"github.com/NeowayLabs/wabbit/amqptest"
	"github.com/NeowayLabs/wabbit/amqptest/server"
	"github.com/OSSystems/auditmq/mocks"
	"github.com/OSSystems/auditmq/pkg"
	"github.com/OSSystems/auditmq/storage"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type PublisherTestSuite struct {
	suite.Suite
	mockCtrl    *gomock.Controller
	storageMock *mocks.MockStorage
	ch          *mocks.MockChannel
	exchange    string
	routingKey  string
	publisher   *publisher
}

func (s *PublisherTestSuite) SetupTest() {
	s.mockCtrl = gomock.NewController(s.T())
	s.storageMock = mocks.NewMockStorage(s.mockCtrl)

	s.ch = mocks.NewMockChannel(s.mockCtrl)
	s.exchange = "test_exchange"
	s.routingKey = "auditmq"
	s.publisher = &publisher{
		conn:       nil,
		ch:         s.ch,
		storage:    s.storageMock,
		exchange:   s.exchange,
		routingKey: s.routingKey,
	}
}

func (s *PublisherTestSuite) TestGetChannel() {
	fakeServer := server.NewServer("amqp://localhost:5672/%2f")
	fakeServer.Start()

	mockConn, err := amqptest.Dial("amqp://localhost:5672/%2f")
	if err != nil {
		s.FailNow(err.Error())
	}

	s.Nil(s.publisher.conn)
	s.NotNil(s.publisher.ch)
	s.publisher.ch = nil
	s.publisher.conn = mockConn

	ch, err := s.publisher.getChannel()
	s.NoError(err)
	s.NotNil(ch)

	s.NotNil(s.publisher.ch)
}

func (s *PublisherTestSuite) TestPublishServiceStatus() {
	service := mocks.NewMockService(s.mockCtrl)
	service.EXPECT().Name().Return("server")
	service.EXPECT().Status().Return(storage.Sync)

	bytes, _ := json.Marshal(pkg.ReportPayload{
		Service: "server",
		Status:  int(storage.Sync),
	})

	s.ch.EXPECT().Publish(s.exchange, s.routingKey, bytes, wabbit.Option{
		"contentType": "application/json",
	}).Return(nil)
	s.NoError(s.publisher.PublishServiceStatus(service))
}

func (s *PublisherTestSuite) TestPublishServiceStatusWithError() {
	service := mocks.NewMockService(s.mockCtrl)
	service.EXPECT().Name().Return("server")
	service.EXPECT().Status().Return(storage.NotSync)

	bytes, _ := json.Marshal(pkg.ReportPayload{
		Service: "server",
		Status:  int(storage.NotSync),
	})

	err := errors.New("Publish error")
	s.ch.EXPECT().Publish(s.exchange, s.routingKey, bytes, wabbit.Option{
		"contentType": "application/json",
	}).Return(err)

	returnedErr := s.publisher.PublishServiceStatus(service)
	if s.Error(returnedErr) {
		s.Equal(err, returnedErr)
	}
}

func (s *PublisherTestSuite) TestPublishGlobalStat() {
	firstService := mocks.NewMockService(s.mockCtrl)
	firstService.EXPECT().Status().Return(storage.Sync)
	secondService := mocks.NewMockService(s.mockCtrl)
	secondService.EXPECT().Status().Return(storage.NotSync)

	s.storageMock.EXPECT().Services().Return(storage.Services{
		"firstService":  firstService,
		"secondService": secondService,
	}).Times(2)

	bytes, _ := json.Marshal(pkg.ReportPayload{
		Service: "all",
		Status:  int(storage.NotSync),
	})

	s.ch.EXPECT().Publish(s.exchange, s.routingKey, bytes, wabbit.Option{
		"contentType": "application/json",
	}).Return(nil)
	s.NoError(s.publisher.PublishGlobalStat())

	firstService.EXPECT().Status().Return(storage.Sync)
	secondService.EXPECT().Status().Return(storage.Sync)

	bytes, _ = json.Marshal(pkg.ReportPayload{
		Service: "all",
		Status:  int(storage.Sync),
	})

	s.ch.EXPECT().Publish(s.exchange, s.routingKey, bytes, wabbit.Option{
		"contentType": "application/json",
	}).Return(nil)
	s.NoError(s.publisher.PublishGlobalStat())
}

func (s *PublisherTestSuite) TestPublishGlobalStatWithError() {
	firstService := mocks.NewMockService(s.mockCtrl)
	firstService.EXPECT().Status().Return(storage.Sync)

	s.storageMock.EXPECT().Services().Return(storage.Services{
		"firstService": firstService,
	})

	bytes, _ := json.Marshal(pkg.ReportPayload{
		Service: "all",
		Status:  int(storage.Sync),
	})

	err := errors.New("Publish error")
	s.ch.EXPECT().Publish(s.exchange, s.routingKey, bytes, wabbit.Option{
		"contentType": "application/json",
	}).Return(err)

	returnedErr := s.publisher.PublishGlobalStat()
	if s.Error(returnedErr) {
		s.Equal(err, returnedErr)
	}
}

func TestPublisher(t *testing.T) {
	suite.Run(t, new(PublisherTestSuite))
}
