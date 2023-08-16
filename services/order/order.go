package order

import (
	"errors"
	"marketplace/dtos"
	"marketplace/repositories"
	"marketplace/services"
)

type Service struct {
	services.IOrderService
	orderRepository   repositories.IOrderRepository
	productRepository repositories.IProductRepository
}

func New(orderRepository repositories.IOrderRepository, productRepository repositories.IProductRepository) services.IOrderService {
	return &Service{
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}

func (s *Service) PlaceBuyerOrder(productID int, sellerID int, buyerID int) (int64, error) {
	id, err := s.orderRepository.AddOrder(productID, sellerID, buyerID)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (s *Service) AcceptOrder(orderID int, sellerID int) error {
	// get buyer requirements
	buyerOrder, err := s.orderRepository.GetOrder(orderID)
	if err != nil {
		return err
	}
	// get seller inventory for the product
	sellerProduct, err := s.productRepository.GetProductByNameAndUser(buyerOrder.Product.Name, sellerID)
	if err != nil {
		return err
	}

	// check if enough quantity present
	if buyerOrder.Product.Quantity > sellerProduct.Quantity {
		return errors.New("not enough inventory")
	}

	// update quantity or invalidate incase of quantity being eqaul
	if buyerOrder.Product.Quantity < sellerProduct.Quantity {
		err = s.productRepository.UpdateProductQuantity(sellerProduct.Quantity-buyerOrder.Product.Quantity, sellerProduct.ID)
	} else {
		err = s.productRepository.InvalidateProduct(sellerProduct.ID)
	}

	if err != nil {
		return err
	}

	// update status to accepted
	err = s.orderRepository.AcceptOrder(orderID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) RejectOrder(orderID int) error {
	err := s.orderRepository.RejectOrder(orderID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) BuyerOrders(buyerID int) ([]dtos.BuyerOrdersResponse, error) {
	buyers, err := s.orderRepository.BuyerOrders(buyerID)
	if err != nil {
		return nil, err
	}
	return buyers, nil
}

func (s *Service) SellerOrders(sellerID int) ([]dtos.SellerOrdersResponse, error) {
	sellers, err := s.orderRepository.SellerOrders(sellerID)
	if err != nil {
		return nil, err
	}
	return sellers, nil
}
