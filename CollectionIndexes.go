// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import driver "github.com/arangodb/go-driver"
import mock "github.com/stretchr/testify/mock"

// CollectionIndexes is an autogenerated mock type for the CollectionIndexes type
type CollectionIndexes struct {
	mock.Mock
}

// EnsureFullTextIndex provides a mock function with given fields: ctx, fields, options
func (_m *CollectionIndexes) EnsureFullTextIndex(ctx context.Context, fields []string, options *driver.EnsureFullTextIndexOptions) (driver.Index, bool, error) {
	ret := _m.Called(ctx, fields, options)

	var r0 driver.Index
	if rf, ok := ret.Get(0).(func(context.Context, []string, *driver.EnsureFullTextIndexOptions) driver.Index); ok {
		r0 = rf(ctx, fields, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Index)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, []string, *driver.EnsureFullTextIndexOptions) bool); ok {
		r1 = rf(ctx, fields, options)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, *driver.EnsureFullTextIndexOptions) error); ok {
		r2 = rf(ctx, fields, options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EnsureGeoIndex provides a mock function with given fields: ctx, fields, options
func (_m *CollectionIndexes) EnsureGeoIndex(ctx context.Context, fields []string, options *driver.EnsureGeoIndexOptions) (driver.Index, bool, error) {
	ret := _m.Called(ctx, fields, options)

	var r0 driver.Index
	if rf, ok := ret.Get(0).(func(context.Context, []string, *driver.EnsureGeoIndexOptions) driver.Index); ok {
		r0 = rf(ctx, fields, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Index)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, []string, *driver.EnsureGeoIndexOptions) bool); ok {
		r1 = rf(ctx, fields, options)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, *driver.EnsureGeoIndexOptions) error); ok {
		r2 = rf(ctx, fields, options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EnsureHashIndex provides a mock function with given fields: ctx, fields, options
func (_m *CollectionIndexes) EnsureHashIndex(ctx context.Context, fields []string, options *driver.EnsureHashIndexOptions) (driver.Index, bool, error) {
	ret := _m.Called(ctx, fields, options)

	var r0 driver.Index
	if rf, ok := ret.Get(0).(func(context.Context, []string, *driver.EnsureHashIndexOptions) driver.Index); ok {
		r0 = rf(ctx, fields, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Index)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, []string, *driver.EnsureHashIndexOptions) bool); ok {
		r1 = rf(ctx, fields, options)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, *driver.EnsureHashIndexOptions) error); ok {
		r2 = rf(ctx, fields, options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EnsurePersistentIndex provides a mock function with given fields: ctx, fields, options
func (_m *CollectionIndexes) EnsurePersistentIndex(ctx context.Context, fields []string, options *driver.EnsurePersistentIndexOptions) (driver.Index, bool, error) {
	ret := _m.Called(ctx, fields, options)

	var r0 driver.Index
	if rf, ok := ret.Get(0).(func(context.Context, []string, *driver.EnsurePersistentIndexOptions) driver.Index); ok {
		r0 = rf(ctx, fields, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Index)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, []string, *driver.EnsurePersistentIndexOptions) bool); ok {
		r1 = rf(ctx, fields, options)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, *driver.EnsurePersistentIndexOptions) error); ok {
		r2 = rf(ctx, fields, options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EnsureSkipListIndex provides a mock function with given fields: ctx, fields, options
func (_m *CollectionIndexes) EnsureSkipListIndex(ctx context.Context, fields []string, options *driver.EnsureSkipListIndexOptions) (driver.Index, bool, error) {
	ret := _m.Called(ctx, fields, options)

	var r0 driver.Index
	if rf, ok := ret.Get(0).(func(context.Context, []string, *driver.EnsureSkipListIndexOptions) driver.Index); ok {
		r0 = rf(ctx, fields, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Index)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, []string, *driver.EnsureSkipListIndexOptions) bool); ok {
		r1 = rf(ctx, fields, options)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, *driver.EnsureSkipListIndexOptions) error); ok {
		r2 = rf(ctx, fields, options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EnsureTTLIndex provides a mock function with given fields: ctx, field, expireAfter, options
func (_m *CollectionIndexes) EnsureTTLIndex(ctx context.Context, field string, expireAfter int, options *driver.EnsureTTLIndexOptions) (driver.Index, bool, error) {
	ret := _m.Called(ctx, field, expireAfter, options)

	var r0 driver.Index
	if rf, ok := ret.Get(0).(func(context.Context, string, int, *driver.EnsureTTLIndexOptions) driver.Index); ok {
		r0 = rf(ctx, field, expireAfter, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Index)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, string, int, *driver.EnsureTTLIndexOptions) bool); ok {
		r1 = rf(ctx, field, expireAfter, options)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, int, *driver.EnsureTTLIndexOptions) error); ok {
		r2 = rf(ctx, field, expireAfter, options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Index provides a mock function with given fields: ctx, name
func (_m *CollectionIndexes) Index(ctx context.Context, name string) (driver.Index, error) {
	ret := _m.Called(ctx, name)

	var r0 driver.Index
	if rf, ok := ret.Get(0).(func(context.Context, string) driver.Index); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Index)
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

// IndexExists provides a mock function with given fields: ctx, name
func (_m *CollectionIndexes) IndexExists(ctx context.Context, name string) (bool, error) {
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

// Indexes provides a mock function with given fields: ctx
func (_m *CollectionIndexes) Indexes(ctx context.Context) ([]driver.Index, error) {
	ret := _m.Called(ctx)

	var r0 []driver.Index
	if rf, ok := ret.Get(0).(func(context.Context) []driver.Index); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]driver.Index)
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