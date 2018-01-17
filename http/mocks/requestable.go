// Code generated by mockery v1.0.0
package mocks

import http "github.com/stone-payments/goclienttools/http"
import mock "github.com/stretchr/testify/mock"

// Requestable is an autogenerated mock type for the Requestable type
type Requestable struct {
	mock.Mock
}

// Delete provides a mock function with given fields: url, request
func (_m *Requestable) Delete(url string, request http.Options) (http.Response, error) {
	ret := _m.Called(url, request)

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(string, http.Options) http.Response); ok {
		r0 = rf(url, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, http.Options) error); ok {
		r1 = rf(url, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: url, request
func (_m *Requestable) Get(url string, request http.Options) (http.Response, error) {
	ret := _m.Called(url, request)

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(string, http.Options) http.Response); ok {
		r0 = rf(url, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, http.Options) error); ok {
		r1 = rf(url, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Post provides a mock function with given fields: url, request
func (_m *Requestable) Post(url string, request http.Options) (http.Response, error) {
	ret := _m.Called(url, request)

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(string, http.Options) http.Response); ok {
		r0 = rf(url, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, http.Options) error); ok {
		r1 = rf(url, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: url, request
func (_m *Requestable) Put(url string, request http.Options) (http.Response, error) {
	ret := _m.Called(url, request)

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(string, http.Options) http.Response); ok {
		r0 = rf(url, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, http.Options) error); ok {
		r1 = rf(url, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
