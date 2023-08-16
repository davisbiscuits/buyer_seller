package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"marketplace/dtos"
	"marketplace/services"
	"net/http"
)

type OrderController struct {
	orderService services.IOrderService
}

func NewOrderController(orderService services.IOrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (c *OrderController) PlaceBuyerOrder(w http.ResponseWriter, r *http.Request) {
	var order dtos.OrderAddition
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = order.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := c.orderService.PlaceBuyerOrder(order.ProductID, order.SellerID, order.BuyerID)
	if err != nil {
		http.Error(w, "Unable to process", http.StatusInternalServerError)
		return
	}

	jsonData := fmt.Sprintf("{id: %d, message: %s }", id, "SuccessFully placed Order")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(jsonData))
}

func (c *OrderController) AcceptOrder(w http.ResponseWriter, r *http.Request) {
	var order dtos.OrdersAccept
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = order.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.orderService.AcceptOrder(order.ID, order.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Unable to process", http.StatusInternalServerError)
		return
	}

	jsonData := fmt.Sprintf("{message: %s }", "SuccessFully accepted Order")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonData))
}

func (c *OrderController) RejectOrder(w http.ResponseWriter, r *http.Request) {
	var order dtos.OrdersReject
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = order.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.orderService.RejectOrder(order.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Unable to process", http.StatusInternalServerError)
		return
	}

	jsonData := fmt.Sprintf("{message: %s }", "Successfully rejected Order")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonData))
}

func (c *OrderController) GetOrdersForBuyer(w http.ResponseWriter, r *http.Request) {
	var buyer dtos.UserRequest
	err := json.NewDecoder(r.Body).Decode(&buyer)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = buyer.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orders, err := c.orderService.BuyerOrders(buyer.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Orders not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Unable to process", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(orders)
	if err != nil {
		http.Error(w, "Unable to Process: Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonData))
}

func (c *OrderController) GetOrdersForSeller(w http.ResponseWriter, r *http.Request) {
	var seller dtos.UserRequest
	err := json.NewDecoder(r.Body).Decode(&seller)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = seller.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orders, err := c.orderService.SellerOrders(seller.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Orders not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Unable to process", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(orders)
	if err != nil {
		http.Error(w, "Unable to Process: Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonData))
}
