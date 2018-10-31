package storage

import (
	"testing"

	"github.com/rodrigoapereira/auditmq/config"
	"github.com/stretchr/testify/suite"
)

type BuilderTestSuite struct {
	suite.Suite
	storage Storage
}

var oneFieldConfig = config.DataFields{
	"fieldOne": config.ServiceData{
		Type:    "int",
		Owner:   "server1",
		Samples: 10,
		Replicas: map[string]config.DataOptions{
			"server2": config.DataOptions{},
		},
	},
}

var twoFieldsConfig = config.DataFields{
	"fieldOne": config.ServiceData{
		Type:    "int",
		Owner:   "server1",
		Samples: 10,
		Replicas: map[string]config.DataOptions{
			"server2": config.DataOptions{
				Offset: 10,
			},
		},
	},
	"fieldTwo": config.ServiceData{
		Type:    "int",
		Owner:   "server1",
		Samples: 3,
		Replicas: map[string]config.DataOptions{
			"server2": config.DataOptions{},
		},
	},
}

var complexConfig = config.DataFields{
	"fieldOne": config.ServiceData{
		Type:    "int",
		Owner:   "server1",
		Samples: 10,
		Replicas: map[string]config.DataOptions{
			"server2": config.DataOptions{
				Offset: 10,
			},
		},
	},
	"fieldTwo": config.ServiceData{
		Type:    "int",
		Owner:   "server2",
		Samples: 10,
		Replicas: map[string]config.DataOptions{
			"server1": config.DataOptions{},
		},
	},
}

func (s *BuilderTestSuite) SetupTest() {
	s.storage = newStorage()
}

func (s *BuilderTestSuite) TestBuildWithOneField() {
	builder := &StorageBuilder{
		storage:    s.storage,
		baseConfig: oneFieldConfig,
	}
	s.NoError(builder.Build())

	server1, err := s.storage.GetService("server1")
	s.NoError(err)
	s.Len(server1.FieldsToMatch(), 0)
	ownField := server1.Fields()["fieldOne"]
	s.Equal("int", ownField.Type)
	s.Equal(10, ownField.Buffer.Len())
	s.Nil(ownField.OwnedBy)
	s.Equal(config.DataOptions{}, ownField.DataOptions)

	server2, err := s.storage.GetService("server2")
	s.NoError(err)
	s.Len(server2.FieldsToMatch(), 1)
	childField := server2.Fields()["fieldOne"]
	s.Equal("int", childField.Type)
	s.Equal(10, childField.Buffer.Len())
	s.Equal(server1, childField.OwnedBy)
	s.Equal(config.DataOptions{}, childField.DataOptions)
}

func (s *BuilderTestSuite) TestBuildWithManyFields() {
	builder := &StorageBuilder{
		storage:    s.storage,
		baseConfig: twoFieldsConfig,
	}
	s.NoError(builder.Build())

	server1, err := s.storage.GetService("server1")
	s.NoError(err)
	s.Len(server1.FieldsToMatch(), 0)
	ownFirstField := server1.Fields()["fieldOne"]
	s.Equal("int", ownFirstField.Type)
	s.Equal(10, ownFirstField.Buffer.Len())
	s.Nil(ownFirstField.OwnedBy)
	s.Equal(config.DataOptions{}, ownFirstField.DataOptions)

	ownSecondField := server1.Fields()["fieldTwo"]
	s.Equal("int", ownSecondField.Type)
	s.Equal(3, ownSecondField.Buffer.Len())
	s.Nil(ownSecondField.OwnedBy)
	s.Equal(config.DataOptions{}, ownSecondField.DataOptions)

	server2, err := s.storage.GetService("server2")
	s.NoError(err)
	s.Len(server2.FieldsToMatch(), 2)
	firstChildField := server2.Fields()["fieldOne"]
	s.Equal("int", firstChildField.Type)
	s.Equal(10, firstChildField.Buffer.Len())
	s.Equal(server1, firstChildField.OwnedBy)
	s.Equal(config.DataOptions{
		Offset: 10,
	}, firstChildField.DataOptions)

	secondChildField := server2.Fields()["fieldTwo"]
	s.Equal("int", secondChildField.Type)
	s.Equal(3, secondChildField.Buffer.Len())
	s.Equal(server1, secondChildField.OwnedBy)
	s.Equal(config.DataOptions{}, secondChildField.DataOptions)
}

func (s *BuilderTestSuite) TestBuildWithComplexFields() {
	builder := &StorageBuilder{
		storage:    s.storage,
		baseConfig: complexConfig,
	}
	s.NoError(builder.Build())

	server1, err := s.storage.GetService("server1")
	s.NoError(err)
	s.Len(server1.FieldsToMatch(), 1)

	server2, err := s.storage.GetService("server2")
	s.NoError(err)
	s.Len(server2.FieldsToMatch(), 1)

	ownFirstField := server1.Fields()["fieldOne"]
	s.Equal("int", ownFirstField.Type)
	s.Equal(10, ownFirstField.Buffer.Len())
	s.Nil(ownFirstField.OwnedBy)
	s.Equal(config.DataOptions{}, ownFirstField.DataOptions)

	childSecondField := server1.Fields()["fieldTwo"]
	s.Equal("int", childSecondField.Type)
	s.Equal(10, childSecondField.Buffer.Len())
	s.Equal(server2, childSecondField.OwnedBy)
	s.Equal(config.DataOptions{}, childSecondField.DataOptions)

	ownSecondField := server2.Fields()["fieldTwo"]
	s.Equal("int", ownSecondField.Type)
	s.Equal(10, ownSecondField.Buffer.Len())
	s.Nil(ownSecondField.OwnedBy)
	s.Equal(config.DataOptions{}, ownSecondField.DataOptions)

	childFirstField := server2.Fields()["fieldOne"]
	s.Equal("int", childFirstField.Type)
	s.Equal(10, childFirstField.Buffer.Len())
	s.Equal(server1, childFirstField.OwnedBy)
	s.Equal(config.DataOptions{
		Offset: 10,
	}, childFirstField.DataOptions)
}

func (s *BuilderTestSuite) TestInitializeStorage() {
	s.Nil(storageInstance)
	InitializeStorage(complexConfig)
	s.NotNil(storageInstance)
}

func TestBuilder(t *testing.T) {
	suite.Run(t, new(BuilderTestSuite))
}
