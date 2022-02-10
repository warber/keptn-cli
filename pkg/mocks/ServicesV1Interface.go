// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	models "github.com/keptn/go-utils/pkg/api/models"
	mock "github.com/stretchr/testify/mock"
)

// ServicesV1Interface is an autogenerated mock type for the ServicesV1Interface type
type ServicesV1Interface struct {
	mock.Mock
}

type ServicesV1Interface_Expecter struct {
	mock *mock.Mock
}

func (_m *ServicesV1Interface) EXPECT() *ServicesV1Interface_Expecter {
	return &ServicesV1Interface_Expecter{mock: &_m.Mock}
}

// CreateServiceInStage provides a mock function with given fields: project, stage, serviceName
func (_m *ServicesV1Interface) CreateServiceInStage(project string, stage string, serviceName string) (*models.EventContext, *models.Error) {
	ret := _m.Called(project, stage, serviceName)

	var r0 *models.EventContext
	if rf, ok := ret.Get(0).(func(string, string, string) *models.EventContext); ok {
		r0 = rf(project, stage, serviceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.EventContext)
		}
	}

	var r1 *models.Error
	if rf, ok := ret.Get(1).(func(string, string, string) *models.Error); ok {
		r1 = rf(project, stage, serviceName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.Error)
		}
	}

	return r0, r1
}

// ServicesV1Interface_CreateServiceInStage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateServiceInStage'
type ServicesV1Interface_CreateServiceInStage_Call struct {
	*mock.Call
}

// CreateServiceInStage is a helper method to define mock.On call
//  - project string
//  - stage string
//  - serviceName string
func (_e *ServicesV1Interface_Expecter) CreateServiceInStage(project interface{}, stage interface{}, serviceName interface{}) *ServicesV1Interface_CreateServiceInStage_Call {
	return &ServicesV1Interface_CreateServiceInStage_Call{Call: _e.mock.On("CreateServiceInStage", project, stage, serviceName)}
}

func (_c *ServicesV1Interface_CreateServiceInStage_Call) Run(run func(project string, stage string, serviceName string)) *ServicesV1Interface_CreateServiceInStage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ServicesV1Interface_CreateServiceInStage_Call) Return(_a0 *models.EventContext, _a1 *models.Error) *ServicesV1Interface_CreateServiceInStage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// DeleteServiceFromStage provides a mock function with given fields: project, stage, serviceName
func (_m *ServicesV1Interface) DeleteServiceFromStage(project string, stage string, serviceName string) (*models.EventContext, *models.Error) {
	ret := _m.Called(project, stage, serviceName)

	var r0 *models.EventContext
	if rf, ok := ret.Get(0).(func(string, string, string) *models.EventContext); ok {
		r0 = rf(project, stage, serviceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.EventContext)
		}
	}

	var r1 *models.Error
	if rf, ok := ret.Get(1).(func(string, string, string) *models.Error); ok {
		r1 = rf(project, stage, serviceName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.Error)
		}
	}

	return r0, r1
}

// ServicesV1Interface_DeleteServiceFromStage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteServiceFromStage'
type ServicesV1Interface_DeleteServiceFromStage_Call struct {
	*mock.Call
}

// DeleteServiceFromStage is a helper method to define mock.On call
//  - project string
//  - stage string
//  - serviceName string
func (_e *ServicesV1Interface_Expecter) DeleteServiceFromStage(project interface{}, stage interface{}, serviceName interface{}) *ServicesV1Interface_DeleteServiceFromStage_Call {
	return &ServicesV1Interface_DeleteServiceFromStage_Call{Call: _e.mock.On("DeleteServiceFromStage", project, stage, serviceName)}
}

func (_c *ServicesV1Interface_DeleteServiceFromStage_Call) Run(run func(project string, stage string, serviceName string)) *ServicesV1Interface_DeleteServiceFromStage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ServicesV1Interface_DeleteServiceFromStage_Call) Return(_a0 *models.EventContext, _a1 *models.Error) *ServicesV1Interface_DeleteServiceFromStage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetAllServices provides a mock function with given fields: project, stage
func (_m *ServicesV1Interface) GetAllServices(project string, stage string) ([]*models.Service, error) {
	ret := _m.Called(project, stage)

	var r0 []*models.Service
	if rf, ok := ret.Get(0).(func(string, string) []*models.Service); ok {
		r0 = rf(project, stage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(project, stage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServicesV1Interface_GetAllServices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllServices'
type ServicesV1Interface_GetAllServices_Call struct {
	*mock.Call
}

// GetAllServices is a helper method to define mock.On call
//  - project string
//  - stage string
func (_e *ServicesV1Interface_Expecter) GetAllServices(project interface{}, stage interface{}) *ServicesV1Interface_GetAllServices_Call {
	return &ServicesV1Interface_GetAllServices_Call{Call: _e.mock.On("GetAllServices", project, stage)}
}

func (_c *ServicesV1Interface_GetAllServices_Call) Run(run func(project string, stage string)) *ServicesV1Interface_GetAllServices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *ServicesV1Interface_GetAllServices_Call) Return(_a0 []*models.Service, _a1 error) *ServicesV1Interface_GetAllServices_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetService provides a mock function with given fields: project, stage, service
func (_m *ServicesV1Interface) GetService(project string, stage string, service string) (*models.Service, error) {
	ret := _m.Called(project, stage, service)

	var r0 *models.Service
	if rf, ok := ret.Get(0).(func(string, string, string) *models.Service); ok {
		r0 = rf(project, stage, service)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(project, stage, service)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServicesV1Interface_GetService_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetService'
type ServicesV1Interface_GetService_Call struct {
	*mock.Call
}

// GetService is a helper method to define mock.On call
//  - project string
//  - stage string
//  - service string
func (_e *ServicesV1Interface_Expecter) GetService(project interface{}, stage interface{}, service interface{}) *ServicesV1Interface_GetService_Call {
	return &ServicesV1Interface_GetService_Call{Call: _e.mock.On("GetService", project, stage, service)}
}

func (_c *ServicesV1Interface_GetService_Call) Run(run func(project string, stage string, service string)) *ServicesV1Interface_GetService_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ServicesV1Interface_GetService_Call) Return(_a0 *models.Service, _a1 error) *ServicesV1Interface_GetService_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}
