// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	models "github.com/keptn/go-utils/pkg/api/models"
	mock "github.com/stretchr/testify/mock"
)

// APIV1Interface is an autogenerated mock type for the APIV1Interface type
type APIV1Interface struct {
	mock.Mock
}

type APIV1Interface_Expecter struct {
	mock *mock.Mock
}

func (_m *APIV1Interface) EXPECT() *APIV1Interface_Expecter {
	return &APIV1Interface_Expecter{mock: &_m.Mock}
}

// CreateProject provides a mock function with given fields: project
func (_m *APIV1Interface) CreateProject(project models.CreateProject) (string, *models.Error) {
	ret := _m.Called(project)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.CreateProject) string); ok {
		r0 = rf(project)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *models.Error
	if rf, ok := ret.Get(1).(func(models.CreateProject) *models.Error); ok {
		r1 = rf(project)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.Error)
		}
	}

	return r0, r1
}

// APIV1Interface_CreateProject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProject'
type APIV1Interface_CreateProject_Call struct {
	*mock.Call
}

// CreateProject is a helper method to define mock.On call
//  - project models.CreateProject
func (_e *APIV1Interface_Expecter) CreateProject(project interface{}) *APIV1Interface_CreateProject_Call {
	return &APIV1Interface_CreateProject_Call{Call: _e.mock.On("CreateProject", project)}
}

func (_c *APIV1Interface_CreateProject_Call) Run(run func(project models.CreateProject)) *APIV1Interface_CreateProject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.CreateProject))
	})
	return _c
}

func (_c *APIV1Interface_CreateProject_Call) Return(_a0 string, _a1 *models.Error) *APIV1Interface_CreateProject_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// CreateService provides a mock function with given fields: project, service
func (_m *APIV1Interface) CreateService(project string, service models.CreateService) (string, *models.Error) {
	ret := _m.Called(project, service)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, models.CreateService) string); ok {
		r0 = rf(project, service)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *models.Error
	if rf, ok := ret.Get(1).(func(string, models.CreateService) *models.Error); ok {
		r1 = rf(project, service)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.Error)
		}
	}

	return r0, r1
}

// APIV1Interface_CreateService_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateService'
type APIV1Interface_CreateService_Call struct {
	*mock.Call
}

// CreateService is a helper method to define mock.On call
//  - project string
//  - service models.CreateService
func (_e *APIV1Interface_Expecter) CreateService(project interface{}, service interface{}) *APIV1Interface_CreateService_Call {
	return &APIV1Interface_CreateService_Call{Call: _e.mock.On("CreateService", project, service)}
}

func (_c *APIV1Interface_CreateService_Call) Run(run func(project string, service models.CreateService)) *APIV1Interface_CreateService_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(models.CreateService))
	})
	return _c
}

func (_c *APIV1Interface_CreateService_Call) Return(_a0 string, _a1 *models.Error) *APIV1Interface_CreateService_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// DeleteProject provides a mock function with given fields: project
func (_m *APIV1Interface) DeleteProject(project models.Project) (*models.DeleteProjectResponse, *models.Error) {
	ret := _m.Called(project)

	var r0 *models.DeleteProjectResponse
	if rf, ok := ret.Get(0).(func(models.Project) *models.DeleteProjectResponse); ok {
		r0 = rf(project)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.DeleteProjectResponse)
		}
	}

	var r1 *models.Error
	if rf, ok := ret.Get(1).(func(models.Project) *models.Error); ok {
		r1 = rf(project)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.Error)
		}
	}

	return r0, r1
}

// APIV1Interface_DeleteProject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteProject'
type APIV1Interface_DeleteProject_Call struct {
	*mock.Call
}

// DeleteProject is a helper method to define mock.On call
//  - project models.Project
func (_e *APIV1Interface_Expecter) DeleteProject(project interface{}) *APIV1Interface_DeleteProject_Call {
	return &APIV1Interface_DeleteProject_Call{Call: _e.mock.On("DeleteProject", project)}
}

func (_c *APIV1Interface_DeleteProject_Call) Run(run func(project models.Project)) *APIV1Interface_DeleteProject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.Project))
	})
	return _c
}

func (_c *APIV1Interface_DeleteProject_Call) Return(_a0 *models.DeleteProjectResponse, _a1 *models.Error) *APIV1Interface_DeleteProject_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// DeleteService provides a mock function with given fields: project, service
func (_m *APIV1Interface) DeleteService(project string, service string) (*models.DeleteServiceResponse, *models.Error) {
	ret := _m.Called(project, service)

	var r0 *models.DeleteServiceResponse
	if rf, ok := ret.Get(0).(func(string, string) *models.DeleteServiceResponse); ok {
		r0 = rf(project, service)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.DeleteServiceResponse)
		}
	}

	var r1 *models.Error
	if rf, ok := ret.Get(1).(func(string, string) *models.Error); ok {
		r1 = rf(project, service)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.Error)
		}
	}

	return r0, r1
}

// APIV1Interface_DeleteService_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteService'
type APIV1Interface_DeleteService_Call struct {
	*mock.Call
}

// DeleteService is a helper method to define mock.On call
//  - project string
//  - service string
func (_e *APIV1Interface_Expecter) DeleteService(project interface{}, service interface{}) *APIV1Interface_DeleteService_Call {
	return &APIV1Interface_DeleteService_Call{Call: _e.mock.On("DeleteService", project, service)}
}

func (_c *APIV1Interface_DeleteService_Call) Run(run func(project string, service string)) *APIV1Interface_DeleteService_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *APIV1Interface_DeleteService_Call) Return(_a0 *models.DeleteServiceResponse, _a1 *models.Error) *APIV1Interface_DeleteService_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetMetadata provides a mock function with given fields:
func (_m *APIV1Interface) GetMetadata() (*models.Metadata, *models.Error) {
	ret := _m.Called()

	var r0 *models.Metadata
	if rf, ok := ret.Get(0).(func() *models.Metadata); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Metadata)
		}
	}

	var r1 *models.Error
	if rf, ok := ret.Get(1).(func() *models.Error); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.Error)
		}
	}

	return r0, r1
}

// APIV1Interface_GetMetadata_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMetadata'
type APIV1Interface_GetMetadata_Call struct {
	*mock.Call
}

// GetMetadata is a helper method to define mock.On call
func (_e *APIV1Interface_Expecter) GetMetadata() *APIV1Interface_GetMetadata_Call {
	return &APIV1Interface_GetMetadata_Call{Call: _e.mock.On("GetMetadata")}
}

func (_c *APIV1Interface_GetMetadata_Call) Run(run func()) *APIV1Interface_GetMetadata_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *APIV1Interface_GetMetadata_Call) Return(_a0 *models.Metadata, _a1 *models.Error) *APIV1Interface_GetMetadata_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// SendEvent provides a mock function with given fields: event
func (_m *APIV1Interface) SendEvent(event models.KeptnContextExtendedCE) (*models.EventContext, *models.Error) {
	ret := _m.Called(event)

	var r0 *models.EventContext
	if rf, ok := ret.Get(0).(func(models.KeptnContextExtendedCE) *models.EventContext); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.EventContext)
		}
	}

	var r1 *models.Error
	if rf, ok := ret.Get(1).(func(models.KeptnContextExtendedCE) *models.Error); ok {
		r1 = rf(event)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.Error)
		}
	}

	return r0, r1
}

// APIV1Interface_SendEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendEvent'
type APIV1Interface_SendEvent_Call struct {
	*mock.Call
}

// SendEvent is a helper method to define mock.On call
//  - event models.KeptnContextExtendedCE
func (_e *APIV1Interface_Expecter) SendEvent(event interface{}) *APIV1Interface_SendEvent_Call {
	return &APIV1Interface_SendEvent_Call{Call: _e.mock.On("SendEvent", event)}
}

func (_c *APIV1Interface_SendEvent_Call) Run(run func(event models.KeptnContextExtendedCE)) *APIV1Interface_SendEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.KeptnContextExtendedCE))
	})
	return _c
}

func (_c *APIV1Interface_SendEvent_Call) Return(_a0 *models.EventContext, _a1 *models.Error) *APIV1Interface_SendEvent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// TriggerEvaluation provides a mock function with given fields: project, stage, service, evaluation
func (_m *APIV1Interface) TriggerEvaluation(project string, stage string, service string, evaluation models.Evaluation) (*models.EventContext, *models.Error) {
	ret := _m.Called(project, stage, service, evaluation)

	var r0 *models.EventContext
	if rf, ok := ret.Get(0).(func(string, string, string, models.Evaluation) *models.EventContext); ok {
		r0 = rf(project, stage, service, evaluation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.EventContext)
		}
	}

	var r1 *models.Error
	if rf, ok := ret.Get(1).(func(string, string, string, models.Evaluation) *models.Error); ok {
		r1 = rf(project, stage, service, evaluation)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.Error)
		}
	}

	return r0, r1
}

// APIV1Interface_TriggerEvaluation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TriggerEvaluation'
type APIV1Interface_TriggerEvaluation_Call struct {
	*mock.Call
}

// TriggerEvaluation is a helper method to define mock.On call
//  - project string
//  - stage string
//  - service string
//  - evaluation models.Evaluation
func (_e *APIV1Interface_Expecter) TriggerEvaluation(project interface{}, stage interface{}, service interface{}, evaluation interface{}) *APIV1Interface_TriggerEvaluation_Call {
	return &APIV1Interface_TriggerEvaluation_Call{Call: _e.mock.On("TriggerEvaluation", project, stage, service, evaluation)}
}

func (_c *APIV1Interface_TriggerEvaluation_Call) Run(run func(project string, stage string, service string, evaluation models.Evaluation)) *APIV1Interface_TriggerEvaluation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string), args[3].(models.Evaluation))
	})
	return _c
}

func (_c *APIV1Interface_TriggerEvaluation_Call) Return(_a0 *models.EventContext, _a1 *models.Error) *APIV1Interface_TriggerEvaluation_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateProject provides a mock function with given fields: project
func (_m *APIV1Interface) UpdateProject(project models.CreateProject) (string, *models.Error) {
	ret := _m.Called(project)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.CreateProject) string); ok {
		r0 = rf(project)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *models.Error
	if rf, ok := ret.Get(1).(func(models.CreateProject) *models.Error); ok {
		r1 = rf(project)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.Error)
		}
	}

	return r0, r1
}

// APIV1Interface_UpdateProject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateProject'
type APIV1Interface_UpdateProject_Call struct {
	*mock.Call
}

// UpdateProject is a helper method to define mock.On call
//  - project models.CreateProject
func (_e *APIV1Interface_Expecter) UpdateProject(project interface{}) *APIV1Interface_UpdateProject_Call {
	return &APIV1Interface_UpdateProject_Call{Call: _e.mock.On("UpdateProject", project)}
}

func (_c *APIV1Interface_UpdateProject_Call) Run(run func(project models.CreateProject)) *APIV1Interface_UpdateProject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.CreateProject))
	})
	return _c
}

func (_c *APIV1Interface_UpdateProject_Call) Return(_a0 string, _a1 *models.Error) *APIV1Interface_UpdateProject_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}
