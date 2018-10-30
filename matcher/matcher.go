package matcher

import (
	"errors"

	"github.com/gonum/stat"
	"github.com/rodrigoapereira/auditmq/config"
)

type Data struct {
	Type    string
	OwnedBy string
	Options config.DataOptions
}

func (d *Data) convertToFloat64(sample []interface{}) (result []float64, err error) {
	for _, value := range sample {
		newValue, ok := value.(float64)
		if !ok {
			err = errors.New("Conversion error")
			return
		}

		result = append(result, newValue)
	}

	return
}

func (d *Data) compareFloat64(expectedData, actualData []interface{}) (bool, error) {
	expected, err := d.convertToFloat64(expectedData)
	if err != nil {
		return false, err
	}

	actual, err := d.convertToFloat64(actualData)
	if err != nil {
		return false, err
	}

	expectedVar := stat.Variance(expected, nil)
	if expectedVar != 0 {
		return false, errors.New("Expected sample is not stable")
	}

	actualVar := stat.Variance(actual, nil)
	if actualVar != 0 {
		return false, errors.New("Actual sample is not stable")
	}

	return stat.Mean(expected, nil) == stat.Mean(actual, nil), nil
}

func (d *Data) Compare(expectedData, actualData []interface{}) (bool, error) {
	if len(expectedData) != len(actualData) {
		return false, errors.New("Sample size not match")
	}

	switch d.Type {
	case "int":
		fallthrough
	case "float":
		return d.compareFloat64(expectedData, actualData)
	}

	return false, errors.New("Unsupported type")
}

type Service struct {
	MatchedData map[string]Data
	OwnFields   []string
}

type Matcher struct {
	Services map[string]Service
	fields   []string
}

var m *Matcher

func init() {
	m = New()
}

func NewService() Service {
	return Service{
		MatchedData: map[string]Data{},
		OwnFields:   []string{},
	}
}

func (s Service) FieldExist(requestedField string) bool {
	for field := range s.MatchedData {
		if field == requestedField {
			return true
		}
	}

	for _, field := range s.OwnFields {
		if field == requestedField {
			return true
		}
	}

	return false
}

func (m *Matcher) ServiceExist(service string) bool {
	for srv := range m.Services {
		if srv == service {
			return true
		}
	}

	return false
}

func (m *Matcher) FieldExist(field string) bool {
	for _, foundField := range m.fields {
		if field == foundField {
			return true
		}
	}

	return false
}

func (m *Matcher) extractMatchedData(fields config.DataFields) {
	for field, fieldAttributes := range fields {
		srv := m.Services[fieldAttributes.Owner]
		srv.OwnFields = append(srv.OwnFields, field)
		m.Services[fieldAttributes.Owner] = srv

		for service, options := range fieldAttributes.Replicas {
			m.Services[service].MatchedData[field] = Data{
				Type:    fieldAttributes.Type,
				OwnedBy: fieldAttributes.Owner,
				Options: options,
			}
		}
	}
}

func (m *Matcher) extractServices(fields config.DataFields) {
	for _, field := range fields {
		if !m.ServiceExist(field.Owner) {
			m.Services[field.Owner] = NewService()
		}

		for serviceReplica := range field.Replicas {
			if !m.ServiceExist(serviceReplica) {
				m.Services[serviceReplica] = NewService()
			}
		}
	}
}

func (m *Matcher) buildFieldMap() {
	for _, srv := range m.Services {
		for field := range srv.MatchedData {
			if !m.FieldExist(field) {
				m.fields = append(m.fields, field)
			}
		}
	}
}

func Get() *Matcher {
	return m
}

func New() *Matcher {
	return &Matcher{
		Services: map[string]Service{},
		fields:   []string{},
	}
}

func InitializeWithData(data config.DataFields) {
	m.extractServices(data)
	m.extractMatchedData(data)
	m.buildFieldMap()
}
