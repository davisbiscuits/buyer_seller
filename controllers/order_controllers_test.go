package controllers

import (
	"bytes"
	"database/sql"
	"errors"
	"marketplace/dtos"
	"marketplace/mocks/serviceMocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestPlaceBuyerOrder(t *testing.T) {
	t.Run("Successfully Place Buyer Order", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("PlaceBuyerOrder", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
			int64(1), nil,
		)

		orderController := NewOrderController(orderServiceMock)

		newOrder := `{"seller_id": 1, "buyer_id": 2, "product_id": 10}`

		req, err := http.NewRequest("POST", "/buyer/place-order", bytes.NewBufferString(newOrder))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.PlaceBuyerOrder(httpInput, req)

		if httpInput.Code != http.StatusCreated {
			t.Errorf("expected status 201; got %v", httpInput.Code)
		}

		expectedResponse := `{id: 1, message: SuccessFully placed Order }`

		if httpInput.Body.String() != expectedResponse {
			t.Errorf("expected response %s; got %s", expectedResponse, httpInput.Body.String())
		}
	})

	t.Run("Invalid Request", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("PlaceBuyerOrder", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
			int64(1), nil,
		)

		orderController := NewOrderController(orderServiceMock)

		newOrder := `{"seller_id": 1, "buyer_id": 2}`

		req, err := http.NewRequest("POST", "/buyer/place-order", bytes.NewBufferString(newOrder))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.PlaceBuyerOrder(httpInput, req)

		if httpInput.Code != http.StatusBadRequest {
			t.Errorf("expected status 400; got %v", httpInput.Code)
		}
	})

	t.Run("Server Error", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("PlaceBuyerOrder", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
			int64(1), errors.New("server error"),
		)

		orderController := NewOrderController(orderServiceMock)

		newOrder := `{"seller_id": 1, "buyer_id": 2, "product_id": 10}`

		req, err := http.NewRequest("POST", "/buyer/place-order", bytes.NewBufferString(newOrder))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.PlaceBuyerOrder(httpInput, req)

		if httpInput.Code != http.StatusInternalServerError {
			t.Errorf("expected status 500; got %v", httpInput.Code)
		}
	})
}

func TestAcceptOrder(t *testing.T) {
	t.Run("Successfully Accept Order", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("AcceptOrder", mock.Anything, mock.Anything).Return(
			nil,
		)

		orderController := NewOrderController(orderServiceMock)

		newOrder := `{"user_id": 1, "order_id": 2}`

		req, err := http.NewRequest("POST", "/seller/accept-order", bytes.NewBufferString(newOrder))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.AcceptOrder(httpInput, req)

		if httpInput.Code != http.StatusOK {
			t.Errorf("expected status 200; got %v", httpInput.Code)
		}

		expectedResponse := `{message: SuccessFully accepted Order }`

		if httpInput.Body.String() != expectedResponse {
			t.Errorf("expected response %s; got %s", expectedResponse, httpInput.Body.String())
		}
	})

	t.Run("Order not found", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("AcceptOrder", mock.Anything, mock.Anything).Return(
			sql.ErrNoRows,
		)

		orderController := NewOrderController(orderServiceMock)

		newOrder := `{"user_id": 1, "order_id": 2}`

		req, err := http.NewRequest("POST", "/seller/accept-order", bytes.NewBufferString(newOrder))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.AcceptOrder(httpInput, req)

		if httpInput.Code != http.StatusNotFound {
			t.Errorf("expected status 404; got %v", httpInput.Code)
		}

	})

	t.Run("Invalid Request", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("AcceptOrder", mock.Anything, mock.Anything).Return(
			nil,
		)

		orderController := NewOrderController(orderServiceMock)

		newOrder := `{"user_id": 1}`

		req, err := http.NewRequest("POST", "/seller/accept-order", bytes.NewBufferString(newOrder))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.AcceptOrder(httpInput, req)

		if httpInput.Code != http.StatusBadRequest {
			t.Errorf("expected status 400; got %v", httpInput.Code)
		}
	})

	t.Run("Server Error", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("AcceptOrder", mock.Anything, mock.Anything).Return(
			errors.New("server error"),
		)

		orderController := NewOrderController(orderServiceMock)

		newOrder := `{"user_id": 1, "order_id": 2}`

		req, err := http.NewRequest("POST", "/seller/accept-order", bytes.NewBufferString(newOrder))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.AcceptOrder(httpInput, req)

		if httpInput.Code != http.StatusInternalServerError {
			t.Errorf("expected status 500; got %v", httpInput.Code)
		}
	})
}

func TestRejectOrder(t *testing.T) {
	t.Run("Successfully Reject Order", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("RejectOrder", mock.Anything, mock.Anything).Return(
			nil,
		)

		orderController := NewOrderController(orderServiceMock)

		newOrder := `{"user_id": 1, "order_id": 2}`

		req, err := http.NewRequest("POST", "/seller/reject-order", bytes.NewBufferString(newOrder))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.RejectOrder(httpInput, req)

		if httpInput.Code != http.StatusOK {
			t.Errorf("expected status 200; got %v", httpInput.Code)
		}

		expectedResponse := `{message: Successfully rejected Order }`

		if httpInput.Body.String() != expectedResponse {
			t.Errorf("expected response %s; got %s", expectedResponse, httpInput.Body.String())
		}
	})

	t.Run("Order not found", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("RejectOrder", mock.Anything, mock.Anything).Return(
			sql.ErrNoRows,
		)

		orderController := NewOrderController(orderServiceMock)

		newOrder := `{"user_id": 1, "order_id": 2}`

		req, err := http.NewRequest("POST", "/seller/reject-order", bytes.NewBufferString(newOrder))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.RejectOrder(httpInput, req)

		if httpInput.Code != http.StatusNotFound {
			t.Errorf("expected status 404; got %v", httpInput.Code)
		}

	})

	t.Run("Invalid Request", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("RejectOrder", mock.Anything, mock.Anything).Return(
			nil,
		)

		orderController := NewOrderController(orderServiceMock)

		newOrder := `{"user_id": 1}`

		req, err := http.NewRequest("POST", "/seller/reject-order", bytes.NewBufferString(newOrder))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.RejectOrder(httpInput, req)

		if httpInput.Code != http.StatusBadRequest {
			t.Errorf("expected status 400; got %v", httpInput.Code)
		}
	})

	t.Run("Server Error", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("RejectOrder", mock.Anything, mock.Anything).Return(
			errors.New("server error"),
		)

		orderController := NewOrderController(orderServiceMock)

		newOrder := `{"user_id": 1, "order_id": 2}`

		req, err := http.NewRequest("POST", "/seller/reject-order", bytes.NewBufferString(newOrder))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.RejectOrder(httpInput, req)

		if httpInput.Code != http.StatusInternalServerError {
			t.Errorf("expected status 500; got %v", httpInput.Code)
		}
	})
}

func TestGetOrdersForBuyer(t *testing.T) {
	t.Run("Successfully Get Buyer Orders", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("BuyerOrders", mock.Anything).Return(
			[]dtos.BuyerOrdersResponse{
				{ID: 1,
					Seller:      dtos.UserResponse{ID: 1, Name: "seller_name", Email: "seller@email.com"},
					Requirement: dtos.ProductResponse{ID: 2, Name: "req_name", Quantity: 10, Price: 10},
					Status:      "pending"},
			}, nil,
		)

		orderController := NewOrderController(orderServiceMock)

		user := `{"user_id": 1}`

		req, err := http.NewRequest("POST", "/buyer/orders", bytes.NewBufferString(user))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.GetOrdersForBuyer(httpInput, req)

		if httpInput.Code != http.StatusOK {
			t.Errorf("expected status 200; got %v", httpInput.Code)
		}

		expectedResponse := `[{"id":1,"seller":{"id":1,"name":"seller_name","email":"seller@email.com"},"requirement":{"id":2,"name":"req_name","quantity":10,"price":10},"status":"pending"}]`

		if httpInput.Body.String() != expectedResponse {
			t.Errorf("expected response %s; got %s", expectedResponse, httpInput.Body.String())
		}
	})

	t.Run("Orders not found", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("BuyerOrders", mock.Anything).Return(
			nil, sql.ErrNoRows,
		)

		orderController := NewOrderController(orderServiceMock)

		user := `{"user_id": 1}`

		req, err := http.NewRequest("POST", "/buyer/orders", bytes.NewBufferString(user))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.GetOrdersForBuyer(httpInput, req)

		if httpInput.Code != http.StatusNotFound {
			t.Errorf("expected status 404; got %v", httpInput.Code)
		}

	})

	t.Run("Invalid Request", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("BuyerOrders", mock.Anything).Return(
			[]dtos.BuyerOrdersResponse{
				{ID: 1,
					Seller:      dtos.UserResponse{ID: 1, Name: "seller_name", Email: "seller@email.com"},
					Requirement: dtos.ProductResponse{ID: 2, Name: "req_name", Quantity: 10, Price: 10},
					Status:      "pending"},
			}, nil,
		)

		orderController := NewOrderController(orderServiceMock)

		user := `{}`

		req, err := http.NewRequest("POST", "/buyer/orders", bytes.NewBufferString(user))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.GetOrdersForBuyer(httpInput, req)

		if httpInput.Code != http.StatusBadRequest {
			t.Errorf("expected status 400; got %v", httpInput.Code)
		}
	})

	t.Run("Server Error", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("BuyerOrders", mock.Anything).Return(
			nil, errors.New("server error"),
		)

		orderController := NewOrderController(orderServiceMock)

		user := `{"user_id": 1}`

		req, err := http.NewRequest("POST", "/buyer/orders", bytes.NewBufferString(user))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.GetOrdersForBuyer(httpInput, req)

		if httpInput.Code != http.StatusInternalServerError {
			t.Errorf("expected status 500; got %v", httpInput.Code)
		}
	})
}

func TestGetOrdersForSeller(t *testing.T) {
	t.Run("Successfully Get Buyer Orders", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("SellerOrders", mock.Anything).Return(
			[]dtos.SellerOrdersResponse{
				{ID: 1,
					Buyer:   dtos.UserResponse{ID: 1, Name: "buyer_name", Email: "buyer@email.com"},
					Product: dtos.ProductResponse{ID: 2, Name: "roduct_name", Quantity: 10, Price: 10},
					Status:  "pending"},
			}, nil,
		)

		orderController := NewOrderController(orderServiceMock)

		user := `{"user_id": 1}`

		req, err := http.NewRequest("POST", "/seller/orders", bytes.NewBufferString(user))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.GetOrdersForSeller(httpInput, req)

		if httpInput.Code != http.StatusOK {
			t.Errorf("expected status 200; got %v", httpInput.Code)
		}

		expectedResponse := `[{"id":1,"buyer":{"id":1,"name":"buyer_name","email":"buyer@email.com"},"product":{"id":2,"name":"roduct_name","quantity":10,"price":10},"status":"pending"}]`

		if httpInput.Body.String() != expectedResponse {
			t.Errorf("expected response %s; got %s", expectedResponse, httpInput.Body.String())
		}
	})

	t.Run("Orders not found", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("SellerOrders", mock.Anything).Return(
			nil, sql.ErrNoRows,
		)

		orderController := NewOrderController(orderServiceMock)

		user := `{"user_id": 1}`

		req, err := http.NewRequest("POST", "/seller/orders", bytes.NewBufferString(user))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.GetOrdersForSeller(httpInput, req)

		if httpInput.Code != http.StatusNotFound {
			t.Errorf("expected status 404; got %v", httpInput.Code)
		}

	})

	t.Run("Invalid Request", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("SellerOrders", mock.Anything).Return(
			[]dtos.SellerOrdersResponse{
				{ID: 1,
					Buyer:   dtos.UserResponse{ID: 1, Name: "buyer_name", Email: "buyer@email.com"},
					Product: dtos.ProductResponse{ID: 2, Name: "roduct_name", Quantity: 10, Price: 10},
					Status:  "pending"},
			}, nil,
		)

		orderController := NewOrderController(orderServiceMock)

		user := `{}`

		req, err := http.NewRequest("POST", "/seller/orders", bytes.NewBufferString(user))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.GetOrdersForSeller(httpInput, req)

		if httpInput.Code != http.StatusBadRequest {
			t.Errorf("expected status 400; got %v", httpInput.Code)
		}
	})

	t.Run("Server Error", func(t *testing.T) {
		orderServiceMock := new(serviceMocks.IOrderService)

		orderServiceMock.On("SellerOrders", mock.Anything).Return(
			nil, errors.New("server error"),
		)

		orderController := NewOrderController(orderServiceMock)

		user := `{"user_id": 1}`

		req, err := http.NewRequest("POST", "/seller/orders", bytes.NewBufferString(user))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		orderController.GetOrdersForSeller(httpInput, req)

		if httpInput.Code != http.StatusInternalServerError {
			t.Errorf("expected status 500; got %v", httpInput.Code)
		}
	})
}
