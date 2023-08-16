package controllers

import (
	"encoding/json"
	"fmt"
	"marketplace/dtos"
	"marketplace/services"
	"net/http"
)

type ProductController struct {
	productService services.IProductService
}

func NewProductController(productService services.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (c *ProductController) AddBuyerRequirementController(w http.ResponseWriter, r *http.Request) {
	var product dtos.ProductAddition
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	err = product.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := c.productService.AddBuyerRequirements(product.Name, product.Quantity, product.Price, product.UserID)
	if err != nil {
		http.Error(w, "Unable to process", http.StatusInternalServerError)
		return
	}
	jsonData := fmt.Sprintf("{id: %d, message: %s }", id, "SuccessFully added Requirement")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(jsonData))
}

func (c *ProductController) AddSellerProductController(w http.ResponseWriter, r *http.Request) {
	var product dtos.ProductAddition
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = product.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := c.productService.AddSellerProducts(product.Name, product.Quantity, product.Price, product.UserID)
	if err != nil {
		http.Error(w, "Unable to process", http.StatusInternalServerError)
		return
	}
	jsonData := fmt.Sprintf("{id: %d, message: %s }", id, "SuccessFully added Product")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(jsonData))
}
