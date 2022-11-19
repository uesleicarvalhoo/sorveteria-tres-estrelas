// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	users "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users"

	uuid "github.com/google/uuid"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, name, email, password
func (_m *UseCase) Create(ctx context.Context, name string, email string, password string) (users.User, error) {
	ret := _m.Called(ctx, name, email, password)

	var r0 users.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) users.User); ok {
		r0 = rf(ctx, name, email, password)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, name, email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, id
func (_m *UseCase) Get(ctx context.Context, id uuid.UUID) (users.User, error) {
	ret := _m.Called(ctx, id)

	var r0 users.User
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) users.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *UseCase) GetByEmail(ctx context.Context, email string) (users.User, error) {
	ret := _m.Called(ctx, email)

	var r0 users.User
	if rf, ok := ret.Get(0).(func(context.Context, string) users.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
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
