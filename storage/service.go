package storage

import (
	"errors"

	"github.com/gonum/stat"
	"github.com/rodrigoapereira/auditmq/config"
	"github.com/spf13/cast"
)

type ServiceStatus int

const (
	Unknown ServiceStatus = iota + 1
	Sync
	NotSync
)

type Service interface {
	Name() string
	Status() ServiceStatus
	SetStatus(status ServiceStatus)
	HasField(fieldName string) bool
	GetField(fieldName string) (Field, error)
	AddField(fieldName string, opts config.ServiceData, owner Service) error
	AddFieldWithOptions(fieldName string, opts config.ServiceData, owner Service, options config.DataOptions) error
	Push(fieldName string, data interface{}) error
	DuplicateLast(fieldName string) error
	Fields() Fields
	FieldsToMatch() Fields
}

type service struct {
	name   string
	fields Fields
	status ServiceStatus
}

type Fields map[string]Field

type Field struct {
	Name        string
	Type        string
	Buffer      *Buffer
	DataOptions config.DataOptions
	OwnedBy     Service
}

func (f Field) PushNumber(data interface{}) error {
	value, err := cast.ToFloat64E(data)
	if err != nil {
		return errors.New("Cast error")
	}

	value += f.DataOptions.Offset

	f.Buffer.Push(value)
	return nil
}

func (f Field) DuplicateLast() error {
	f.Buffer.CopyLastValue()
	return nil
}

func (f Field) ToSlice() []interface{} {
	return f.Buffer.ToSlice()
}

func (f Field) convertToFloat64(sample []interface{}) ([]float64, error) {
	result := []float64{}
	for _, value := range sample {
		newValue, err := cast.ToFloat64E(value)
		if err != nil {
			return result, err
		}

		result = append(result, newValue)
	}

	return result, nil
}

func (f Field) equalsFloat64(expectedData, comparedData []interface{}) (bool, error) {
	expected, err := f.convertToFloat64(expectedData)
	if err != nil {
		return false, err
	}

	compared, err := f.convertToFloat64(comparedData)
	if err != nil {
		return false, err
	}

	expectedVar := stat.Variance(expected, nil)
	if expectedVar != 0 {
		return false, errors.New("Sample is not stable")
	}

	comparedVar := stat.Variance(compared, nil)
	if comparedVar != 0 {
		return false, errors.New("Compared sample is not stable")
	}

	return stat.Mean(expected, nil) == stat.Mean(compared, nil), nil
}

func (f Field) Equals(otherField Field) (bool, error) {
	actualData := f.Buffer.ToSlice()
	comparedData := otherField.ToSlice()
	if len(actualData) != len(comparedData) {
		return false, errors.New("Sample size not match")
	}

	switch f.Type {
	case "int":
		fallthrough
	case "float":
		return f.equalsFloat64(actualData, comparedData)
	default:
		return false, errors.New("Unsupported type")
	}
}

func (s *service) Name() string {
	return s.name
}

func (s *service) Status() ServiceStatus {
	return s.status
}

func (s *service) SetStatus(status ServiceStatus) {
	s.status = status
}

func (s *service) HasField(fieldName string) bool {
	_, ok := s.fields[fieldName]
	return ok
}

func (s *service) GetField(fieldName string) (Field, error) {
	if s.HasField(fieldName) {
		return s.fields[fieldName], nil
	}

	return Field{}, errors.New("Field not configured")
}

func (s *service) AddField(fieldName string, opts config.ServiceData, owner Service) error {
	if s.HasField(fieldName) {
		return errors.New("Field already configured")
	}

	s.fields[fieldName] = Field{
		Name:    fieldName,
		Type:    opts.Type,
		Buffer:  NewBuffer(opts.Samples),
		OwnedBy: owner,
	}

	return nil
}

func (s *service) AddFieldWithOptions(fieldName string, opts config.ServiceData, owner Service, options config.DataOptions) error {
	err := s.AddField(fieldName, opts, owner)
	if err != nil {
		return err
	}

	field := s.fields[fieldName]
	field.DataOptions = options
	s.fields[fieldName] = field

	return nil
}

func (s *service) Push(fieldName string, data interface{}) error {
	field, err := s.GetField(fieldName)
	if err != nil {
		return err
	}

	switch field.Type {
	case "int":
		fallthrough
	case "float":
		return field.PushNumber(data)
	default:
		return errors.New("Unknown field type")
	}
}

func (s *service) DuplicateLast(fieldName string) error {
	field, err := s.GetField(fieldName)
	if err != nil {
		return err
	}

	return field.DuplicateLast()
}

func (s *service) FieldsToMatch() Fields {
	matchedFields := Fields{}
	for fieldName, field := range s.fields {
		if field.OwnedBy != nil {
			matchedFields[fieldName] = field
		}
	}

	return matchedFields
}

func (s *service) Fields() Fields {
	return s.fields
}

func NewService(name string) Service {
	return &service{
		name:   name,
		fields: make(Fields),
		status: Unknown,
	}
}
