// Code generated by MockGen. DO NOT EDIT.
// Source: go_mock/calc/calc.go

// Package mock_calc is a generated GoMock package.
package mock_calc

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCalculate is a mock of Calculate interface.
type MockCalculate struct {
	ctrl     *gomock.Controller
	recorder *MockCalculateMockRecorder
}

// MockCalculateMockRecorder is the mock recorder for MockCalculate.
type MockCalculateMockRecorder struct {
	mock *MockCalculate
}

// NewMockCalculate creates a new mock instance.
func NewMockCalculate(ctrl *gomock.Controller) *MockCalculate {
	mock := &MockCalculate{ctrl: ctrl}
	mock.recorder = &MockCalculateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCalculate) EXPECT() *MockCalculateMockRecorder {
	return m.recorder
}

// Delta mocks base method.
func (m *MockCalculate) Delta(first, second int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", first, second)
	ret0, _ := ret[0].(int)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockCalculateMockRecorder) Delta(first, second interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockCalculate)(nil).Delta), first, second)
}

// Sum mocks base method.
func (m *MockCalculate) Sum(first, second int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sum", first, second)
	ret0, _ := ret[0].(int)
	return ret0
}

// Sum indicates an expected call of Sum.
func (mr *MockCalculateMockRecorder) Sum(first, second interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sum", reflect.TypeOf((*MockCalculate)(nil).Sum), first, second)
}
