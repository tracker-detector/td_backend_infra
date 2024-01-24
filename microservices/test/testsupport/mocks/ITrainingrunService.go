// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"
	models "tds/shared/models"

	mock "github.com/stretchr/testify/mock"
)

// ITrainingrunService is an autogenerated mock type for the ITrainingrunService type
type ITrainingrunService struct {
	mock.Mock
}

// DeleteAllByModelId provides a mock function with given fields: ctx, id
func (_m *ITrainingrunService) DeleteAllByModelId(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAllByModelId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByID provides a mock function with given fields: ctx, id
func (_m *ITrainingrunService) DeleteByID(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAllByModelId provides a mock function with given fields: ctx, modelId
func (_m *ITrainingrunService) FindAllByModelId(ctx context.Context, modelId string) ([]*models.TrainingRun, error) {
	ret := _m.Called(ctx, modelId)

	if len(ret) == 0 {
		panic("no return value specified for FindAllByModelId")
	}

	var r0 []*models.TrainingRun
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*models.TrainingRun, error)); ok {
		return rf(ctx, modelId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*models.TrainingRun); ok {
		r0 = rf(ctx, modelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.TrainingRun)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, modelId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllTrainingRuns provides a mock function with given fields: ctx
func (_m *ITrainingrunService) FindAllTrainingRuns(ctx context.Context) ([]*models.TrainingRun, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FindAllTrainingRuns")
	}

	var r0 []*models.TrainingRun
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*models.TrainingRun, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*models.TrainingRun); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.TrainingRun)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllTrainingRunsForModelname provides a mock function with given fields: ctx, modelName
func (_m *ITrainingrunService) FindAllTrainingRunsForModelname(ctx context.Context, modelName string) ([]*models.TrainingRun, error) {
	ret := _m.Called(ctx, modelName)

	if len(ret) == 0 {
		panic("no return value specified for FindAllTrainingRunsForModelname")
	}

	var r0 []*models.TrainingRun
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*models.TrainingRun, error)); ok {
		return rf(ctx, modelName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*models.TrainingRun); ok {
		r0 = rf(ctx, modelName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.TrainingRun)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, modelName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewITrainingrunService creates a new instance of ITrainingrunService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewITrainingrunService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ITrainingrunService {
	mock := &ITrainingrunService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}