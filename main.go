package main

import (
	"fmt"
	"log"
	"marketplace/application"
	"net/http"
)

func main() {

	app := application.App{}
	app.InstantiateApp()

	http.HandleFunc("/register", app.Controllers.User.RegisterUserController)
	http.HandleFunc("/buyer/add-requirement", app.Controllers.Product.AddBuyerRequirementController)
	http.HandleFunc("/buyer/match", app.Controllers.User.GetSellersByProductController)
	http.HandleFunc("/buyer/place-order", app.Controllers.Order.PlaceBuyerOrder)
	http.HandleFunc("/buyer/orders", app.Controllers.Order.GetOrdersForBuyer)
	http.HandleFunc("/seller/add-product", app.Controllers.Product.AddSellerProductController)
	http.HandleFunc("/seller/accept-order", app.Controllers.Order.AcceptOrder)
	http.HandleFunc("/seller/reject-order", app.Controllers.Order.RejectOrder)
	http.HandleFunc("/seller/orders", app.Controllers.Order.GetOrdersForBuyer)

	port := 8080
	log.Printf("Server started at :%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
