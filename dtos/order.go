package dtos

import "errors"

type BuyerOrdersResponse struct {
	ID          int             `json:"id"`
	Seller      UserResponse    `json:"seller"`
	Requirement ProductResponse `json:"requirement"`
	Status      string          `json:"status"`
}

type SellerOrdersResponse struct {
	ID      int             `json:"id"`
	Buyer   UserResponse    `json:"buyer"`
	Product ProductResponse `json:"product"`
	Status  string          `json:"status"`
}

type OrderAddition struct {
	SellerID  int `json:"seller_id"`
	BuyerID   int `json:"buyer_id"`
	ProductID int `json:"product_id"`
}
type OrdersAccept struct {
	ID     int `json:"order_id"`
	UserID int `json:"user_id"`
}

type OrdersReject struct {
	ID int `json:"order_id"`
}

type OrderResponse struct {
	ID      int
	Buyer   UserResponse
	Seller  UserResponse
	Product ProductResponse
	Status  string
}

func (o *OrderAddition) Validate() error {
	if o.SellerID == 0 {
		return errors.New("seller_id is required")
	}
	if o.BuyerID == 0 {
		return errors.New("buyer_id is required")
	}
	if o.ProductID == 0 {
		return errors.New("product_id is required")
	}
	return nil
}

func (o *OrdersAccept) Validate() error {
	if o.ID == 0 {
		return errors.New("order_id is required")
	}
	if o.UserID == 0 {
		return errors.New("user_id is required")
	}
	return nil
}

func (o *OrdersReject) Validate() error {
	if o.ID == 0 {
		return errors.New("order_id is required")
	}
	return nil
}
