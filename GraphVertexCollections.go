// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import driver "github.com/arangodb/go-driver"
import mock "github.com/stretchr/testify/mock"

// GraphVertexCollections is an autogenerated mock type for the GraphVertexCollections type
type GraphVertexCollections struct {
	mock.Mock
}

// CreateVertexCollection provides a mock function with given fields: ctx, collection
func (_m *GraphVertexCollections) CreateVertexCollection(ctx context.Context, collection string) (driver.Collection, error) {
	ret := _m.Called(ctx, collection)

	var r0 driver.Collection
	if rf, ok := ret.Get(0).(func(context.Context, string) driver.Collection); ok {
		r0 = rf(ctx, collection)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, collection)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VertexCollection provides a mock function with given fields: ctx, name
func (_m *GraphVertexCollections) VertexCollection(ctx context.Context, name string) (driver.Collection, error) {
	ret := _m.Called(ctx, name)

	var r0 driver.Collection
	if rf, ok := ret.Get(0).(func(context.Context, string) driver.Collection); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Collection)
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

// VertexCollectionExists provides a mock function with given fields: ctx, name
func (_m *GraphVertexCollections) VertexCollectionExists(ctx context.Context, name string) (bool, error) {
	ret := _m.Called(ctx, name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VertexCollections provides a mock function with given fields: ctx
func (_m *GraphVertexCollections) VertexCollections(ctx context.Context) ([]driver.Collection, error) {
	ret := _m.Called(ctx)

	var r0 []driver.Collection
	if rf, ok := ret.Get(0).(func(context.Context) []driver.Collection); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]driver.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
