// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import driver "github.com/arangodb/go-driver"
import mock "github.com/stretchr/testify/mock"

// ClientBackup is an autogenerated mock type for the ClientBackup type
type ClientBackup struct {
	mock.Mock
}

// Abort provides a mock function with given fields: ctx, job
func (_m *ClientBackup) Abort(ctx context.Context, job driver.BackupTransferJobID) error {
	ret := _m.Called(ctx, job)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, driver.BackupTransferJobID) error); ok {
		r0 = rf(ctx, job)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx, opt
func (_m *ClientBackup) Create(ctx context.Context, opt *driver.BackupCreateOptions) (driver.BackupID, driver.BackupCreateResponse, error) {
	ret := _m.Called(ctx, opt)

	var r0 driver.BackupID
	if rf, ok := ret.Get(0).(func(context.Context, *driver.BackupCreateOptions) driver.BackupID); ok {
		r0 = rf(ctx, opt)
	} else {
		r0 = ret.Get(0).(driver.BackupID)
	}

	var r1 driver.BackupCreateResponse
	if rf, ok := ret.Get(1).(func(context.Context, *driver.BackupCreateOptions) driver.BackupCreateResponse); ok {
		r1 = rf(ctx, opt)
	} else {
		r1 = ret.Get(1).(driver.BackupCreateResponse)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *driver.BackupCreateOptions) error); ok {
		r2 = rf(ctx, opt)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ClientBackup) Delete(ctx context.Context, id driver.BackupID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, driver.BackupID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Download provides a mock function with given fields: ctx, id, remoteRepository, config
func (_m *ClientBackup) Download(ctx context.Context, id driver.BackupID, remoteRepository string, config interface{}) (driver.BackupTransferJobID, error) {
	ret := _m.Called(ctx, id, remoteRepository, config)

	var r0 driver.BackupTransferJobID
	if rf, ok := ret.Get(0).(func(context.Context, driver.BackupID, string, interface{}) driver.BackupTransferJobID); ok {
		r0 = rf(ctx, id, remoteRepository, config)
	} else {
		r0 = ret.Get(0).(driver.BackupTransferJobID)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, driver.BackupID, string, interface{}) error); ok {
		r1 = rf(ctx, id, remoteRepository, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, opt
func (_m *ClientBackup) List(ctx context.Context, opt *driver.BackupListOptions) (map[driver.BackupID]driver.BackupMeta, error) {
	ret := _m.Called(ctx, opt)

	var r0 map[driver.BackupID]driver.BackupMeta
	if rf, ok := ret.Get(0).(func(context.Context, *driver.BackupListOptions) map[driver.BackupID]driver.BackupMeta); ok {
		r0 = rf(ctx, opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[driver.BackupID]driver.BackupMeta)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *driver.BackupListOptions) error); ok {
		r1 = rf(ctx, opt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Progress provides a mock function with given fields: ctx, job
func (_m *ClientBackup) Progress(ctx context.Context, job driver.BackupTransferJobID) (driver.BackupTransferProgressReport, error) {
	ret := _m.Called(ctx, job)

	var r0 driver.BackupTransferProgressReport
	if rf, ok := ret.Get(0).(func(context.Context, driver.BackupTransferJobID) driver.BackupTransferProgressReport); ok {
		r0 = rf(ctx, job)
	} else {
		r0 = ret.Get(0).(driver.BackupTransferProgressReport)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, driver.BackupTransferJobID) error); ok {
		r1 = rf(ctx, job)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Restore provides a mock function with given fields: ctx, id, opt
func (_m *ClientBackup) Restore(ctx context.Context, id driver.BackupID, opt *driver.BackupRestoreOptions) error {
	ret := _m.Called(ctx, id, opt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, driver.BackupID, *driver.BackupRestoreOptions) error); ok {
		r0 = rf(ctx, id, opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Upload provides a mock function with given fields: ctx, id, remoteRepository, config
func (_m *ClientBackup) Upload(ctx context.Context, id driver.BackupID, remoteRepository string, config interface{}) (driver.BackupTransferJobID, error) {
	ret := _m.Called(ctx, id, remoteRepository, config)

	var r0 driver.BackupTransferJobID
	if rf, ok := ret.Get(0).(func(context.Context, driver.BackupID, string, interface{}) driver.BackupTransferJobID); ok {
		r0 = rf(ctx, id, remoteRepository, config)
	} else {
		r0 = ret.Get(0).(driver.BackupTransferJobID)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, driver.BackupID, string, interface{}) error); ok {
		r1 = rf(ctx, id, remoteRepository, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
