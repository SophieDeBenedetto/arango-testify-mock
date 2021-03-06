// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import driver "github.com/arangodb/go-driver"
import mock "github.com/stretchr/testify/mock"

// Cursor is an autogenerated mock type for the Cursor type
type Cursor struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Cursor) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Count provides a mock function with given fields:
func (_m *Cursor) Count() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// HasMore provides a mock function with given fields:
func (_m *Cursor) HasMore() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ReadDocument provides a mock function with given fields: ctx, result
func (_m *Cursor) ReadDocument(ctx context.Context, result interface{}) (driver.DocumentMeta, error) {
	ret := _m.Called(ctx, result)

	var r0 driver.DocumentMeta
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) driver.DocumentMeta); ok {
		r0 = rf(ctx, result)
	} else {
		r0 = ret.Get(0).(driver.DocumentMeta)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, result)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Statistics provides a mock function with given fields:
func (_m *Cursor) Statistics() driver.QueryStatistics {
	ret := _m.Called()

	var r0 driver.QueryStatistics
	if rf, ok := ret.Get(0).(func() driver.QueryStatistics); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.QueryStatistics)
		}
	}

	return r0
}
