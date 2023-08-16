package product

import (
	"database/sql"
	"errors"
	"marketplace/mocks/repoMocks"
	"marketplace/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddBuyerRequirements(t *testing.T) {
	t.Run("successfully Add buyer requirement", func(t *testing.T) {
		mockProductRepo := new(repoMocks.IProductRepository)
		expectedProduct := models.Product{}
		mockProductRepo.On("GetProductByNameAndUser", mock.Anything, mock.Anything).Return(
			expectedProduct, sql.ErrNoRows,
		)
		expectedID := int64(1)
		mockProductRepo.On("AddProduct", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
			expectedID, nil,
		)
		productService := New(mockProductRepo)

		ID, err := productService.AddBuyerRequirements("product_name", 10, 100, 1)
		assert.Equal(t, ID, expectedID)
		assert.Equal(t, err, nil)
	})

	t.Run("unable Add buyer requirement, product already exists", func(t *testing.T) {
		mockProductRepo := new(repoMocks.IProductRepository)
		expectedProduct := models.Product{}
		mockProductRepo.On("GetProductByNameAndUser", mock.Anything, mock.Anything).Return(
			expectedProduct, nil,
		)
		productService := New(mockProductRepo)

		expectedError := errors.New("product already exists")

		_, err := productService.AddBuyerRequirements("product_name", 10, 100, 1)
		assert.Equal(t, err, expectedError)
	})

	t.Run("unable Add buyer requirement, server error", func(t *testing.T) {
		mockProductRepo := new(repoMocks.IProductRepository)
		expectedProduct := models.Product{}
		mockProductRepo.On("GetProductByNameAndUser", mock.Anything, mock.Anything).Return(
			expectedProduct, sql.ErrNoRows,
		)
		expectedID := int64(1)
		mockProductRepo.On("AddProduct", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
			expectedID, errors.New("server error"),
		)
		productService := New(mockProductRepo)

		_, err := productService.AddBuyerRequirements("product_name", 10, 100, 1)
		assert.NotEqual(t, err, nil)
	})
}

func TestAddSellerProducts(t *testing.T) {
	t.Run("successfully Add seller products", func(t *testing.T) {
		mockProductRepo := new(repoMocks.IProductRepository)
		expectedProduct := models.Product{}
		mockProductRepo.On("GetProductByNameAndUser", mock.Anything, mock.Anything).Return(
			expectedProduct, sql.ErrNoRows,
		)
		expectedID := int64(1)
		mockProductRepo.On("AddProduct", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
			expectedID, nil,
		)
		productService := New(mockProductRepo)

		ID, err := productService.AddSellerProducts("product_name", 10, 100, 1)
		assert.Equal(t, ID, expectedID)
		assert.Equal(t, err, nil)
	})

	t.Run("unable Add seller products, product already exists", func(t *testing.T) {
		mockProductRepo := new(repoMocks.IProductRepository)
		expectedProduct := models.Product{}
		mockProductRepo.On("GetProductByNameAndUser", mock.Anything, mock.Anything).Return(
			expectedProduct, nil,
		)
		productService := New(mockProductRepo)

		expectedError := errors.New("product already exists")

		_, err := productService.AddSellerProducts("product_name", 10, 100, 1)
		assert.Equal(t, err, expectedError)
	})

	t.Run("unable Add seller products, server error", func(t *testing.T) {
		mockProductRepo := new(repoMocks.IProductRepository)
		expectedProduct := models.Product{}
		mockProductRepo.On("GetProductByNameAndUser", mock.Anything, mock.Anything).Return(
			expectedProduct, sql.ErrNoRows,
		)
		expectedID := int64(1)
		mockProductRepo.On("AddProduct", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
			expectedID, errors.New("server error"),
		)
		productService := New(mockProductRepo)

		_, err := productService.AddSellerProducts("product_name", 10, 100, 1)
		assert.NotEqual(t, err, nil)
	})
}
