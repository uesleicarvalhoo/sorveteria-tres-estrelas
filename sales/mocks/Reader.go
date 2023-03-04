// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	sales "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"

	time "time"
)

// Reader is an autogenerated mock type for the Reader type
type Reader struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx
func (_m *Reader) GetAll(ctx context.Context) ([]sales.Sale, error) {
	ret := _m.Called(ctx)

	var r0 []sales.Sale
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]sales.Sale, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []sales.Sale); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sales.Sale)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: ctx, start, end
func (_m *Reader) Search(ctx context.Context, start time.Time, end time.Time) ([]sales.Sale, error) {
	ret := _m.Called(ctx, start, end)

	var r0 []sales.Sale
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time) ([]sales.Sale, error)); ok {
		return rf(ctx, start, end)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time) []sales.Sale); ok {
		r0 = rf(ctx, start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sales.Sale)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time, time.Time) error); ok {
		r1 = rf(ctx, start, end)
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
