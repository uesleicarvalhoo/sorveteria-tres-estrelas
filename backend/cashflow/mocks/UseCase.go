// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	cashflow "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/cashflow"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// GetCashFlow provides a mock function with given fields: ctx
func (_m *UseCase) GetCashFlow(ctx context.Context) (cashflow.CashFlow, error) {
	ret := _m.Called(ctx)

	var r0 cashflow.CashFlow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (cashflow.CashFlow, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) cashflow.CashFlow); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(cashflow.CashFlow)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCashFlowBetween provides a mock function with given fields: ctx, startAt, endAt
func (_m *UseCase) GetCashFlowBetween(ctx context.Context, startAt time.Time, endAt time.Time) (cashflow.CashFlow, error) {
	ret := _m.Called(ctx, startAt, endAt)

	var r0 cashflow.CashFlow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time) (cashflow.CashFlow, error)); ok {
		return rf(ctx, startAt, endAt)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time) cashflow.CashFlow); ok {
		r0 = rf(ctx, startAt, endAt)
	} else {
		r0 = ret.Get(0).(cashflow.CashFlow)
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time, time.Time) error); ok {
		r1 = rf(ctx, startAt, endAt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCase(t mockConstructorTestingTNewUseCase) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}