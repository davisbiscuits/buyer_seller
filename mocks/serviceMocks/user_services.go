// Code generated by mockery v2.15.0. DO NOT EDIT.

package serviceMocks

import (
	models "marketplace/models"

	mock "github.com/stretchr/testify/mock"
)

// IUserService is an autogenerated mock type for the IUserService type
type IUserService struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: userType, userEmail, userName
func (_m *IUserService) CreateUser(userType string, userEmail string, userName string) (int64, error) {
	ret := _m.Called(userType, userEmail, userName)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, string, string) int64); ok {
		r0 = rf(userType, userEmail, userName)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(userType, userEmail, userName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MatchSellers provides a mock function with given fields: productID
func (_m *IUserService) MatchSellers(productID int) ([]models.User, error) {
	ret := _m.Called(productID)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(int) []models.User); ok {
		r0 = rf(productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUserService creates a new instance of IUserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUserService(t mockConstructorTestingTNewIUserService) *IUserService {
	mock := &IUserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}