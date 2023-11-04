// Code generated by MockGen. DO NOT EDIT.
// Source: ./receiver.go

// Package mock_logger is a generated GoMock package.
package mock_logger

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockReceiver is a mock of Receiver interface.
type MockReceiver struct {
	ctrl     *gomock.Controller
	recorder *MockReceiverMockRecorder
}

// MockReceiverMockRecorder is the mock recorder for MockReceiver.
type MockReceiverMockRecorder struct {
	mock *MockReceiver
}

// NewMockReceiver creates a new mock instance.
func NewMockReceiver(ctrl *gomock.Controller) *MockReceiver {
	mock := &MockReceiver{ctrl: ctrl}
	mock.recorder = &MockReceiverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReceiver) EXPECT() *MockReceiverMockRecorder {
	return m.recorder
}

// Subcribe mocks base method.
func (m *MockReceiver) Subcribe(topic string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subcribe", topic)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subcribe indicates an expected call of Subcribe.
func (mr *MockReceiverMockRecorder) Subcribe(topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subcribe", reflect.TypeOf((*MockReceiver)(nil).Subcribe), topic)
}
