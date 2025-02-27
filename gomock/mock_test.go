// Code generated by MockGen. DO NOT EDIT.
// Source: example_test.go
//
// Generated by this command:
//
//	mockgen -destination mock_test.go -package gomock_test -source example_test.go
//

// Package gomock_test is a generated GoMock package.
package gomock_test

import (
	reflect "reflect"

	"go.uber.org/mock/gomock"
)

// MockFoo is a mock of Foo interface.
type MockFoo struct {
	ctrl     *gomock.Controller
	recorder *MockFooMockRecorder
}

// MockFooMockRecorder is the mock recorder for MockFoo.
type MockFooMockRecorder struct {
	mock *MockFoo
}

// NewMockFoo creates a new mock instance.
func NewMockFoo(ctrl *gomock.Controller) *MockFoo {
	mock := &MockFoo{ctrl: ctrl}
	mock.recorder = &MockFooMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFoo) EXPECT() *MockFooMockRecorder {
	return m.recorder
}

// Bar mocks base method.
func (m *MockFoo) Bar(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bar", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Bar indicates an expected call of Bar.
func (mr *MockFooMockRecorder) Bar(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bar", reflect.TypeOf((*MockFoo)(nil).Bar), arg0)
}
