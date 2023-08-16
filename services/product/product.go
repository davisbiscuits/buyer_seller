package product

import (
	"database/sql"
	"errors"
	"marketplace/repositories"
	"marketplace/services"
)

type Service struct {
	services.IProductService
	productRepository repositories.IProductRepository
}

func New(productRepository repositories.IProductRepository) services.IProductService {
	return &Service{
		productRepository: productRepository,
	}
}

func (s *Service) AddBuyerRequirements(productName string, productQuantity int, productPrice int, buyerID int) (int64, error) {
	_, err := s.productRepository.GetProductByNameAndUser(productName, buyerID)
	if err != sql.ErrNoRows {
		return -1, errors.New("product already exists")
	}

	productID, err := s.productRepository.AddProduct(productName, productQuantity, productPrice, buyerID)
	if err != nil {
		return -1, err
	}
	return productID, nil
}

func (s *Service) AddSellerProducts(productName string, productQuantity int, productPrice int, sellerID int) (int64, error) {

	_, err := s.productRepository.GetProductByNameAndUser(productName, sellerID)
	if err != sql.ErrNoRows {
		return -1, errors.New("product already exists")
	}
	productID, err := s.productRepository.AddProduct(productName, productQuantity, productPrice, sellerID)
	if err != nil {
		return -1, err
	}
	return productID, nil
}
