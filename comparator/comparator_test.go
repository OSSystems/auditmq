package comparator

import (
	"errors"
	"testing"

	"github.com/OSSystems/auditmq/config"
	"github.com/OSSystems/auditmq/mocks"
	"github.com/OSSystems/auditmq/storage"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type ComparatorTestSuite struct {
	suite.Suite
	mockCtrl      *gomock.Controller
	storageMock   *mocks.MockStorage
	publisherMock *mocks.MockPublisher
	comparator    *comparator
}

func (s *ComparatorTestSuite) SetupTest() {
	s.mockCtrl = gomock.NewController(s.T())
	s.storageMock = mocks.NewMockStorage(s.mockCtrl)
	s.publisherMock = mocks.NewMockPublisher(s.mockCtrl)
	s.comparator = &comparator{
		pub:     s.publisherMock,
		storage: s.storageMock,
		stop:    make(chan bool),
		done:    make(chan error),
	}
}

func (s *ComparatorTestSuite) TestCompareWithOwnerFieldNotPresent() {
	ownerService := storage.NewService("owner-service")
	childService := storage.NewService("child-service")
	childService.AddField("field1", config.ServiceData{
		Samples: 3,
		Type:    "int",
	}, ownerService)
	s.storageMock.EXPECT().Services().Return(storage.Services{
		"owner-service": ownerService,
		"child-service": childService,
	})

	ownerService.Push("field1", 10)
	ownerService.Push("field1", 10)
	ownerService.Push("field1", 10)

	childService.Push("field1", 10)
	childService.Push("field1", 10)
	childService.Push("field1", 10)

	s.Equal(ownerService.Status(), storage.Unknown)
	s.Equal(childService.Status(), storage.Unknown)

	s.comparator.compare()

	s.Equal(ownerService.Status(), storage.Unknown)
	s.Equal(childService.Status(), storage.Unknown)
}

func (s *ComparatorTestSuite) TestCompareWithDifferentSampleSize() {
	ownerService := storage.NewService("owner-service")
	ownerService.AddField("field1", config.ServiceData{
		Samples: 3,
		Type:    "int",
	}, nil)
	childService := storage.NewService("child-service")
	childService.AddField("field1", config.ServiceData{
		Samples: 3,
		Type:    "int",
	}, ownerService)
	s.storageMock.EXPECT().Services().Return(storage.Services{
		"owner-service": ownerService,
		"child-service": childService,
	})

	ownerService.Push("field1", 10)
	ownerService.Push("field1", 10)
	ownerService.Push("field1", 10)

	childService.Push("field1", 10)
	childService.Push("field1", 10)

	s.Equal(ownerService.Status(), storage.Unknown)
	s.Equal(childService.Status(), storage.Unknown)

	s.comparator.compare()

	s.Equal(ownerService.Status(), storage.Unknown)
	s.Equal(childService.Status(), storage.Unknown)
}

func (s *ComparatorTestSuite) TestCompareWithManyFields() {
	ownerService := storage.NewService("owner-service")
	ownerService.AddField("field1", config.ServiceData{
		Samples: 3,
		Type:    "int",
	}, nil)
	ownerService.AddField("field2", config.ServiceData{
		Samples: 3,
		Type:    "int",
	}, nil)
	childService := storage.NewService("child-service")
	childService.AddField("field1", config.ServiceData{
		Samples: 3,
		Type:    "int",
	}, ownerService)
	childService.AddField("field2", config.ServiceData{
		Samples: 3,
		Type:    "int",
	}, ownerService)
	s.storageMock.EXPECT().Services().Return(storage.Services{
		"owner-service": ownerService,
		"child-service": childService,
	})

	ownerService.Push("field1", 10)
	ownerService.Push("field1", 10)
	ownerService.Push("field1", 10)
	ownerService.Push("field2", 10)
	ownerService.Push("field2", 10)
	ownerService.Push("field2", 10)

	childService.Push("field1", 10)
	childService.Push("field1", 10)
	childService.Push("field1", 10)
	childService.Push("field2", 30)
	childService.Push("field2", 30)
	childService.Push("field2", 30)

	s.publisherMock.EXPECT().PublishServiceStatus(childService).Return(nil)
	s.publisherMock.EXPECT().PublishGlobalStat().Return(nil)
	s.Equal(ownerService.Status(), storage.Unknown)
	s.Equal(childService.Status(), storage.Unknown)

	s.comparator.compare()

	s.Equal(ownerService.Status(), storage.Unknown)
	s.Equal(childService.Status(), storage.NotSync)
}

func (s *ComparatorTestSuite) TestCompareWithEqualValues() {
	ownerService := storage.NewService("owner-service")
	ownerService.AddField("field1", config.ServiceData{
		Samples: 3,
		Type:    "int",
	}, nil)
	childService := storage.NewService("child-service")
	childService.AddField("field1", config.ServiceData{
		Samples: 3,
		Type:    "int",
	}, ownerService)
	s.storageMock.EXPECT().Services().Return(storage.Services{
		"owner-service": ownerService,
		"child-service": childService,
	})

	ownerService.Push("field1", 10)
	ownerService.Push("field1", 10)
	ownerService.Push("field1", 10)

	childService.Push("field1", 10)
	childService.Push("field1", 10)
	childService.Push("field1", 10)

	s.publisherMock.EXPECT().PublishServiceStatus(childService).Return(nil)
	s.publisherMock.EXPECT().PublishGlobalStat().Return(nil)
	s.Equal(ownerService.Status(), storage.Unknown)
	s.Equal(childService.Status(), storage.Unknown)

	s.comparator.compare()

	s.Equal(ownerService.Status(), storage.Unknown)
	s.Equal(childService.Status(), storage.Sync)
}

func (s *ComparatorTestSuite) TestCompareWithUnequalValues() {
	ownerService := storage.NewService("owner-service")
	ownerService.AddField("field1", config.ServiceData{
		Samples: 3,
		Type:    "int",
	}, nil)
	childService := storage.NewService("child-service")
	childService.AddField("field1", config.ServiceData{
		Samples: 3,
		Type:    "int",
	}, ownerService)
	s.storageMock.EXPECT().Services().Return(storage.Services{
		"owner-service": ownerService,
		"child-service": childService,
	})

	ownerService.Push("field1", 10)
	ownerService.Push("field1", 10)
	ownerService.Push("field1", 10)

	childService.Push("field1", 11)
	childService.Push("field1", 11)
	childService.Push("field1", 11)

	s.publisherMock.EXPECT().PublishServiceStatus(childService).Return(nil)
	s.publisherMock.EXPECT().PublishGlobalStat().Return(nil)

	s.Equal(ownerService.Status(), storage.Unknown)
	s.Equal(childService.Status(), storage.Unknown)

	s.comparator.compare()

	s.Equal(ownerService.Status(), storage.Unknown)
	s.Equal(childService.Status(), storage.NotSync)
}

func (s *ComparatorTestSuite) TestCompareWithPublishServiceError() {
	ownerService := storage.NewService("owner-service")
	ownerService.AddField("field1", config.ServiceData{
		Samples: 3,
		Type:    "int",
	}, nil)
	childService := storage.NewService("child-service")
	childService.AddField("field1", config.ServiceData{
		Samples: 3,
		Type:    "int",
	}, ownerService)
	s.storageMock.EXPECT().Services().Return(storage.Services{
		"owner-service": ownerService,
		"child-service": childService,
	})

	ownerService.Push("field1", 10)
	ownerService.Push("field1", 10)
	ownerService.Push("field1", 10)

	childService.Push("field1", 11)
	childService.Push("field1", 11)
	childService.Push("field1", 11)

	err := errors.New("Publish error")
	s.publisherMock.EXPECT().PublishServiceStatus(childService).Return(err)

	returnedErr := s.comparator.compare()
	if s.Error(returnedErr) {
		s.Equal(err, returnedErr)
	}
}

func (s *ComparatorTestSuite) TestCompareWithPublishGlobalStatusError() {
	ownerService := storage.NewService("owner-service")
	ownerService.AddField("field1", config.ServiceData{
		Samples: 3,
		Type:    "int",
	}, nil)
	childService := storage.NewService("child-service")
	childService.AddField("field1", config.ServiceData{
		Samples: 3,
		Type:    "int",
	}, ownerService)
	s.storageMock.EXPECT().Services().Return(storage.Services{
		"owner-service": ownerService,
		"child-service": childService,
	})

	ownerService.Push("field1", 10)
	ownerService.Push("field1", 10)
	ownerService.Push("field1", 10)

	childService.Push("field1", 11)
	childService.Push("field1", 11)
	childService.Push("field1", 11)

	err := errors.New("Publish error")
	s.publisherMock.EXPECT().PublishServiceStatus(childService).Return(nil)
	s.publisherMock.EXPECT().PublishGlobalStat().Return(err)

	returnedErr := s.comparator.compare()
	if s.Error(returnedErr) {
		s.Equal(err, returnedErr)
	}
}

func TestComparator(t *testing.T) {
	suite.Run(t, new(ComparatorTestSuite))
}
