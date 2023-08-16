package controllers

import (
	"bytes"
	"errors"
	"marketplace/mocks/serviceMocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestAddBuyerRequirementController(t *testing.T) {

	t.Run("Successfully Add Buyer Requirement", func(t *testing.T) {
		productServiceMock := new(serviceMocks.IProductService)

		productServiceMock.On("AddBuyerRequirements", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
			int64(1), nil,
		)

		productController := NewProductController(productServiceMock)

		newBuyerReq := `{"name": "product_name", "quantity": 2, "price": 10, "user_id": 1}`

		req, err := http.NewRequest("POST", "/buyer/add-requirement", bytes.NewBufferString(newBuyerReq))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		productController.AddBuyerRequirementController(httpInput, req)

		if httpInput.Code != http.StatusCreated {
			t.Errorf("expected status 201; got %v", httpInput.Code)
		}

		expectedResponse := `{id: 1, message: SuccessFully added Requirement }`

		if httpInput.Body.String() != expectedResponse {
			t.Errorf("expected response %s; got %s", expectedResponse, httpInput.Body.String())
		}
	})

	t.Run("Invalid Request", func(t *testing.T) {
		productServiceMock := new(serviceMocks.IProductService)

		productServiceMock.On("AddBuyerRequirements", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
			int64(1), nil,
		)

		productController := NewProductController(productServiceMock)

		newBuyerReq := `{"name": "product_name", "quantity": 2}`

		req, err := http.NewRequest("POST", "/buyer/add-requirement", bytes.NewBufferString(newBuyerReq))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		productController.AddBuyerRequirementController(httpInput, req)

		if httpInput.Code != http.StatusBadRequest {
			t.Errorf("expected status 400; got %v", httpInput.Code)
		}
	})

	t.Run("Server Error", func(t *testing.T) {
		productServiceMock := new(serviceMocks.IProductService)

		productServiceMock.On("AddBuyerRequirements", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
			int64(-1), errors.New("server error"),
		)

		productController := NewProductController(productServiceMock)

		newBuyerReq := `{"name": "product_name", "quantity": 2, "price": 10, "user_id": 1}`

		req, err := http.NewRequest("POST", "/buyer/add-requirement", bytes.NewBufferString(newBuyerReq))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		productController.AddBuyerRequirementController(httpInput, req)

		if httpInput.Code != http.StatusInternalServerError {
			t.Errorf("expected status 500; got %v", httpInput.Code)
		}
	})
}

func TestAddSellerProductController(t *testing.T) {
	t.Run("Successfully Add Seller Product", func(t *testing.T) {
		productServiceMock := new(serviceMocks.IProductService)

		productServiceMock.On("AddSellerProducts", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
			int64(1), nil,
		)

		productController := NewProductController(productServiceMock)

		newSellerProduct := `{"name": "product_name", "quantity": 2, "price": 10, "user_id": 1}`

		req, err := http.NewRequest("POST", "/seller/add-product", bytes.NewBufferString(newSellerProduct))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		productController.AddSellerProductController(httpInput, req)

		if httpInput.Code != http.StatusCreated {
			t.Errorf("expected status 201; got %v", httpInput.Code)
		}

		expectedResponse := `{id: 1, message: SuccessFully added Product }`

		if httpInput.Body.String() != expectedResponse {
			t.Errorf("expected response %s; got %s", expectedResponse, httpInput.Body.String())
		}
	})

	t.Run("Invalid Request", func(t *testing.T) {
		productServiceMock := new(serviceMocks.IProductService)

		productServiceMock.On("AddSellerProducts", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
			int64(1), nil,
		)

		productController := NewProductController(productServiceMock)

		newSellerProduct := `{"name": "product_name", "quantity": 2}`

		req, err := http.NewRequest("POST", "/seller/add-product", bytes.NewBufferString(newSellerProduct))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		productController.AddSellerProductController(httpInput, req)

		if httpInput.Code != http.StatusBadRequest {
			t.Errorf("expected status 400; got %v", httpInput.Code)
		}
	})

	t.Run("Server Error", func(t *testing.T) {
		productServiceMock := new(serviceMocks.IProductService)

		productServiceMock.On("AddSellerProducts", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
			int64(-1), errors.New("server error"),
		)

		productController := NewProductController(productServiceMock)

		newSellerProduct := `{"name": "product_name", "quantity": 2, "price": 10, "user_id": 1}`

		req, err := http.NewRequest("POST", "/seller/add-product", bytes.NewBufferString(newSellerProduct))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()
		productController.AddSellerProductController(httpInput, req)

		if httpInput.Code != http.StatusInternalServerError {
			t.Errorf("expected status 500; got %v", httpInput.Code)
		}
	})
}
