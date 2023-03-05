// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// DatabasePing is an autogenerated mock type for the DatabasePing type
type DatabasePing struct {
	mock.Mock
}

// Ping provides a mock function with given fields:
func (_m *DatabasePing) Ping() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDatabasePing interface {
	mock.TestingT
	Cleanup(func())
}

// NewDatabasePing creates a new instance of DatabasePing. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDatabasePing(t mockConstructorTestingTNewDatabasePing) *DatabasePing {
	mock := &DatabasePing{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}