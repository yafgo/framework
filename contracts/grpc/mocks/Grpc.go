// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// Grpc is an autogenerated mock type for the Grpc type
type Grpc struct {
	mock.Mock
}

// Client provides a mock function with given fields: ctx, name
func (_m *Grpc) Client(ctx context.Context, name string) (*grpc.ClientConn, error) {
	ret := _m.Called(ctx, name)

	var r0 *grpc.ClientConn
	if rf, ok := ret.Get(0).(func(context.Context, string) *grpc.ClientConn); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grpc.ClientConn)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Run provides a mock function with given fields: host
func (_m *Grpc) Run(host string) error {
	ret := _m.Called(host)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(host)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Server provides a mock function with given fields:
func (_m *Grpc) Server() *grpc.Server {
	ret := _m.Called()

	var r0 *grpc.Server
	if rf, ok := ret.Get(0).(func() *grpc.Server); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grpc.Server)
		}
	}

	return r0
}

// UnaryClientInterceptorGroups provides a mock function with given fields: _a0
func (_m *Grpc) UnaryClientInterceptorGroups(_a0 map[string][]grpc.UnaryClientInterceptor) {
	_m.Called(_a0)
}

// UnaryServerInterceptors provides a mock function with given fields: _a0
func (_m *Grpc) UnaryServerInterceptors(_a0 []grpc.UnaryServerInterceptor) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewGrpc interface {
	mock.TestingT
	Cleanup(func())
}

// NewGrpc creates a new instance of Grpc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGrpc(t mockConstructorTestingTNewGrpc) *Grpc {
	mock := &Grpc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
