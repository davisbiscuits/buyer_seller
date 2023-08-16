package user

import (
	"marketplace/models"
	"marketplace/repositories"
	"marketplace/services"
)

type Service struct {
	services.IUserService
	userRepository    repositories.IUserRepository
	productRepository repositories.IProductRepository
}

func New(userRepository repositories.IUserRepository, productRepository repositories.IProductRepository) services.IUserService {
	return &Service{
		userRepository:    userRepository,
		productRepository: productRepository,
	}
}

// create user both seller or buyer , diffrentiated by type
func (s *Service) CreateUser(userType string, userEmail string, userName string) (int64, error) {
	user_id, err := s.userRepository.CreateUser(userType, userEmail, userName)
	if err != nil {
		return -1, err
	}
	return user_id, nil
}

// get seller for a product
func (s *Service) MatchSellers(productID int) ([]models.User, error) {
	product, err := s.productRepository.GetProduct(productID)
	if err != nil {
		return nil, err
	}

	users, err := s.userRepository.FindSellers(product.Name, product.Quantity, product.Price)
	if err != nil {
		return nil, err
	}
	return users, nil
}
