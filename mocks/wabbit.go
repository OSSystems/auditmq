// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NeowayLabs/wabbit (interfaces: Delivery)

// Package mocks is a generated GoMock package.
package mocks

import (
	wabbit "github.com/NeowayLabs/wabbit"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDelivery is a mock of Delivery interface
type MockDelivery struct {
	ctrl     *gomock.Controller
	recorder *MockDeliveryMockRecorder
}

// MockDeliveryMockRecorder is the mock recorder for MockDelivery
type MockDeliveryMockRecorder struct {
	mock *MockDelivery
}

// NewMockDelivery creates a new mock instance
func NewMockDelivery(ctrl *gomock.Controller) *MockDelivery {
	mock := &MockDelivery{ctrl: ctrl}
	mock.recorder = &MockDeliveryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDelivery) EXPECT() *MockDeliveryMockRecorder {
	return m.recorder
}

// Ack mocks base method
func (m *MockDelivery) Ack(arg0 bool) error {
	ret := m.ctrl.Call(m, "Ack", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ack indicates an expected call of Ack
func (mr *MockDeliveryMockRecorder) Ack(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ack", reflect.TypeOf((*MockDelivery)(nil).Ack), arg0)
}

// Body mocks base method
func (m *MockDelivery) Body() []byte {
	ret := m.ctrl.Call(m, "Body")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Body indicates an expected call of Body
func (mr *MockDeliveryMockRecorder) Body() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockDelivery)(nil).Body))
}

// ConsumerTag mocks base method
func (m *MockDelivery) ConsumerTag() string {
	ret := m.ctrl.Call(m, "ConsumerTag")
	ret0, _ := ret[0].(string)
	return ret0
}

// ConsumerTag indicates an expected call of ConsumerTag
func (mr *MockDeliveryMockRecorder) ConsumerTag() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumerTag", reflect.TypeOf((*MockDelivery)(nil).ConsumerTag))
}

// DeliveryTag mocks base method
func (m *MockDelivery) DeliveryTag() uint64 {
	ret := m.ctrl.Call(m, "DeliveryTag")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// DeliveryTag indicates an expected call of DeliveryTag
func (mr *MockDeliveryMockRecorder) DeliveryTag() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeliveryTag", reflect.TypeOf((*MockDelivery)(nil).DeliveryTag))
}

// Headers mocks base method
func (m *MockDelivery) Headers() wabbit.Option {
	ret := m.ctrl.Call(m, "Headers")
	ret0, _ := ret[0].(wabbit.Option)
	return ret0
}

// Headers indicates an expected call of Headers
func (mr *MockDeliveryMockRecorder) Headers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Headers", reflect.TypeOf((*MockDelivery)(nil).Headers))
}

// MessageId mocks base method
func (m *MockDelivery) MessageId() string {
	ret := m.ctrl.Call(m, "MessageId")
	ret0, _ := ret[0].(string)
	return ret0
}

// MessageId indicates an expected call of MessageId
func (mr *MockDeliveryMockRecorder) MessageId() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageId", reflect.TypeOf((*MockDelivery)(nil).MessageId))
}

// Nack mocks base method
func (m *MockDelivery) Nack(arg0, arg1 bool) error {
	ret := m.ctrl.Call(m, "Nack", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Nack indicates an expected call of Nack
func (mr *MockDeliveryMockRecorder) Nack(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nack", reflect.TypeOf((*MockDelivery)(nil).Nack), arg0, arg1)
}

// Reject mocks base method
func (m *MockDelivery) Reject(arg0 bool) error {
	ret := m.ctrl.Call(m, "Reject", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reject indicates an expected call of Reject
func (mr *MockDeliveryMockRecorder) Reject(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reject", reflect.TypeOf((*MockDelivery)(nil).Reject), arg0)
}
