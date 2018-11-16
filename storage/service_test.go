package storage

import (
	"testing"

	"github.com/OSSystems/auditmq/config"
	"github.com/stretchr/testify/suite"
)

type FieldTestSuite struct {
	suite.Suite
}

func (s *FieldTestSuite) NewTestField() Field {
	return Field{
		Name:    "fieldOne",
		Type:    "int",
		Buffer:  NewBuffer(3),
		OwnedBy: nil,
	}
}

func (s *FieldTestSuite) TestPushNumber() {
	field := s.NewTestField()
	err := field.PushNumber("error")
	if s.Error(err) {
		s.Equal("Cast error", err.Error())
	}
}

func (s *FieldTestSuite) TestEquals() {
	fieldOne := s.NewTestField()
	fieldTwo := s.NewTestField()

	fieldOne.PushNumber(1)
	result, err := fieldOne.Equals(fieldTwo)
	s.False(result)
	if s.Error(err) {
		s.Equal("Sample size not match", err.Error())
	}

	fieldTwo.PushNumber(1)
	result, err = fieldOne.Equals(fieldTwo)
	s.False(result)
	if s.Error(err) {
		s.Equal("Sample is not stable", err.Error())
	}

	fieldOne.PushNumber(1)
	fieldOne.PushNumber(1)

	fieldTwo.PushNumber(2)
	fieldTwo.PushNumber(1)
	result, err = fieldOne.Equals(fieldTwo)
	s.False(result)
	if s.Error(err) {
		s.Equal("Compared sample is not stable", err.Error())
	}

	fieldTwo.PushNumber(1)
	fieldTwo.PushNumber(1)

	result, err = fieldOne.Equals(fieldTwo)
	s.True(result)
	s.NoError(err)
}

func TestField(t *testing.T) {
	suite.Run(t, new(FieldTestSuite))
}

type ServiceTestSuite struct {
	suite.Suite
	service Service
}

func (s *ServiceTestSuite) SetupTest() {
	s.service = NewService("server1")
}

func (s *ServiceTestSuite) TestGetField() {
	field, err := s.service.GetField("fieldOne")
	s.Empty(field)
	if s.Error(err) {
		s.Equal("Field not configured", err.Error())
	}

	err = s.service.AddField("fieldOne", config.ServiceData{
		Type:    "int",
		Owner:   "server1",
		Samples: 3,
	}, nil)
	s.NoError(err)

	field, err = s.service.GetField("fieldOne")
	s.NoError(err)
	s.NotEmpty(field)
}

func (s *ServiceTestSuite) TestPush() {
	err := s.service.Push("fieldOne", 1)
	if s.Error(err) {
		s.Equal("Field not configured", err.Error())
	}

	err = s.service.AddField("fieldOne", config.ServiceData{
		Type:    "intx",
		Owner:   "server1",
		Samples: 3,
	}, nil)
	s.NoError(err)

	err = s.service.Push("fieldOne", 1)
	s.Error(err)

	field := s.service.Fields()["fieldOne"]
	field.Type = "int"
	s.service.Fields()["fieldOne"] = field

	err = s.service.Push("fieldOne", 1)
	s.NoError(err)
}

func (s *ServiceTestSuite) TestDuplicateLast() {
	err := s.service.AddField("fieldOne", config.ServiceData{
		Type:    "int",
		Owner:   "server1",
		Samples: 3,
	}, nil)
	s.NoError(err)

	s.service.DuplicateLast("fieldOne")
	field, err := s.service.GetField("fieldOne")
	s.NoError(err)
	s.Equal([]interface{}{}, field.ToSlice())

	err = s.service.Push("fieldOne", 1)
	s.NoError(err)

	s.service.DuplicateLast("fieldOne")

	field, err = s.service.GetField("fieldOne")
	s.NoError(err)
	s.EqualValues([]interface{}{float64(1), float64(1)}, field.ToSlice())
}

func (s *ServiceTestSuite) TestAddField() {
	err := s.service.AddField("fieldOne", config.ServiceData{
		Type:    "int",
		Owner:   "server1",
		Samples: 3,
	}, nil)
	s.NoError(err)

	err = s.service.AddField("fieldOne", config.ServiceData{
		Type:    "int",
		Owner:   "server1",
		Samples: 3,
	}, nil)
	if s.Error(err) {
		s.Equal("Field already configured", err.Error())
	}
}

func TestService(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
