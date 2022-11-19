// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	balances "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/balances"

	mock "github.com/stretchr/testify/mock"
)

// Reader is an autogenerated mock type for the Reader type
type Reader struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx
func (_m *Reader) GetAll(ctx context.Context) ([]balances.Balance, error) {
	ret := _m.Called(ctx)

	var r0 []balances.Balance
	if rf, ok := ret.Get(0).(func(context.Context) []balances.Balance); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]balances.Balance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewReader interface {
	mock.TestingT
	Cleanup(func())
}

// NewReader creates a new instance of Reader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReader(t mockConstructorTestingTNewReader) *Reader {
	mock := &Reader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}