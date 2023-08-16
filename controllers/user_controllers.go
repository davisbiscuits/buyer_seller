package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"marketplace/dtos"
	"marketplace/models"
	"marketplace/services"
	"net/http"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) RegisterUserController(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = user.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := c.userService.CreateUser(user.Type, user.Email, user.Name)

	if err != nil {
		http.Error(w, "Unable to Process: Server Error", http.StatusInternalServerError)
		return
	}

	jsonData := fmt.Sprintf("{id: %d, message: %s }", id, "SuccessFully added User")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(jsonData))
}

func (c *UserController) GetSellersByProductController(w http.ResponseWriter, r *http.Request) {

	var product dtos.ProductRequest
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

	users, err := c.userService.MatchSellers(product.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No Sellers found", http.StatusNotFound)
			return
		}
		http.Error(w, "Unable to Process: Server Error", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Unable to Process: Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
