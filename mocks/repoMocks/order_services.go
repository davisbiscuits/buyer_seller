// Code generated by mockery v2.15.0. DO NOT EDIT.

package repoMocks

import (
	dtos "marketplace/dtos"

	mock "github.com/stretchr/testify/mock"
)

// IOrderRepository is an autogenerated mock type for the IOrderRepository type
type IOrderRepository struct {
	mock.Mock
}

// AcceptOrder provides a mock function with given fields: orderID
func (_m *IOrderRepository) AcceptOrder(orderID int) error {
	ret := _m.Called(orderID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(orderID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddOrder provides a mock function with given fields: productID, sellerID, buyerID
func (_m *IOrderRepository) AddOrder(productID int, sellerID int, buyerID int) (int64, error) {
	ret := _m.Called(productID, sellerID, buyerID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int, int, int) int64); ok {
		r0 = rf(productID, sellerID, buyerID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, int) error); ok {
		r1 = rf(productID, sellerID, buyerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuyerOrders provides a mock function with given fields: buyerID
func (_m *IOrderRepository) BuyerOrders(buyerID int) ([]dtos.BuyerOrdersResponse, error) {
	ret := _m.Called(buyerID)

	var r0 []dtos.BuyerOrdersResponse
	if rf, ok := ret.Get(0).(func(int) []dtos.BuyerOrdersResponse); ok {
		r0 = rf(buyerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dtos.BuyerOrdersResponse)
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

// GetOrder provides a mock function with given fields: orderID
func (_m *IOrderRepository) GetOrder(orderID int) (dtos.OrderResponse, error) {
	ret := _m.Called(orderID)

	var r0 dtos.OrderResponse
	if rf, ok := ret.Get(0).(func(int) dtos.OrderResponse); ok {
		r0 = rf(orderID)
	} else {
		r0 = ret.Get(0).(dtos.OrderResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RejectOrder provides a mock function with given fields: orderID
func (_m *IOrderRepository) RejectOrder(orderID int) error {
	ret := _m.Called(orderID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(orderID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SellerOrders provides a mock function with given fields: buyerID
func (_m *IOrderRepository) SellerOrders(buyerID int) ([]dtos.SellerOrdersResponse, error) {
	ret := _m.Called(buyerID)

	var r0 []dtos.SellerOrdersResponse
	if rf, ok := ret.Get(0).(func(int) []dtos.SellerOrdersResponse); ok {
		r0 = rf(buyerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dtos.SellerOrdersResponse)
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

type mockConstructorTestingTNewIOrderRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIOrderRepository creates a new instance of IOrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIOrderRepository(t mockConstructorTestingTNewIOrderRepository) *IOrderRepository {
	mock := &IOrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}