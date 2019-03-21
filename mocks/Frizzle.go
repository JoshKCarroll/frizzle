// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import frizzle "github.com/qntfy/frizzle"
import mock "github.com/stretchr/testify/mock"
import time "time"

// Frizzle is an autogenerated mock type for the Frizzle type
type Frizzle struct {
	mock.Mock
}

// Ack provides a mock function with given fields: _a0
func (_m *Frizzle) Ack(_a0 frizzle.Msg) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(frizzle.Msg) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddOptions provides a mock function with given fields: _a0
func (_m *Frizzle) AddOptions(_a0 ...frizzle.Option) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Close provides a mock function with given fields:
func (_m *Frizzle) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Events provides a mock function with given fields:
func (_m *Frizzle) Events() <-chan frizzle.Event {
	ret := _m.Called()

	var r0 <-chan frizzle.Event
	if rf, ok := ret.Get(0).(func() <-chan frizzle.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan frizzle.Event)
		}
	}

	return r0
}

// Fail provides a mock function with given fields: _a0
func (_m *Frizzle) Fail(_a0 frizzle.Msg) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(frizzle.Msg) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FlushAndClose provides a mock function with given fields: timeout
func (_m *Frizzle) FlushAndClose(timeout time.Duration) error {
	ret := _m.Called(timeout)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration) error); ok {
		r0 = rf(timeout)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Receive provides a mock function with given fields:
func (_m *Frizzle) Receive() <-chan frizzle.Msg {
	ret := _m.Called()

	var r0 <-chan frizzle.Msg
	if rf, ok := ret.Get(0).(func() <-chan frizzle.Msg); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan frizzle.Msg)
		}
	}

	return r0
}

// Send provides a mock function with given fields: m, dest
func (_m *Frizzle) Send(m frizzle.Msg, dest string) error {
	ret := _m.Called(m, dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(frizzle.Msg, string) error); ok {
		r0 = rf(m, dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
