// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	auth "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"

	mock "github.com/stretchr/testify/mock"

	user "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/user"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Authorize provides a mock function with given fields: ctx, token
func (_m *UseCase) Authorize(ctx context.Context, token string) (user.User, error) {
	ret := _m.Called(ctx, token)

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (user.User, error)); ok {
		return rf(ctx, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) user.User); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, payload
func (_m *UseCase) Login(ctx context.Context, payload auth.LoginPayload) (auth.JwtToken, error) {
	ret := _m.Called(ctx, payload)

	var r0 auth.JwtToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, auth.LoginPayload) (auth.JwtToken, error)); ok {
		return rf(ctx, payload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, auth.LoginPayload) auth.JwtToken); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(auth.JwtToken)
	}

	if rf, ok := ret.Get(1).(func(context.Context, auth.LoginPayload) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RefreshToken provides a mock function with given fields: ctx, payload
func (_m *UseCase) RefreshToken(ctx context.Context, payload auth.RefreshTokenPayload) (auth.JwtToken, error) {
	ret := _m.Called(ctx, payload)

	var r0 auth.JwtToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, auth.RefreshTokenPayload) (auth.JwtToken, error)); ok {
		return rf(ctx, payload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, auth.RefreshTokenPayload) auth.JwtToken); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(auth.JwtToken)
	}

	if rf, ok := ret.Get(1).(func(context.Context, auth.RefreshTokenPayload) error); ok {
		r1 = rf(ctx, payload)
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
