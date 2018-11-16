// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NeowayLabs/wabbit (interfaces: Delivery,Channel)

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

// MockChannel is a mock of Channel interface
type MockChannel struct {
	ctrl     *gomock.Controller
	recorder *MockChannelMockRecorder
}

// MockChannelMockRecorder is the mock recorder for MockChannel
type MockChannelMockRecorder struct {
	mock *MockChannel
}

// NewMockChannel creates a new mock instance
func NewMockChannel(ctrl *gomock.Controller) *MockChannel {
	mock := &MockChannel{ctrl: ctrl}
	mock.recorder = &MockChannelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChannel) EXPECT() *MockChannelMockRecorder {
	return m.recorder
}

// Ack mocks base method
func (m *MockChannel) Ack(arg0 uint64, arg1 bool) error {
	ret := m.ctrl.Call(m, "Ack", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ack indicates an expected call of Ack
func (mr *MockChannelMockRecorder) Ack(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ack", reflect.TypeOf((*MockChannel)(nil).Ack), arg0, arg1)
}

// Cancel mocks base method
func (m *MockChannel) Cancel(arg0 string, arg1 bool) error {
	ret := m.ctrl.Call(m, "Cancel", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Cancel indicates an expected call of Cancel
func (mr *MockChannelMockRecorder) Cancel(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockChannel)(nil).Cancel), arg0, arg1)
}

// Close mocks base method
func (m *MockChannel) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockChannelMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockChannel)(nil).Close))
}

// Confirm mocks base method
func (m *MockChannel) Confirm(arg0 bool) error {
	ret := m.ctrl.Call(m, "Confirm", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Confirm indicates an expected call of Confirm
func (mr *MockChannelMockRecorder) Confirm(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Confirm", reflect.TypeOf((*MockChannel)(nil).Confirm), arg0)
}

// Consume mocks base method
func (m *MockChannel) Consume(arg0, arg1 string, arg2 wabbit.Option) (<-chan wabbit.Delivery, error) {
	ret := m.ctrl.Call(m, "Consume", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan wabbit.Delivery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume
func (mr *MockChannelMockRecorder) Consume(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockChannel)(nil).Consume), arg0, arg1, arg2)
}

// ExchangeDeclare mocks base method
func (m *MockChannel) ExchangeDeclare(arg0, arg1 string, arg2 wabbit.Option) error {
	ret := m.ctrl.Call(m, "ExchangeDeclare", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExchangeDeclare indicates an expected call of ExchangeDeclare
func (mr *MockChannelMockRecorder) ExchangeDeclare(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangeDeclare", reflect.TypeOf((*MockChannel)(nil).ExchangeDeclare), arg0, arg1, arg2)
}

// ExchangeDeclarePassive mocks base method
func (m *MockChannel) ExchangeDeclarePassive(arg0, arg1 string, arg2 wabbit.Option) error {
	ret := m.ctrl.Call(m, "ExchangeDeclarePassive", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExchangeDeclarePassive indicates an expected call of ExchangeDeclarePassive
func (mr *MockChannelMockRecorder) ExchangeDeclarePassive(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangeDeclarePassive", reflect.TypeOf((*MockChannel)(nil).ExchangeDeclarePassive), arg0, arg1, arg2)
}

// Nack mocks base method
func (m *MockChannel) Nack(arg0 uint64, arg1, arg2 bool) error {
	ret := m.ctrl.Call(m, "Nack", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Nack indicates an expected call of Nack
func (mr *MockChannelMockRecorder) Nack(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nack", reflect.TypeOf((*MockChannel)(nil).Nack), arg0, arg1, arg2)
}

// NotifyClose mocks base method
func (m *MockChannel) NotifyClose(arg0 chan wabbit.Error) chan wabbit.Error {
	ret := m.ctrl.Call(m, "NotifyClose", arg0)
	ret0, _ := ret[0].(chan wabbit.Error)
	return ret0
}

// NotifyClose indicates an expected call of NotifyClose
func (mr *MockChannelMockRecorder) NotifyClose(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyClose", reflect.TypeOf((*MockChannel)(nil).NotifyClose), arg0)
}

// NotifyPublish mocks base method
func (m *MockChannel) NotifyPublish(arg0 chan wabbit.Confirmation) chan wabbit.Confirmation {
	ret := m.ctrl.Call(m, "NotifyPublish", arg0)
	ret0, _ := ret[0].(chan wabbit.Confirmation)
	return ret0
}

// NotifyPublish indicates an expected call of NotifyPublish
func (mr *MockChannelMockRecorder) NotifyPublish(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyPublish", reflect.TypeOf((*MockChannel)(nil).NotifyPublish), arg0)
}

// Publish mocks base method
func (m *MockChannel) Publish(arg0, arg1 string, arg2 []byte, arg3 wabbit.Option) error {
	ret := m.ctrl.Call(m, "Publish", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish
func (mr *MockChannelMockRecorder) Publish(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockChannel)(nil).Publish), arg0, arg1, arg2, arg3)
}

// Qos mocks base method
func (m *MockChannel) Qos(arg0, arg1 int, arg2 bool) error {
	ret := m.ctrl.Call(m, "Qos", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Qos indicates an expected call of Qos
func (mr *MockChannelMockRecorder) Qos(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Qos", reflect.TypeOf((*MockChannel)(nil).Qos), arg0, arg1, arg2)
}

// QueueBind mocks base method
func (m *MockChannel) QueueBind(arg0, arg1, arg2 string, arg3 wabbit.Option) error {
	ret := m.ctrl.Call(m, "QueueBind", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueueBind indicates an expected call of QueueBind
func (mr *MockChannelMockRecorder) QueueBind(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueBind", reflect.TypeOf((*MockChannel)(nil).QueueBind), arg0, arg1, arg2, arg3)
}

// QueueDeclare mocks base method
func (m *MockChannel) QueueDeclare(arg0 string, arg1 wabbit.Option) (wabbit.Queue, error) {
	ret := m.ctrl.Call(m, "QueueDeclare", arg0, arg1)
	ret0, _ := ret[0].(wabbit.Queue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueDeclare indicates an expected call of QueueDeclare
func (mr *MockChannelMockRecorder) QueueDeclare(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueDeclare", reflect.TypeOf((*MockChannel)(nil).QueueDeclare), arg0, arg1)
}

// QueueDeclarePassive mocks base method
func (m *MockChannel) QueueDeclarePassive(arg0 string, arg1 wabbit.Option) (wabbit.Queue, error) {
	ret := m.ctrl.Call(m, "QueueDeclarePassive", arg0, arg1)
	ret0, _ := ret[0].(wabbit.Queue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueDeclarePassive indicates an expected call of QueueDeclarePassive
func (mr *MockChannelMockRecorder) QueueDeclarePassive(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueDeclarePassive", reflect.TypeOf((*MockChannel)(nil).QueueDeclarePassive), arg0, arg1)
}

// QueueDelete mocks base method
func (m *MockChannel) QueueDelete(arg0 string, arg1 wabbit.Option) (int, error) {
	ret := m.ctrl.Call(m, "QueueDelete", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueDelete indicates an expected call of QueueDelete
func (mr *MockChannelMockRecorder) QueueDelete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueDelete", reflect.TypeOf((*MockChannel)(nil).QueueDelete), arg0, arg1)
}

// QueueUnbind mocks base method
func (m *MockChannel) QueueUnbind(arg0, arg1, arg2 string, arg3 wabbit.Option) error {
	ret := m.ctrl.Call(m, "QueueUnbind", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueueUnbind indicates an expected call of QueueUnbind
func (mr *MockChannelMockRecorder) QueueUnbind(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueUnbind", reflect.TypeOf((*MockChannel)(nil).QueueUnbind), arg0, arg1, arg2, arg3)
}

// Reject mocks base method
func (m *MockChannel) Reject(arg0 uint64, arg1 bool) error {
	ret := m.ctrl.Call(m, "Reject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reject indicates an expected call of Reject
func (mr *MockChannelMockRecorder) Reject(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reject", reflect.TypeOf((*MockChannel)(nil).Reject), arg0, arg1)
}
