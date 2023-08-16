package services

import (
	"marketplace/dtos"
	"marketplace/models"
)

// IEntityProfileService is the interface for the EntityProfile service
type IUserService interface {
	CreateUser(userType string, userEmail string, userName string) (int64, error)
	MatchSellers(productID int) ([]models.User, error)
}

type IOrderService interface {
	PlaceBuyerOrder(productID int, sellerID int, buyerID int) (int64, error)
	RejectOrder(orderID int) error
	AcceptOrder(orderID int, sellerID int) error
	BuyerOrders(buyerID int) ([]dtos.BuyerOrdersResponse, error)
	SellerOrders(sellerID int) ([]dtos.SellerOrdersResponse, error)
}

type IProductService interface {
	AddBuyerRequirements(productName string, productQuantity int, productPrice int, buyerID int) (int64, error)
	AddSellerProducts(productName string, productQuantity int, productPrice int, sellerID int) (int64, error)
}
