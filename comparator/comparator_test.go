package comparator

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rodrigoapereira/auditmq/config"
	"github.com/rodrigoapereira/auditmq/mocks"
	"github.com/rodrigoapereira/auditmq/storage"
	"github.com/stretchr/testify/suite"
)

type ComparatorTestSuite struct {
	suite.Suite
	mockCtrl    *gomock.Controller
	storageMock *mocks.MockStorage
	comparator  *comparator
}

func (s *ComparatorTestSuite) SetupTest() {
	s.mockCtrl = gomock.NewController(s.T())
	s.storageMock = mocks.NewMockStorage(s.mockCtrl)
	s.comparator = &comparator{
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

	s.Equal(ownerService.Status(), storage.Unknown)
	s.Equal(childService.Status(), storage.Unknown)

	s.comparator.compare()

	s.Equal(ownerService.Status(), storage.Unknown)
	s.Equal(childService.Status(), storage.NotSync)
}

func TestComparator(t *testing.T) {
	suite.Run(t, new(ComparatorTestSuite))
}