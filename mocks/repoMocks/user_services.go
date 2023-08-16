// Code generated by mockery v2.15.0. DO NOT EDIT.

package repoMocks

import (
	models "marketplace/models"

	mock "github.com/stretchr/testify/mock"
)

// IUserRepository is an autogenerated mock type for the IUserRepository type
type IUserRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: userType, userEmail, userName
func (_m *IUserRepository) CreateUser(userType string, userEmail string, userName string) (int64, error) {
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

// FindSellers provides a mock function with given fields: productName, productQuantity, productPrice
func (_m *IUserRepository) FindSellers(productName string, productQuantity int, productPrice int) ([]models.User, error) {
	ret := _m.Called(productName, productQuantity, productPrice)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(string, int, int) []models.User); ok {
		r0 = rf(productName, productQuantity, productPrice)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(productName, productQuantity, productPrice)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUserRepository creates a new instance of IUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUserRepository(t mockConstructorTestingTNewIUserRepository) *IUserRepository {
	mock := &IUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}