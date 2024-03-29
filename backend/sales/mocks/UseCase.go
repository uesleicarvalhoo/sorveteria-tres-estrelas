// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	sales "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/sales"

	time "time"

	uuid "github.com/google/uuid"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// DeleteByID provides a mock function with given fields: ctx, id
func (_m *UseCase) DeleteByID(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *UseCase) GetAll(ctx context.Context) ([]sales.Sale, error) {
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

// GetByPeriod provides a mock function with given fields: ctx, start, end
func (_m *UseCase) GetByPeriod(ctx context.Context, start time.Time, end time.Time) ([]sales.Sale, error) {
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

// RegisterSale provides a mock function with given fields: ctx, desc, payment, cart
func (_m *UseCase) RegisterSale(ctx context.Context, desc string, payment sales.PaymentType, cart sales.Cart) (sales.Sale, error) {
	ret := _m.Called(ctx, desc, payment, cart)

	var r0 sales.Sale
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, sales.PaymentType, sales.Cart) (sales.Sale, error)); ok {
		return rf(ctx, desc, payment, cart)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, sales.PaymentType, sales.Cart) sales.Sale); ok {
		r0 = rf(ctx, desc, payment, cart)
	} else {
		r0 = ret.Get(0).(sales.Sale)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, sales.PaymentType, sales.Cart) error); ok {
		r1 = rf(ctx, desc, payment, cart)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
