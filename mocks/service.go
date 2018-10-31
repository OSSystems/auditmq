// Code generated by MockGen. DO NOT EDIT.
// Source: storage/service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	config "github.com/rodrigoapereira/auditmq/config"
	storage "github.com/rodrigoapereira/auditmq/storage"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockService) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockServiceMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockService)(nil).Name))
}

// Status mocks base method
func (m *MockService) Status() storage.ServiceStatus {
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(storage.ServiceStatus)
	return ret0
}

// Status indicates an expected call of Status
func (mr *MockServiceMockRecorder) Status() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockService)(nil).Status))
}

// SetStatus mocks base method
func (m *MockService) SetStatus(status storage.ServiceStatus) {
	m.ctrl.Call(m, "SetStatus", status)
}

// SetStatus indicates an expected call of SetStatus
func (mr *MockServiceMockRecorder) SetStatus(status interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockService)(nil).SetStatus), status)
}

// HasField mocks base method
func (m *MockService) HasField(fieldName string) bool {
	ret := m.ctrl.Call(m, "HasField", fieldName)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasField indicates an expected call of HasField
func (mr *MockServiceMockRecorder) HasField(fieldName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasField", reflect.TypeOf((*MockService)(nil).HasField), fieldName)
}

// GetField mocks base method
func (m *MockService) GetField(fieldName string) (storage.Field, error) {
	ret := m.ctrl.Call(m, "GetField", fieldName)
	ret0, _ := ret[0].(storage.Field)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetField indicates an expected call of GetField
func (mr *MockServiceMockRecorder) GetField(fieldName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetField", reflect.TypeOf((*MockService)(nil).GetField), fieldName)
}

// AddField mocks base method
func (m *MockService) AddField(fieldName string, opts config.ServiceData, owner storage.Service) error {
	ret := m.ctrl.Call(m, "AddField", fieldName, opts, owner)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddField indicates an expected call of AddField
func (mr *MockServiceMockRecorder) AddField(fieldName, opts, owner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddField", reflect.TypeOf((*MockService)(nil).AddField), fieldName, opts, owner)
}

// AddFieldWithOptions mocks base method
func (m *MockService) AddFieldWithOptions(fieldName string, opts config.ServiceData, owner storage.Service, options config.DataOptions) error {
	ret := m.ctrl.Call(m, "AddFieldWithOptions", fieldName, opts, owner, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFieldWithOptions indicates an expected call of AddFieldWithOptions
func (mr *MockServiceMockRecorder) AddFieldWithOptions(fieldName, opts, owner, options interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFieldWithOptions", reflect.TypeOf((*MockService)(nil).AddFieldWithOptions), fieldName, opts, owner, options)
}

// Push mocks base method
func (m *MockService) Push(fieldName string, data interface{}) error {
	ret := m.ctrl.Call(m, "Push", fieldName, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push
func (mr *MockServiceMockRecorder) Push(fieldName, data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockService)(nil).Push), fieldName, data)
}

// DuplicateLast mocks base method
func (m *MockService) DuplicateLast(fieldName string) error {
	ret := m.ctrl.Call(m, "DuplicateLast", fieldName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DuplicateLast indicates an expected call of DuplicateLast
func (mr *MockServiceMockRecorder) DuplicateLast(fieldName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DuplicateLast", reflect.TypeOf((*MockService)(nil).DuplicateLast), fieldName)
}

// Fields mocks base method
func (m *MockService) Fields() storage.Fields {
	ret := m.ctrl.Call(m, "Fields")
	ret0, _ := ret[0].(storage.Fields)
	return ret0
}

// Fields indicates an expected call of Fields
func (mr *MockServiceMockRecorder) Fields() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fields", reflect.TypeOf((*MockService)(nil).Fields))
}

// FieldsToMatch mocks base method
func (m *MockService) FieldsToMatch() storage.Fields {
	ret := m.ctrl.Call(m, "FieldsToMatch")
	ret0, _ := ret[0].(storage.Fields)
	return ret0
}

// FieldsToMatch indicates an expected call of FieldsToMatch
func (mr *MockServiceMockRecorder) FieldsToMatch() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FieldsToMatch", reflect.TypeOf((*MockService)(nil).FieldsToMatch))
}
