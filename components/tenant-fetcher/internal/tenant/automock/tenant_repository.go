// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package automock

import (
	context "context"

	model "github.com/kyma-incubator/compass/components/tenant-fetcher/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// TenantRepository is an autogenerated mock type for the TenantRepository type
type TenantRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, item
func (_m *TenantRepository) Create(ctx context.Context, item model.TenantModel) error {
	ret := _m.Called(ctx, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.TenantModel) error); ok {
		r0 = rf(ctx, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByExternalID provides a mock function with given fields: ctx, tenantId
func (_m *TenantRepository) GetByExternalID(ctx context.Context, tenantId string) (model.TenantModel, error) {
	ret := _m.Called(ctx, tenantId)

	var r0 model.TenantModel
	if rf, ok := ret.Get(0).(func(context.Context, string) model.TenantModel); ok {
		r0 = rf(ctx, tenantId)
	} else {
		r0 = ret.Get(0).(model.TenantModel)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tenantId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, item
func (_m *TenantRepository) Update(ctx context.Context, item model.TenantModel) error {
	ret := _m.Called(ctx, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.TenantModel) error); ok {
		r0 = rf(ctx, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
