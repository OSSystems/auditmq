package storage

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type StorageTestSuite struct {
	suite.Suite
}

func (s *StorageTestSuite) SetupTest() {
	new()
}

func (s *StorageTestSuite) TestGetStorageData() {
	s.Equal(0, len(storage.services))

	storageData := GetStorageData("api")
	s.NotNil(storageData)

	s.Equal(1, len(storage.services))
}

func (s *StorageTestSuite) TestGetBuffer() {
	storageData := GetStorageData("api")
	s.Equal(0, len(storageData))

	buff := storageData.GetBuffer("device-count")
	s.NotNil(buff)
	s.Equal(1, len(storageData))

	_, ok := storageData["device-count"]
	s.True(ok)
}

func TestStorage(t *testing.T) {
	suite.Run(t, &StorageTestSuite{})
}
