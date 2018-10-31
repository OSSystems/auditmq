package storage

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type StorageTestSuite struct {
	suite.Suite
	storage *Storage
}

func (s *StorageTestSuite) SetupTest() {
	storageInstance = newStorage()
}

func (s *StorageTestSuite) TestGetStorage() {
	storageInstance = nil
	localStorage, err := GetStorage()
	if s.Error(err) {
		s.Nil(localStorage)
	}

	storageInstance = newStorage()
	localStorage, err = GetStorage()
	if s.NoError(err) {
		s.NotNil(localStorage)
	}
}

func (s *StorageTestSuite) TestGetService() {
	_, err := storageInstance.GetService("test")
	if s.Error(err) {
		s.Equal("Service not found", err.Error())
	}
}

func TestStorage(t *testing.T) {
	suite.Run(t, new(StorageTestSuite))
}
