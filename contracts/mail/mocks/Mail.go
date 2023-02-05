// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	mail "github.com/yafgo/framework/contracts/mail"
)

// Mail is an autogenerated mock type for the Mail type
type Mail struct {
	mock.Mock
}

// Attach provides a mock function with given fields: files
func (_m *Mail) Attach(files []string) mail.Mail {
	ret := _m.Called(files)

	var r0 mail.Mail
	if rf, ok := ret.Get(0).(func([]string) mail.Mail); ok {
		r0 = rf(files)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mail.Mail)
		}
	}

	return r0
}

// Bcc provides a mock function with given fields: addresses
func (_m *Mail) Bcc(addresses []string) mail.Mail {
	ret := _m.Called(addresses)

	var r0 mail.Mail
	if rf, ok := ret.Get(0).(func([]string) mail.Mail); ok {
		r0 = rf(addresses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mail.Mail)
		}
	}

	return r0
}

// Cc provides a mock function with given fields: addresses
func (_m *Mail) Cc(addresses []string) mail.Mail {
	ret := _m.Called(addresses)

	var r0 mail.Mail
	if rf, ok := ret.Get(0).(func([]string) mail.Mail); ok {
		r0 = rf(addresses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mail.Mail)
		}
	}

	return r0
}

// Content provides a mock function with given fields: content
func (_m *Mail) Content(content mail.Content) mail.Mail {
	ret := _m.Called(content)

	var r0 mail.Mail
	if rf, ok := ret.Get(0).(func(mail.Content) mail.Mail); ok {
		r0 = rf(content)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mail.Mail)
		}
	}

	return r0
}

// From provides a mock function with given fields: address
func (_m *Mail) From(address mail.From) mail.Mail {
	ret := _m.Called(address)

	var r0 mail.Mail
	if rf, ok := ret.Get(0).(func(mail.From) mail.Mail); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mail.Mail)
		}
	}

	return r0
}

// Queue provides a mock function with given fields: queue
func (_m *Mail) Queue(queue *mail.Queue) error {
	ret := _m.Called(queue)

	var r0 error
	if rf, ok := ret.Get(0).(func(*mail.Queue) error); ok {
		r0 = rf(queue)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields:
func (_m *Mail) Send() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// To provides a mock function with given fields: addresses
func (_m *Mail) To(addresses []string) mail.Mail {
	ret := _m.Called(addresses)

	var r0 mail.Mail
	if rf, ok := ret.Get(0).(func([]string) mail.Mail); ok {
		r0 = rf(addresses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mail.Mail)
		}
	}

	return r0
}

type mockConstructorTestingTNewMail interface {
	mock.TestingT
	Cleanup(func())
}

// NewMail creates a new instance of Mail. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMail(t mockConstructorTestingTNewMail) *Mail {
	mock := &Mail{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
