// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import driver "github.com/arangodb/go-driver"
import mock "github.com/stretchr/testify/mock"

// Cluster is an autogenerated mock type for the Cluster type
type Cluster struct {
	mock.Mock
}

// CleanOutServer provides a mock function with given fields: ctx, serverID
func (_m *Cluster) CleanOutServer(ctx context.Context, serverID string) error {
	ret := _m.Called(ctx, serverID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, serverID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DatabaseInventory provides a mock function with given fields: ctx, db
func (_m *Cluster) DatabaseInventory(ctx context.Context, db driver.Database) (driver.DatabaseInventory, error) {
	ret := _m.Called(ctx, db)

	var r0 driver.DatabaseInventory
	if rf, ok := ret.Get(0).(func(context.Context, driver.Database) driver.DatabaseInventory); ok {
		r0 = rf(ctx, db)
	} else {
		r0 = ret.Get(0).(driver.DatabaseInventory)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, driver.Database) error); ok {
		r1 = rf(ctx, db)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Health provides a mock function with given fields: ctx
func (_m *Cluster) Health(ctx context.Context) (driver.ClusterHealth, error) {
	ret := _m.Called(ctx)

	var r0 driver.ClusterHealth
	if rf, ok := ret.Get(0).(func(context.Context) driver.ClusterHealth); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(driver.ClusterHealth)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsCleanedOut provides a mock function with given fields: ctx, serverID
func (_m *Cluster) IsCleanedOut(ctx context.Context, serverID string) (bool, error) {
	ret := _m.Called(ctx, serverID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, serverID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, serverID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MoveShard provides a mock function with given fields: ctx, col, shard, fromServer, toServer
func (_m *Cluster) MoveShard(ctx context.Context, col driver.Collection, shard driver.ShardID, fromServer driver.ServerID, toServer driver.ServerID) error {
	ret := _m.Called(ctx, col, shard, fromServer, toServer)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, driver.Collection, driver.ShardID, driver.ServerID, driver.ServerID) error); ok {
		r0 = rf(ctx, col, shard, fromServer, toServer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveServer provides a mock function with given fields: ctx, serverID
func (_m *Cluster) RemoveServer(ctx context.Context, serverID driver.ServerID) error {
	ret := _m.Called(ctx, serverID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, driver.ServerID) error); ok {
		r0 = rf(ctx, serverID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResignServer provides a mock function with given fields: ctx, serverID
func (_m *Cluster) ResignServer(ctx context.Context, serverID string) error {
	ret := _m.Called(ctx, serverID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, serverID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
