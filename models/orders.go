package models

type Order struct {
	SellerID  int    `json:"seller_id" binding:"required"`
	BuyerID   int    `json:"buyer_id" binding:"required"`
	ProductID int    `json:"product_id" binding:"required"`
	Status    string `json:"status" binding:"required"`
}
