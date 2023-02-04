// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	filesystem "github.com/yafgo/framework/contracts/filesystem"
	http "github.com/yafgo/framework/contracts/http"

	mock "github.com/stretchr/testify/mock"

	nethttp "net/http"

	validation "github.com/yafgo/framework/contracts/validation"
)

// Request is an autogenerated mock type for the Request type
type Request struct {
	mock.Mock
}

// AbortWithStatus provides a mock function with given fields: code
func (_m *Request) AbortWithStatus(code int) {
	_m.Called(code)
}

// AbortWithStatusJson provides a mock function with given fields: code, jsonObj
func (_m *Request) AbortWithStatusJson(code int, jsonObj any) {
	_m.Called(code, jsonObj)
}

// Bind provides a mock function with given fields: obj
func (_m *Request) Bind(obj any) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(any) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// File provides a mock function with given fields: name
func (_m *Request) File(name string) (filesystem.File, error) {
	ret := _m.Called(name)

	var r0 filesystem.File
	if rf, ok := ret.Get(0).(func(string) filesystem.File); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(filesystem.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Form provides a mock function with given fields: key, defaultValue
func (_m *Request) Form(key string, defaultValue string) string {
	ret := _m.Called(key, defaultValue)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(key, defaultValue)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// FullUrl provides a mock function with given fields:
func (_m *Request) FullUrl() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Header provides a mock function with given fields: key, defaultValue
func (_m *Request) Header(key string, defaultValue string) string {
	ret := _m.Called(key, defaultValue)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(key, defaultValue)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Headers provides a mock function with given fields:
func (_m *Request) Headers() nethttp.Header {
	ret := _m.Called()

	var r0 nethttp.Header
	if rf, ok := ret.Get(0).(func() nethttp.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(nethttp.Header)
		}
	}

	return r0
}

// Input provides a mock function with given fields: key
func (_m *Request) Input(key string) string {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Ip provides a mock function with given fields:
func (_m *Request) Ip() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Method provides a mock function with given fields:
func (_m *Request) Method() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Next provides a mock function with given fields:
func (_m *Request) Next() {
	_m.Called()
}

// Origin provides a mock function with given fields:
func (_m *Request) Origin() *nethttp.Request {
	ret := _m.Called()

	var r0 *nethttp.Request
	if rf, ok := ret.Get(0).(func() *nethttp.Request); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*nethttp.Request)
		}
	}

	return r0
}

// Path provides a mock function with given fields:
func (_m *Request) Path() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Query provides a mock function with given fields: key, defaultValue
func (_m *Request) Query(key string, defaultValue string) string {
	ret := _m.Called(key, defaultValue)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(key, defaultValue)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// QueryArray provides a mock function with given fields: key
func (_m *Request) QueryArray(key string) []string {
	ret := _m.Called(key)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// QueryMap provides a mock function with given fields: key
func (_m *Request) QueryMap(key string) map[string]string {
	ret := _m.Called(key)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string) map[string]string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// Url provides a mock function with given fields:
func (_m *Request) Url() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Validate provides a mock function with given fields: rules, options
func (_m *Request) Validate(rules map[string]string, options ...validation.Option) (validation.Validator, error) {
	_va := make([]any, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []any
	_ca = append(_ca, rules)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 validation.Validator
	if rf, ok := ret.Get(0).(func(map[string]string, ...validation.Option) validation.Validator); ok {
		r0 = rf(rules, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(validation.Validator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]string, ...validation.Option) error); ok {
		r1 = rf(rules, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateRequest provides a mock function with given fields: request
func (_m *Request) ValidateRequest(request http.FormRequest) (validation.Errors, error) {
	ret := _m.Called(request)

	var r0 validation.Errors
	if rf, ok := ret.Get(0).(func(http.FormRequest) validation.Errors); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(validation.Errors)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(http.FormRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRequest interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequest creates a new instance of Request. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequest(t mockConstructorTestingTNewRequest) *Request {
	mock := &Request{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
