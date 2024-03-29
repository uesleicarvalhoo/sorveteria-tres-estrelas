// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	product "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/product"

	uuid "github.com/google/uuid"
)

// Writer is an autogenerated mock type for the Writer type
type Writer struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, p
func (_m *Writer) Create(ctx context.Context, p product.Product) error {
	ret := _m.Called(ctx, p)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, product.Product) error); ok {
		r0 = rf(ctx, p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Writer) Delete(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, p
func (_m *Writer) Update(ctx context.Context, p *product.Product) error {
	ret := _m.Called(ctx, p)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *product.Product) error); ok {
		r0 = rf(ctx, p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewWriter creates a new instance of Writer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWriter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Writer {
	mock := &Writer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
