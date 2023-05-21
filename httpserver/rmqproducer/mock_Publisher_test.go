// Code generated by mockery v2.27.1. DO NOT EDIT.

package rmqproducer

import (
	mock "github.com/stretchr/testify/mock"
	rabbitmq "github.com/wagslane/go-rabbitmq"
)

// MockPublisher is an autogenerated mock type for the Publisher type
type MockPublisher struct {
	mock.Mock
}

type MockPublisher_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPublisher) EXPECT() *MockPublisher_Expecter {
	return &MockPublisher_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockPublisher) Close() {
	_m.Called()
}

// MockPublisher_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockPublisher_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockPublisher_Expecter) Close() *MockPublisher_Close_Call {
	return &MockPublisher_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockPublisher_Close_Call) Run(run func()) *MockPublisher_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPublisher_Close_Call) Return() *MockPublisher_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPublisher_Close_Call) RunAndReturn(run func()) *MockPublisher_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Publish provides a mock function with given fields: data, routingKeys, optionFuncs
func (_m *MockPublisher) Publish(data []byte, routingKeys []string, optionFuncs ...func(*rabbitmq.PublishOptions)) error {
	_va := make([]interface{}, len(optionFuncs))
	for _i := range optionFuncs {
		_va[_i] = optionFuncs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, data, routingKeys)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, []string, ...func(*rabbitmq.PublishOptions)) error); ok {
		r0 = rf(data, routingKeys, optionFuncs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPublisher_Publish_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Publish'
type MockPublisher_Publish_Call struct {
	*mock.Call
}

// Publish is a helper method to define mock.On call
//   - data []byte
//   - routingKeys []string
//   - optionFuncs ...func(*rabbitmq.PublishOptions)
func (_e *MockPublisher_Expecter) Publish(data interface{}, routingKeys interface{}, optionFuncs ...interface{}) *MockPublisher_Publish_Call {
	return &MockPublisher_Publish_Call{Call: _e.mock.On("Publish",
		append([]interface{}{data, routingKeys}, optionFuncs...)...)}
}

func (_c *MockPublisher_Publish_Call) Run(run func(data []byte, routingKeys []string, optionFuncs ...func(*rabbitmq.PublishOptions))) *MockPublisher_Publish_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*rabbitmq.PublishOptions), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*rabbitmq.PublishOptions))
			}
		}
		run(args[0].([]byte), args[1].([]string), variadicArgs...)
	})
	return _c
}

func (_c *MockPublisher_Publish_Call) Return(_a0 error) *MockPublisher_Publish_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPublisher_Publish_Call) RunAndReturn(run func([]byte, []string, ...func(*rabbitmq.PublishOptions)) error) *MockPublisher_Publish_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockPublisher interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPublisher creates a new instance of MockPublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPublisher(t mockConstructorTestingTNewMockPublisher) *MockPublisher {
	mock := &MockPublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}