package order

import (
	"errors"
	"marketplace/dtos"
	"marketplace/mocks/repoMocks"
	"marketplace/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPlaceBuyerOrder(t *testing.T) {
	t.Run("successfully place order", func(t *testing.T) {
		mockOrderRepo := new(repoMocks.IOrderRepository)
		mockProductRepo := new(repoMocks.IProductRepository)
		expectedID := int64(1)

		mockOrderRepo.On("AddOrder", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
			expectedID, nil,
		)
		orderService := New(mockOrderRepo, mockProductRepo)

		ID, err := orderService.PlaceBuyerOrder(14, 10, 11)
		assert.Equal(t, ID, expectedID)
		assert.Equal(t, err, nil)
	})
	t.Run("error placing order", func(t *testing.T) {
		mockOrderRepo := new(repoMocks.IOrderRepository)
		mockProductRepo := new(repoMocks.IProductRepository)

		mockOrderRepo.On("AddOrder", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
			int64(-1), errors.New("some error"),
		)
		orderService := New(mockOrderRepo, mockProductRepo)

		ID, err := orderService.PlaceBuyerOrder(14, 10, 11)
		assert.Equal(t, ID, int64(-1))
		assert.NotEqual(t, err, nil)
	})
}

func TestAcceptOrder(t *testing.T) {
	t.Run("successfully Accept order seller inventory >  buyer order", func(t *testing.T) {
		mockOrderRepo := new(repoMocks.IOrderRepository)
		mockProductRepo := new(repoMocks.IProductRepository)
		buyerOrder := &dtos.OrderResponse{
			ID:      1,
			Buyer:   dtos.UserResponse{ID: 1, Name: "buyer_name", Email: "buyer@email.com"},
			Seller:  dtos.UserResponse{ID: 2, Name: "seller_name", Email: "seller@email.com"},
			Product: dtos.ProductResponse{ID: 1, Name: "product_name", Quantity: 10, Price: 100},
			Status:  "pending",
		}

		sellerProduct := &models.Product{ID: 2, Name: "product_name", Quantity: 12, Price: 100}

		mockOrderRepo.On("GetOrder", mock.Anything).Return(
			*buyerOrder, nil,
		)

		mockProductRepo.On("GetProductByNameAndUser", mock.Anything, mock.Anything).Return(
			*sellerProduct, nil,
		)
		mockProductRepo.On("UpdateProductQuantity", mock.Anything, mock.Anything).Return(
			nil,
		)

		mockOrderRepo.On("AcceptOrder", mock.Anything).Return(
			nil,
		)

		orderService := New(mockOrderRepo, mockProductRepo)

		err := orderService.AcceptOrder(1, 2)
		assert.Equal(t, err, nil)
	})
	t.Run("successfully Accept order seller inventory = buyer order", func(t *testing.T) {
		mockOrderRepo := new(repoMocks.IOrderRepository)
		mockProductRepo := new(repoMocks.IProductRepository)
		buyerOrder := dtos.OrderResponse{
			ID:      1,
			Buyer:   dtos.UserResponse{ID: 1, Name: "buyer_name", Email: "buyer@email.com"},
			Seller:  dtos.UserResponse{ID: 2, Name: "seller_name", Email: "seller@email.com"},
			Product: dtos.ProductResponse{ID: 1, Name: "product_name", Quantity: 10, Price: 100},
			Status:  "pending",
		}

		sellerProduct := models.Product{ID: 2, Name: "product_name", Quantity: 10, Price: 100}

		mockOrderRepo.On("GetOrder", mock.Anything).Return(
			buyerOrder, nil,
		)

		mockProductRepo.On("GetProductByNameAndUser", mock.Anything, mock.Anything).Return(
			sellerProduct, nil,
		)
		mockProductRepo.On("InvalidateProduct", mock.Anything).Return(
			nil,
		)

		mockOrderRepo.On("AcceptOrder", mock.Anything).Return(
			nil,
		)

		orderService := New(mockOrderRepo, mockProductRepo)

		err := orderService.AcceptOrder(1, 2)
		assert.Equal(t, err, nil)
	})

	t.Run("unable to Accept order seller inventory < buyer order", func(t *testing.T) {
		mockOrderRepo := new(repoMocks.IOrderRepository)
		mockProductRepo := new(repoMocks.IProductRepository)
		buyerOrder := dtos.OrderResponse{
			ID:      1,
			Buyer:   dtos.UserResponse{ID: 1, Name: "buyer_name", Email: "buyer@email.com"},
			Seller:  dtos.UserResponse{ID: 2, Name: "seller_name", Email: "seller@email.com"},
			Product: dtos.ProductResponse{ID: 1, Name: "product_name", Quantity: 12, Price: 100},
			Status:  "pending",
		}

		sellerProduct := models.Product{ID: 2, Name: "product_name", Quantity: 10, Price: 100}

		mockOrderRepo.On("GetOrder", mock.Anything).Return(
			buyerOrder, nil,
		)

		mockProductRepo.On("GetProductByNameAndUser", mock.Anything, mock.Anything).Return(
			sellerProduct, nil,
		)

		expectedErr := errors.New("not enough inventory")

		orderService := New(mockOrderRepo, mockProductRepo)

		err := orderService.AcceptOrder(1, 2)
		assert.Equal(t, err, expectedErr)
	})

	t.Run("unable to Accept order: sever error", func(t *testing.T) {
		mockOrderRepo := new(repoMocks.IOrderRepository)
		mockProductRepo := new(repoMocks.IProductRepository)
		buyerOrder := &dtos.OrderResponse{
			ID:      1,
			Buyer:   dtos.UserResponse{ID: 1, Name: "buyer_name", Email: "buyer@email.com"},
			Seller:  dtos.UserResponse{ID: 2, Name: "seller_name", Email: "seller@email.com"},
			Product: dtos.ProductResponse{ID: 1, Name: "product_name", Quantity: 12, Price: 100},
			Status:  "pending",
		}
		mockOrderRepo.On("GetOrder", mock.Anything).Return(
			*buyerOrder, errors.New("Server Error"),
		)

		orderService := New(mockOrderRepo, mockProductRepo)

		err := orderService.AcceptOrder(1, 2)
		assert.NotEqual(t, err, nil)
	})
}

func TestRejectOrder(t *testing.T) {
	t.Run("successfully reject order", func(t *testing.T) {
		mockOrderRepo := new(repoMocks.IOrderRepository)
		mockProductRepo := new(repoMocks.IProductRepository)

		mockOrderRepo.On("RejectOrder", mock.Anything).Return(
			nil,
		)
		orderService := New(mockOrderRepo, mockProductRepo)

		err := orderService.RejectOrder(1)
		assert.Equal(t, err, nil)
	})

	t.Run("unable to reject order", func(t *testing.T) {
		mockOrderRepo := new(repoMocks.IOrderRepository)
		mockProductRepo := new(repoMocks.IProductRepository)

		mockOrderRepo.On("RejectOrder", mock.Anything).Return(
			errors.New("Server Error"),
		)
		orderService := New(mockOrderRepo, mockProductRepo)

		err := orderService.RejectOrder(1)
		assert.NotEqual(t, err, nil)
	})
}

func TestBuyerOrders(t *testing.T) {
	t.Run("successfully get buyer orders", func(t *testing.T) {
		mockOrderRepo := new(repoMocks.IOrderRepository)
		mockProductRepo := new(repoMocks.IProductRepository)

		expectedBuyerOrders := []dtos.BuyerOrdersResponse{}

		mockOrderRepo.On("BuyerOrders", mock.Anything).Return(
			expectedBuyerOrders, nil,
		)
		orderService := New(mockOrderRepo, mockProductRepo)

		buyerOrders, err := orderService.BuyerOrders(1)
		assert.Equal(t, expectedBuyerOrders, buyerOrders)
		assert.Equal(t, err, nil)
	})

	t.Run("unable to get buyer orders", func(t *testing.T) {
		mockOrderRepo := new(repoMocks.IOrderRepository)
		mockProductRepo := new(repoMocks.IProductRepository)

		expectedBuyerOrders := []dtos.BuyerOrdersResponse{}

		mockOrderRepo.On("BuyerOrders", mock.Anything).Return(
			nil, errors.New("server error"),
		)
		orderService := New(mockOrderRepo, mockProductRepo)

		buyerOrders, err := orderService.BuyerOrders(1)
		assert.NotEqual(t, expectedBuyerOrders, buyerOrders)
		assert.NotEqual(t, err, nil)
	})
}

func TestSellerOrders(t *testing.T) {
	t.Run("successfully get seller orders", func(t *testing.T) {
		mockOrderRepo := new(repoMocks.IOrderRepository)
		mockProductRepo := new(repoMocks.IProductRepository)

		expectedSellerOrders := []dtos.SellerOrdersResponse{}

		mockOrderRepo.On("SellerOrders", mock.Anything).Return(
			expectedSellerOrders, nil,
		)
		orderService := New(mockOrderRepo, mockProductRepo)

		sellerOrders, err := orderService.SellerOrders(1)
		assert.Equal(t, expectedSellerOrders, sellerOrders)
		assert.Equal(t, err, nil)
	})

	t.Run("unable to get seller orders", func(t *testing.T) {
		mockOrderRepo := new(repoMocks.IOrderRepository)
		mockProductRepo := new(repoMocks.IProductRepository)

		expectedSellerOrders := []dtos.SellerOrdersResponse{}

		mockOrderRepo.On("SellerOrders", mock.Anything).Return(
			nil, errors.New("server error"),
		)
		orderService := New(mockOrderRepo, mockProductRepo)

		sellerOrders, err := orderService.SellerOrders(1)
		assert.NotEqual(t, expectedSellerOrders, sellerOrders)
		assert.NotEqual(t, err, nil)
	})
}
