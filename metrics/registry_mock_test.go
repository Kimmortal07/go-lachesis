// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Fantom-foundation/go-lachesis/metrics (interfaces: Registry)

// Package metrics is a generated GoMock package.
package metrics

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRegistry is a mock of Registry interface
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// Each mocks base method
func (m *MockRegistry) Each(arg0 RegistryEachFunc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Each", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Each indicates an expected call of Each
func (mr *MockRegistryMockRecorder) Each(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Each", reflect.TypeOf((*MockRegistry)(nil).Each), arg0)
}

// Get mocks base method
func (m *MockRegistry) Get(arg0 string) Metric {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(Metric)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockRegistryMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRegistry)(nil).Get), arg0)
}

// GetOrRegister mocks base method
func (m *MockRegistry) GetOrRegister(arg0 string, arg1 Metric) Metric {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrRegister", arg0, arg1)
	ret0, _ := ret[0].(Metric)
	return ret0
}

// GetOrRegister indicates an expected call of GetOrRegister
func (mr *MockRegistryMockRecorder) GetOrRegister(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrRegister", reflect.TypeOf((*MockRegistry)(nil).GetOrRegister), arg0, arg1)
}

// OnNew mocks base method
func (m *MockRegistry) OnNew(arg0 RegistryEachFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnNew", arg0)
}

// OnNew indicates an expected call of OnNew
func (mr *MockRegistryMockRecorder) OnNew(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnNew", reflect.TypeOf((*MockRegistry)(nil).OnNew), arg0)
}

// Register mocks base method
func (m *MockRegistry) Register(arg0 string, arg1 Metric) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockRegistryMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRegistry)(nil).Register), arg0, arg1)
}
