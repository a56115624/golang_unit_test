// Code generated by MockGen. DO NOT EDIT.
// Source: ./car/car.go

// Package gotest is a generated GoMock package.
package gotest

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockdataInterface is a mock of dataInterface interface.
type MockdataInterface struct {
	ctrl     *gomock.Controller
	recorder *MockdataInterfaceMockRecorder
}

// MockdataInterfaceMockRecorder is the mock recorder for MockdataInterface.
type MockdataInterfaceMockRecorder struct {
	mock *MockdataInterface
}

// NewMockdataInterface creates a new mock instance.
func NewMockdataInterface(ctrl *gomock.Controller) *MockdataInterface {
	mock := &MockdataInterface{ctrl: ctrl}
	mock.recorder = &MockdataInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdataInterface) EXPECT() *MockdataInterfaceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockdataInterface) Add(a, b float64) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", a, b)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockdataInterfaceMockRecorder) Add(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockdataInterface)(nil).Add), a, b)
}

// DOT mocks base method.
func (m *MockdataInterface) DOT(a, b float64) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DOT", a, b)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DOT indicates an expected call of DOT.
func (mr *MockdataInterfaceMockRecorder) DOT(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DOT", reflect.TypeOf((*MockdataInterface)(nil).DOT), a, b)
}

// Division mocks base method.
func (m *MockdataInterface) Division(a, b float64) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Division", a, b)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Division indicates an expected call of Division.
func (mr *MockdataInterfaceMockRecorder) Division(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Division", reflect.TypeOf((*MockdataInterface)(nil).Division), a, b)
}
