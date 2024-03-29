// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	product "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/product"

	uuid "github.com/google/uuid"
)

// Reader is an autogenerated mock type for the Reader type
type Reader struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, id
func (_m *Reader) Get(ctx context.Context, id uuid.UUID) (product.Product, error) {
	ret := _m.Called(ctx, id)

	var r0 product.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (product.Product, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) product.Product); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(product.Product)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *Reader) GetAll(ctx context.Context) ([]product.Product, error) {
	ret := _m.Called(ctx)

	var r0 []product.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]product.Product, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []product.Product); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewReader creates a new instance of Reader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *Reader {
	mock := &Reader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
