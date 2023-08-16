package dtos

import "errors"

type ProductResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type ProductRequest struct {
	ID int `json:"product_id"`
}

type ProductAddition struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	UserID   int    `json:"user_id"`
}

func (p *ProductRequest) Validate() error {
	if p.ID == 0 {
		return errors.New("product_id is required")
	}
	return nil
}

func (p *ProductAddition) Validate() error {
	if p.Name == "" {
		return errors.New("name is required")
	}
	if p.Price == 0 {
		return errors.New("price is required")
	}
	if p.Quantity == 0 {
		return errors.New("qunatity is required")
	}
	if p.UserID == 0 {
		return errors.New("user_id is required")
	}
	return nil
}
