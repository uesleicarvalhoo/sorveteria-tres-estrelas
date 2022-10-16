// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	entity "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"

	time "time"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, s
func (_m *Repository) Create(ctx context.Context, s entity.Sale) error {
	ret := _m.Called(ctx, s)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Sale) error); ok {
		r0 = rf(ctx, s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repository) GetAll(ctx context.Context) ([]entity.Sale, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Sale
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Sale); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Sale)
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

// Search provides a mock function with given fields: ctx, start, end
func (_m *Repository) Search(ctx context.Context, start time.Time, end time.Time) ([]entity.Sale, error) {
	ret := _m.Called(ctx, start, end)

	var r0 []entity.Sale
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time) []entity.Sale); ok {
		r0 = rf(ctx, start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Sale)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, time.Time, time.Time) error); ok {
		r1 = rf(ctx, start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}