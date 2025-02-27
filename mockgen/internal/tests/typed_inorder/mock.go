// Code generated by MockGen. DO NOT EDIT.
// Source: input.go
//
// Generated by this command:
//
//	mockgen -package typed_inorder -source=input.go -destination=mock.go -typed
//

// Package typed_inorder is a generated GoMock package.
package typed_inorder

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAnimal is a mock of Animal interface.
type MockAnimal struct {
	ctrl     *gomock.Controller
	recorder *MockAnimalMockRecorder
}

// MockAnimalMockRecorder is the mock recorder for MockAnimal.
type MockAnimalMockRecorder struct {
	mock *MockAnimal
}

// NewMockAnimal creates a new mock instance.
func NewMockAnimal(ctrl *gomock.Controller) *MockAnimal {
	mock := &MockAnimal{ctrl: ctrl}
	mock.recorder = &MockAnimalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnimal) EXPECT() *MockAnimalMockRecorder {
	return m.recorder
}

// Feed mocks base method.
func (m *MockAnimal) Feed(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Feed", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Feed indicates an expected call of Feed.
func (mr *MockAnimalMockRecorder) Feed(arg0 any) *AnimalFeedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Feed", reflect.TypeOf((*MockAnimal)(nil).Feed), arg0)
	return &AnimalFeedCall{Call: call}
}

// AnimalFeedCall wrap *gomock.Call
type AnimalFeedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *AnimalFeedCall) Return(arg0 error) *AnimalFeedCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *AnimalFeedCall) Do(f func(string) error) *AnimalFeedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *AnimalFeedCall) DoAndReturn(f func(string) error) *AnimalFeedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSound mocks base method.
func (m *MockAnimal) GetSound() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSound")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSound indicates an expected call of GetSound.
func (mr *MockAnimalMockRecorder) GetSound() *AnimalGetSoundCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSound", reflect.TypeOf((*MockAnimal)(nil).GetSound))
	return &AnimalGetSoundCall{Call: call}
}

// AnimalGetSoundCall wrap *gomock.Call
type AnimalGetSoundCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *AnimalGetSoundCall) Return(arg0 string) *AnimalGetSoundCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *AnimalGetSoundCall) Do(f func() string) *AnimalGetSoundCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *AnimalGetSoundCall) DoAndReturn(f func() string) *AnimalGetSoundCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
