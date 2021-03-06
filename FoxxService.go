// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import driver "github.com/arangodb/go-driver"
import mock "github.com/stretchr/testify/mock"

// FoxxService is an autogenerated mock type for the FoxxService type
type FoxxService struct {
	mock.Mock
}

// InstallFoxxService provides a mock function with given fields: ctx, zipFile, options
func (_m *FoxxService) InstallFoxxService(ctx context.Context, zipFile string, options driver.FoxxCreateOptions) error {
	ret := _m.Called(ctx, zipFile, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, driver.FoxxCreateOptions) error); ok {
		r0 = rf(ctx, zipFile, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UninstallFoxxService provides a mock function with given fields: ctx, options
func (_m *FoxxService) UninstallFoxxService(ctx context.Context, options driver.FoxxDeleteOptions) error {
	ret := _m.Called(ctx, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, driver.FoxxDeleteOptions) error); ok {
		r0 = rf(ctx, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
