package models

type Product struct {
	ID       int    `json:"id" `
	Name     string `json:"name" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
	Price    int    `json:"price" binding:"required"`
	Valid    bool   `json:"valid"`
	UserID   int    `json:"user_id" binding:"required"`
}
