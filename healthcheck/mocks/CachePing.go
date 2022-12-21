// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CachePing is an autogenerated mock type for the CachePing type
type CachePing struct {
	mock.Mock
}

// Ping provides a mock function with given fields: ctx
func (_m *CachePing) Ping(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCachePing interface {
	mock.TestingT
	Cleanup(func())
}

// NewCachePing creates a new instance of CachePing. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCachePing(t mockConstructorTestingTNewCachePing) *CachePing {
	mock := &CachePing{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
