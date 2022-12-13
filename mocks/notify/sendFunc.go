// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// SendFunc is an autogenerated mock type for the sendFunc type
type SendFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, subject, message
func (_m *SendFunc) Execute(ctx context.Context, subject string, message string) error {
	ret := _m.Called(ctx, subject, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, subject, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewSendFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewSendFunc creates a new instance of SendFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSendFunc(t mockConstructorTestingTNewSendFunc) *SendFunc {
	mock := &SendFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
