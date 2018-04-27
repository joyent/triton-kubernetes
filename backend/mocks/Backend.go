// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import state "github.com/joyent/triton-kubernetes/state"

// Backend is an autogenerated mock type for the Backend type
type Backend struct {
	mock.Mock
}

// DeleteState provides a mock function with given fields: name
func (_m *Backend) DeleteState(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PersistState provides a mock function with given fields: _a0
func (_m *Backend) PersistState(_a0 state.State) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(state.State) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// State provides a mock function with given fields: name
func (_m *Backend) State(name string) (state.State, error) {
	ret := _m.Called(name)

	var r0 state.State
	if rf, ok := ret.Get(0).(func(string) state.State); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(state.State)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StateTerraformConfig provides a mock function with given fields: name
func (_m *Backend) StateTerraformConfig(name string) (string, interface{}) {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 interface{}
	if rf, ok := ret.Get(1).(func(string) interface{}); ok {
		r1 = rf(name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(interface{})
		}
	}

	return r0, r1
}

// States provides a mock function with given fields:
func (_m *Backend) States() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}