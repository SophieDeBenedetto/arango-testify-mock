// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import protocol "github.com/arangodb/go-driver/vst/protocol"

// messageTransport is an autogenerated mock type for the messageTransport type
type messageTransport struct {
	mock.Mock
}

// Send provides a mock function with given fields: ctx, messageParts
func (_m *messageTransport) Send(ctx context.Context, messageParts ...[]byte) (<-chan protocol.Message, error) {
	_va := make([]interface{}, len(messageParts))
	for _i := range messageParts {
		_va[_i] = messageParts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 <-chan protocol.Message
	if rf, ok := ret.Get(0).(func(context.Context, ...[]byte) <-chan protocol.Message); ok {
		r0 = rf(ctx, messageParts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan protocol.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...[]byte) error); ok {
		r1 = rf(ctx, messageParts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
