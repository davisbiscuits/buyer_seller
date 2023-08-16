package repositories

import (
	"marketplace/dtos"
	"marketplace/models"
)

type IProductRepository interface {
	AddProduct(name string, quantity int, price int, userID int) (int64, error)
	GetProduct(productID int) (models.Product, error)
	BuyerProducts(buyerID int) ([]models.Product, error)
	SellerProducts(sellerID int) ([]models.Product, error)
	UpdateProductQuantity(quantity int, productID int) error
	InvalidateProduct(productID int) error
	GetProductByNameAndUser(productName string, userID int) (models.Product, error)
}

type IUserRepository interface {
	CreateUser(userType string, userEmail string, userName string) (int64, error)
	FindSellers(productName string, productQuantity int, productPrice int) ([]models.User, error)
}

type IOrderRepository interface {
	AddOrder(productID int, sellerID int, buyerID int) (int64, error)
	GetOrder(orderID int) (dtos.OrderResponse, error)
	BuyerOrders(buyerID int) ([]dtos.BuyerOrdersResponse, error)
	SellerOrders(buyerID int) ([]dtos.SellerOrdersResponse, error)
	AcceptOrder(orderID int) error
	RejectOrder(orderID int) error
}
