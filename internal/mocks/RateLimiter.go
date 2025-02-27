// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	context "context"

	action "github.com/bangumi/server/internal/web/rate/action"

	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/internal/model"

	rate "github.com/bangumi/server/internal/web/rate"
)

// RateLimiter is an autogenerated mock type for the Manager type
type RateLimiter struct {
	mock.Mock
}

type RateLimiter_Expecter struct {
	mock *mock.Mock
}

func (_m *RateLimiter) EXPECT() *RateLimiter_Expecter {
	return &RateLimiter_Expecter{mock: &_m.Mock}
}

// AllowAction provides a mock function with given fields: ctx, u, _a2, limit
func (_m *RateLimiter) AllowAction(ctx context.Context, u model.UserID, _a2 action.Action, limit rate.Limit) (bool, int, error) {
	ret := _m.Called(ctx, u, _a2, limit)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, model.UserID, action.Action, rate.Limit) bool); ok {
		r0 = rf(ctx, u, _a2, limit)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, model.UserID, action.Action, rate.Limit) int); ok {
		r1 = rf(ctx, u, _a2, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, model.UserID, action.Action, rate.Limit) error); ok {
		r2 = rf(ctx, u, _a2, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RateLimiter_AllowAction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllowAction'
type RateLimiter_AllowAction_Call struct {
	*mock.Call
}

// AllowAction is a helper method to define mock.On call
//   - ctx context.Context
//   - u model.UserID
//   - _a2 action.Action
//   - limit rate.Limit
func (_e *RateLimiter_Expecter) AllowAction(ctx interface{}, u interface{}, _a2 interface{}, limit interface{}) *RateLimiter_AllowAction_Call {
	return &RateLimiter_AllowAction_Call{Call: _e.mock.On("AllowAction", ctx, u, _a2, limit)}
}

func (_c *RateLimiter_AllowAction_Call) Run(run func(ctx context.Context, u model.UserID, _a2 action.Action, limit rate.Limit)) *RateLimiter_AllowAction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserID), args[2].(action.Action), args[3].(rate.Limit))
	})
	return _c
}

func (_c *RateLimiter_AllowAction_Call) Return(allowed bool, remain int, err error) *RateLimiter_AllowAction_Call {
	_c.Call.Return(allowed, remain, err)
	return _c
}

// Login provides a mock function with given fields: ctx, ip
func (_m *RateLimiter) Login(ctx context.Context, ip string) (bool, int, error) {
	ret := _m.Called(ctx, ip)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, ip)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, string) int); ok {
		r1 = rf(ctx, ip)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, ip)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RateLimiter_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type RateLimiter_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - ctx context.Context
//   - ip string
func (_e *RateLimiter_Expecter) Login(ctx interface{}, ip interface{}) *RateLimiter_Login_Call {
	return &RateLimiter_Login_Call{Call: _e.mock.On("Login", ctx, ip)}
}

func (_c *RateLimiter_Login_Call) Run(run func(ctx context.Context, ip string)) *RateLimiter_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *RateLimiter_Login_Call) Return(allowed bool, remain int, err error) *RateLimiter_Login_Call {
	_c.Call.Return(allowed, remain, err)
	return _c
}

// Reset provides a mock function with given fields: ctx, ip
func (_m *RateLimiter) Reset(ctx context.Context, ip string) error {
	ret := _m.Called(ctx, ip)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RateLimiter_Reset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reset'
type RateLimiter_Reset_Call struct {
	*mock.Call
}

// Reset is a helper method to define mock.On call
//   - ctx context.Context
//   - ip string
func (_e *RateLimiter_Expecter) Reset(ctx interface{}, ip interface{}) *RateLimiter_Reset_Call {
	return &RateLimiter_Reset_Call{Call: _e.mock.On("Reset", ctx, ip)}
}

func (_c *RateLimiter_Reset_Call) Run(run func(ctx context.Context, ip string)) *RateLimiter_Reset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *RateLimiter_Reset_Call) Return(_a0 error) *RateLimiter_Reset_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewRateLimiter interface {
	mock.TestingT
	Cleanup(func())
}

// NewRateLimiter creates a new instance of RateLimiter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRateLimiter(t mockConstructorTestingTNewRateLimiter) *RateLimiter {
	mock := &RateLimiter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
