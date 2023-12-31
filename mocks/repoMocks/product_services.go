// Code generated by mockery v2.15.0. DO NOT EDIT.

package repoMocks

import (
	models "marketplace/models"

	mock "github.com/stretchr/testify/mock"
)

// IProductRepository is an autogenerated mock type for the IProductRepository type
type IProductRepository struct {
	mock.Mock
}

// AddProduct provides a mock function with given fields: name, quantity, price, userID
func (_m *IProductRepository) AddProduct(name string, quantity int, price int, userID int) (int64, error) {
	ret := _m.Called(name, quantity, price, userID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, int, int, int) int64); ok {
		r0 = rf(name, quantity, price, userID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, int, int) error); ok {
		r1 = rf(name, quantity, price, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuyerProducts provides a mock function with given fields: buyerID
func (_m *IProductRepository) BuyerProducts(buyerID int) ([]models.Product, error) {
	ret := _m.Called(buyerID)

	var r0 []models.Product
	if rf, ok := ret.Get(0).(func(int) []models.Product); ok {
		r0 = rf(buyerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(buyerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProduct provides a mock function with given fields: productID
func (_m *IProductRepository) GetProduct(productID int) (models.Product, error) {
	ret := _m.Called(productID)

	var r0 models.Product
	if rf, ok := ret.Get(0).(func(int) models.Product); ok {
		r0 = rf(productID)
	} else {
		r0 = ret.Get(0).(models.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductByNameAndUser provides a mock function with given fields: productName, userID
func (_m *IProductRepository) GetProductByNameAndUser(productName string, userID int) (models.Product, error) {
	ret := _m.Called(productName, userID)

	var r0 models.Product
	if rf, ok := ret.Get(0).(func(string, int) models.Product); ok {
		r0 = rf(productName, userID)
	} else {
		r0 = ret.Get(0).(models.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(productName, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InvalidateProduct provides a mock function with given fields: productID
func (_m *IProductRepository) InvalidateProduct(productID int) error {
	ret := _m.Called(productID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(productID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SellerProducts provides a mock function with given fields: sellerID
func (_m *IProductRepository) SellerProducts(sellerID int) ([]models.Product, error) {
	ret := _m.Called(sellerID)

	var r0 []models.Product
	if rf, ok := ret.Get(0).(func(int) []models.Product); ok {
		r0 = rf(sellerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(sellerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProductQuantity provides a mock function with given fields: quantity, productID
func (_m *IProductRepository) UpdateProductQuantity(quantity int, productID int) error {
	ret := _m.Called(quantity, productID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(quantity, productID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIProductRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIProductRepository creates a new instance of IProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIProductRepository(t mockConstructorTestingTNewIProductRepository) *IProductRepository {
	mock := &IProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
