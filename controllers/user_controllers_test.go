package controllers

import (
	"bytes"
	"database/sql"
	"errors"
	"marketplace/mocks/serviceMocks"
	"marketplace/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestRegisterUserController(t *testing.T) {

	t.Run("Successfully Register User", func(t *testing.T) {
		userServiceMock := new(serviceMocks.IUserService)
		expectedID := int64(1)

		expectedResponse := `{id: 1, message: SuccessFully added User }`

		userServiceMock.On("CreateUser", mock.Anything, mock.Anything, mock.Anything).Return(
			expectedID, nil,
		)
		userController := NewUserController(userServiceMock)

		newUser := `{"name": "user_name", "email": "user@email.com", "type": "seller"}`

		req, err := http.NewRequest("POST", "/register", bytes.NewBufferString(newUser))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()

		userController.RegisterUserController(httpInput, req)

		if httpInput.Code != http.StatusCreated {
			t.Errorf("expected status 201; got %v", httpInput.Code)
		}

		if httpInput.Body.String() != expectedResponse {
			t.Errorf("expected response %s; got %s", expectedResponse, httpInput.Body.String())
		}
	})

	t.Run("Invalid Request", func(t *testing.T) {
		userServiceMock := new(serviceMocks.IUserService)
		expectedID := int64(1)

		userServiceMock.On("CreateUser", mock.Anything, mock.Anything, mock.Anything).Return(
			expectedID, nil,
		)
		userController := NewUserController(userServiceMock)

		newUser := `{"name": "user_name", "email": "user@email.com"}`

		req, err := http.NewRequest("POST", "/register", bytes.NewBufferString(newUser))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()

		userController.RegisterUserController(httpInput, req)

		if httpInput.Code != http.StatusBadRequest {
			t.Errorf("expected status 400; got %v", httpInput.Code)
		}
	})

	t.Run("Server Error", func(t *testing.T) {
		userServiceMock := new(serviceMocks.IUserService)
		expectedID := int64(-1)

		userServiceMock.On("CreateUser", mock.Anything, mock.Anything, mock.Anything).Return(
			expectedID, errors.New("server error"),
		)
		userController := NewUserController(userServiceMock)

		newUser := `{"name": "user_name", "email": "user@email.com", "type": "seller"}`

		req, err := http.NewRequest("POST", "/register", bytes.NewBufferString(newUser))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()

		userController.RegisterUserController(httpInput, req)

		if httpInput.Code != http.StatusInternalServerError {
			t.Errorf("expected status 500; got %v", httpInput.Code)
		}
	})
}

func TestGetSellersByProductController(t *testing.T) {
	t.Run("Successful Register User", func(t *testing.T) {
		userServiceMock := new(serviceMocks.IUserService)
		users := []models.User{
			{ID: 1, Name: "user_name", Email: "user@email.com", Type: "seller"},
		}

		userServiceMock.On("MatchSellers", mock.Anything).Return(
			users, nil,
		)
		userController := NewUserController(userServiceMock)

		productID := `{"product_id": 1}`

		req, err := http.NewRequest("GET", "/buyer/match", bytes.NewBufferString(productID))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()

		userController.GetSellersByProductController(httpInput, req)

		if httpInput.Code != http.StatusOK {
			t.Errorf("expected status 200; got %v", httpInput.Code)
		}

		expectedResponse := `[{"id":1,"name":"user_name","email":"user@email.com","type":"seller"}]`
		if httpInput.Body.String() != expectedResponse {
			t.Errorf("expected response %s; got %s", expectedResponse, httpInput.Body.String())
		}
	})

	t.Run("Invalid Request", func(t *testing.T) {
		userServiceMock := new(serviceMocks.IUserService)
		userController := NewUserController(userServiceMock)

		productID := ``

		req, err := http.NewRequest("GET", "/buyer/match", bytes.NewBufferString(productID))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()

		userController.GetSellersByProductController(httpInput, req)

		if httpInput.Code != http.StatusBadRequest {
			t.Errorf("expected status 400; got %v", httpInput.Code)
		}
	})

	t.Run("No Sellers found", func(t *testing.T) {
		userServiceMock := new(serviceMocks.IUserService)
		users := []models.User{}

		userServiceMock.On("MatchSellers", mock.Anything).Return(
			users, sql.ErrNoRows,
		)
		userController := NewUserController(userServiceMock)

		productID := `{"product_id": 1}`

		req, err := http.NewRequest("GET", "/buyer/match", bytes.NewBufferString(productID))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()

		userController.GetSellersByProductController(httpInput, req)

		if httpInput.Code != http.StatusNotFound {
			t.Errorf("expected status 404; got %v", httpInput.Code)
		}

	})

	t.Run("Server Error", func(t *testing.T) {
		userServiceMock := new(serviceMocks.IUserService)
		users := []models.User{}

		userServiceMock.On("MatchSellers", mock.Anything).Return(
			users, errors.New("server error"),
		)
		userController := NewUserController(userServiceMock)

		productID := `{"product_id": 1}`

		req, err := http.NewRequest("GET", "/buyer/match", bytes.NewBufferString(productID))
		if err != nil {
			t.Fatal(err)
		}

		httpInput := httptest.NewRecorder()

		userController.GetSellersByProductController(httpInput, req)

		if httpInput.Code != http.StatusInternalServerError {
			t.Errorf("expected status 500; got %v", httpInput.Code)
		}
	})
}
