// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	entity "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"

	uuid "github.com/google/uuid"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *UseCase) Delete(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *UseCase) Get(ctx context.Context, id uuid.UUID) (entity.Popsicle, error) {
	ret := _m.Called(ctx, id)

	var r0 entity.Popsicle
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) entity.Popsicle); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.Popsicle)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Index provides a mock function with given fields: ctx
func (_m *UseCase) Index(ctx context.Context) ([]entity.Popsicle, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Popsicle
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Popsicle); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Popsicle)
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

// Store provides a mock function with given fields: ctx, flavor, price
func (_m *UseCase) Store(ctx context.Context, flavor string, price float64) (entity.Popsicle, error) {
	ret := _m.Called(ctx, flavor, price)

	var r0 entity.Popsicle
	if rf, ok := ret.Get(0).(func(context.Context, string, float64) entity.Popsicle); ok {
		r0 = rf(ctx, flavor, price)
	} else {
		r0 = ret.Get(0).(entity.Popsicle)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, float64) error); ok {
		r1 = rf(ctx, flavor, price)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, p
func (_m *UseCase) Update(ctx context.Context, p *entity.Popsicle) error {
	ret := _m.Called(ctx, p)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Popsicle) error); ok {
		r0 = rf(ctx, p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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
