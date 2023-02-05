// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	queue "github.com/yafgo/framework/contracts/queue"
)

// Queue is an autogenerated mock type for the Queue type
type Queue struct {
	mock.Mock
}

// Chain provides a mock function with given fields: jobs
func (_m *Queue) Chain(jobs []queue.Jobs) queue.Task {
	ret := _m.Called(jobs)

	var r0 queue.Task
	if rf, ok := ret.Get(0).(func([]queue.Jobs) queue.Task); ok {
		r0 = rf(jobs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(queue.Task)
		}
	}

	return r0
}

// GetJobs provides a mock function with given fields:
func (_m *Queue) GetJobs() []queue.Job {
	ret := _m.Called()

	var r0 []queue.Job
	if rf, ok := ret.Get(0).(func() []queue.Job); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]queue.Job)
		}
	}

	return r0
}

// Job provides a mock function with given fields: job, args
func (_m *Queue) Job(job queue.Job, args []queue.Arg) queue.Task {
	ret := _m.Called(job, args)

	var r0 queue.Task
	if rf, ok := ret.Get(0).(func(queue.Job, []queue.Arg) queue.Task); ok {
		r0 = rf(job, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(queue.Task)
		}
	}

	return r0
}

// Register provides a mock function with given fields: jobs
func (_m *Queue) Register(jobs []queue.Job) {
	_m.Called(jobs)
}

// Worker provides a mock function with given fields: args
func (_m *Queue) Worker(args *queue.Args) queue.Worker {
	ret := _m.Called(args)

	var r0 queue.Worker
	if rf, ok := ret.Get(0).(func(*queue.Args) queue.Worker); ok {
		r0 = rf(args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(queue.Worker)
		}
	}

	return r0
}

type mockConstructorTestingTNewQueue interface {
	mock.TestingT
	Cleanup(func())
}

// NewQueue creates a new instance of Queue. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewQueue(t mockConstructorTestingTNewQueue) *Queue {
	mock := &Queue{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
