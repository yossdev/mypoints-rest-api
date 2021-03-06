// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	entities "github.com/yossdev/mypoints-rest-api/src/rewards/entities"
)

// PsqlRepository is an autogenerated mock type for the PsqlRepository type
type PsqlRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: payload
func (_m *PsqlRepository) Create(payload entities.Domain) (int64, error) {
	ret := _m.Called(payload)

	var r0 int64
	if rf, ok := ret.Get(0).(func(entities.Domain) int64); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.Domain) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *PsqlRepository) Delete(id uint32) (int64, error) {
	ret := _m.Called(id)

	var r0 int64
	if rf, ok := ret.Get(0).(func(uint32) int64); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReward provides a mock function with given fields: id
func (_m *PsqlRepository) GetReward(id uint32) (entities.Domain, error) {
	ret := _m.Called(id)

	var r0 entities.Domain
	if rf, ok := ret.Get(0).(func(uint32) entities.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entities.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: payload
func (_m *PsqlRepository) Update(payload entities.Domain) (int64, error) {
	ret := _m.Called(payload)

	var r0 int64
	if rf, ok := ret.Get(0).(func(entities.Domain) int64); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.Domain) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
