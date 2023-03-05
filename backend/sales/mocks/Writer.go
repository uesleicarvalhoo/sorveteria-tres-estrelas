// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	sales "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/sales"

	uuid "github.com/google/uuid"
)

// Writer is an autogenerated mock type for the Writer type
type Writer struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, s
func (_m *Writer) Create(ctx context.Context, s sales.Sale) error {
	ret := _m.Called(ctx, s)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, sales.Sale) error); ok {
		r0 = rf(ctx, s)
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

type mockConstructorTestingTNewWriter interface {
	mock.TestingT
	Cleanup(func())
}

// NewWriter creates a new instance of Writer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWriter(t mockConstructorTestingTNewWriter) *Writer {
	mock := &Writer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}