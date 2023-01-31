// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	http "github.com/yafgo/framework/contracts/http"
)

// Response is an autogenerated mock type for the Response type
type Response struct {
	mock.Mock
}

// Download provides a mock function with given fields: filepath, filename
func (_m *Response) Download(filepath string, filename string) {
	_m.Called(filepath, filename)
}

// File provides a mock function with given fields: filepath
func (_m *Response) File(filepath string) {
	_m.Called(filepath)
}

// Header provides a mock function with given fields: key, value
func (_m *Response) Header(key string, value string) http.Response {
	ret := _m.Called(key, value)

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(string, string) http.Response); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// Json provides a mock function with given fields: code, obj
func (_m *Response) Json(code int, obj interface{}) {
	_m.Called(code, obj)
}

// Origin provides a mock function with given fields:
func (_m *Response) Origin() http.ResponseOrigin {
	ret := _m.Called()

	var r0 http.ResponseOrigin
	if rf, ok := ret.Get(0).(func() http.ResponseOrigin); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.ResponseOrigin)
		}
	}

	return r0
}

// String provides a mock function with given fields: code, format, values
func (_m *Response) String(code int, format string, values ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, code, format)
	_ca = append(_ca, values...)
	_m.Called(_ca...)
}

// Success provides a mock function with given fields:
func (_m *Response) Success() http.ResponseSuccess {
	ret := _m.Called()

	var r0 http.ResponseSuccess
	if rf, ok := ret.Get(0).(func() http.ResponseSuccess); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.ResponseSuccess)
		}
	}

	return r0
}

type mockConstructorTestingTNewResponse interface {
	mock.TestingT
	Cleanup(func())
}

// NewResponse creates a new instance of Response. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewResponse(t mockConstructorTestingTNewResponse) *Response {
	mock := &Response{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
