package user

import (
	"errors"
	"marketplace/mocks/repoMocks"
	"marketplace/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPlaceBuyerOrder(t *testing.T) {
	t.Run("successfully create user", func(t *testing.T) {
		mockUserRepo := new(repoMocks.IUserRepository)
		mockProductRepo := new(repoMocks.IProductRepository)
		expectedID := int64(1)

		mockUserRepo.On("CreateUser", mock.Anything, mock.Anything, mock.Anything).Return(
			expectedID, nil,
		)
		userService := New(mockUserRepo, mockProductRepo)

		ID, err := userService.CreateUser("user_name", "user_email", "seller")
		assert.Equal(t, ID, expectedID)
		assert.Equal(t, err, nil)
	})
	t.Run("error creating user", func(t *testing.T) {
		mockUserRepo := new(repoMocks.IUserRepository)
		mockProductRepo := new(repoMocks.IProductRepository)
		expectedID := int64(-1)

		mockUserRepo.On("CreateUser", mock.Anything, mock.Anything, mock.Anything).Return(
			expectedID, errors.New("server Error"),
		)
		userService := New(mockUserRepo, mockProductRepo)

		ID, err := userService.CreateUser("user_name", "user_email", "seller")
		assert.Equal(t, ID, expectedID)
		assert.NotEqual(t, err, nil)
	})
}

func TestMatchSellers(t *testing.T) {
	t.Run("successfully create user", func(t *testing.T) {
		mockUserRepo := new(repoMocks.IUserRepository)
		mockProductRepo := new(repoMocks.IProductRepository)
		expectedProduct := models.Product{}
		expectedSellers := []models.User{
			{ID: 1, Name: "user_name", Email: "user_email"},
		}
		mockProductRepo.On("GetProduct", mock.Anything, mock.Anything, mock.Anything).Return(
			expectedProduct, nil,
		)

		mockUserRepo.On("FindSellers", mock.Anything, mock.Anything, mock.Anything).Return(
			expectedSellers, nil,
		)
		userService := New(mockUserRepo, mockProductRepo)

		sellers, err := userService.MatchSellers(1)
		assert.Equal(t, sellers, expectedSellers)
		assert.Equal(t, err, nil)
	})
	t.Run("error creating user", func(t *testing.T) {
		mockUserRepo := new(repoMocks.IUserRepository)
		mockProductRepo := new(repoMocks.IProductRepository)
		expectedProduct := models.Product{}

		mockProductRepo.On("GetProduct", mock.Anything, mock.Anything, mock.Anything).Return(
			expectedProduct, errors.New("server error"),
		)

		userService := New(mockUserRepo, mockProductRepo)

		_, err := userService.MatchSellers(1)
		assert.NotEqual(t, err, nil)
	})
}
