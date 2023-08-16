package application

import (
	"log"
	"marketplace/controllers"
	"marketplace/repositories"
	orderRepo "marketplace/repositories/order"
	productRepo "marketplace/repositories/product"
	userRepo "marketplace/repositories/user"
	"marketplace/services"
	orderService "marketplace/services/order"
	productService "marketplace/services/product"
	userService "marketplace/services/user"
	"marketplace/stores"
	mysqlStore "marketplace/stores/mysqldb"
)

type App struct {
	Store struct {
		MysqlCon stores.Imysqldb
	}

	Repository struct {
		User    repositories.IUserRepository
		Order   repositories.IOrderRepository
		Product repositories.IProductRepository
	}

	Service struct {
		User    services.IUserService
		Order   services.IOrderService
		Product services.IProductService
	}

	Controllers struct {
		User    controllers.UserController
		Order   controllers.OrderController
		Product controllers.ProductController
	}
}

func (a *App) InstantiateStores() {
	a.Store.MysqlCon = mysqlStore.New()
}

func (a *App) InstantiateRepositories() {
	a.Repository.User = userRepo.New(a.Store.MysqlCon)
	a.Repository.Product = productRepo.New(a.Store.MysqlCon)
	a.Repository.Order = orderRepo.New(a.Store.MysqlCon)
}
func (a *App) InstantiateServices() {
	a.Service.User = userService.New(a.Repository.User, a.Repository.Product)
	a.Service.Product = productService.New(a.Repository.Product)
	a.Service.Order = orderService.New(a.Repository.Order, a.Repository.Product)
}

func (a *App) InstantiateControllers() {
	a.Controllers.User = *controllers.NewUserController(a.Service.User)
	a.Controllers.Product = *controllers.NewProductController(a.Service.Product)
	a.Controllers.Order = *controllers.NewOrderController(a.Service.Order)
}

func (a *App) SeedDB() {
	conn := a.Store.MysqlCon.Connection()
	_, err := conn.Exec(`CREATE TABLE IF NOT EXISTS users (
							id INT AUTO_INCREMENT PRIMARY KEY,
							name VARCHAR(50),
							email VARCHAR(50),
							type VARCHAR(50)
						);`)
	if err != nil {
		log.Fatal("unable to initiate db, error: ", err)
	}
	_, err = conn.Exec(`CREATE TABLE IF NOT EXISTS products (
							id INT AUTO_INCREMENT PRIMARY KEY,
							name VARCHAR(50),
							quantity INT,
							price INT,
							valid BOOLEAN,
							user_id INT,
							FOREIGN KEY (user_id) REFERENCES users (id)
						);`)
	if err != nil {
		log.Fatal("unable to initiate db, error: ", err)
	}
	_, err = conn.Exec(`CREATE TABLE IF NOT EXISTS orders (
							id INT AUTO_INCREMENT PRIMARY KEY,
							product_id INT,
							seller_id INT,
							buyer_id INT,
							status VARCHAR(50),
							FOREIGN KEY (product_id) REFERENCES products (id),
							FOREIGN KEY (seller_id) REFERENCES users (id),
							FOREIGN KEY (buyer_id) REFERENCES users (id)
						);`)

	if err != nil {
		log.Fatal("unable to initiate db, error: ", err)
	}

}

func (a *App) InstantiateApp() {
	a.InstantiateStores()
	a.SeedDB()
	a.InstantiateRepositories()
	a.InstantiateServices()
	a.InstantiateControllers()
}
