package product

import (
	"database/sql"
	"marketplace/models"
	"marketplace/repositories"
	"marketplace/stores"
)

type Repository struct {
	repositories.IProductRepository
	client *sql.DB
}

// New initiates the signal repository structure
func New(mysqlStore stores.Imysqldb) repositories.IProductRepository {
	mysql := mysqlStore.Connection()
	return &Repository{
		client: mysql,
	}
}

func (d *Repository) AddProduct(name string, quantity int, price int, userID int) (int64, error) {
	txn, err := d.client.Begin()
	if err != nil {
		return -1, err
	}

	defer txn.Rollback()
	result, err := txn.Exec("INSERT INTO products (user_id, name, price, quantity) VALUES (?, ?, ?, ?)", userID, name, price, quantity)

	if err != nil {
		return -1, err
	}
	insertedID, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	txn.Commit()

	return insertedID, nil
}

// func (d *Repository) AddProducts(products []models.Product, userID int) error {
// 	txn, err := d.client.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	defer txn.Rollback()

// 	prepQuery := "INSERT INTO products (user_id, name, price, quantity, valid) VALUES (?, ?, ?, ?)"
// 	stmt, err := txn.Prepare(prepQuery)

// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	for _, p := range products {
// 		_, err = stmt.Exec(userID, p.Name, p.Price, p.Quantity, true)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	txn.Commit()
// 	return nil
// }

func (d *Repository) GetProduct(productID int) (models.Product, error) {
	result := d.client.QueryRow("SELECT id, name, quantity, price, user_id FROM products WHERE id = ?", productID)

	var product models.Product
	err := result.Scan(&product.ID, &product.Name, &product.Quantity, &product.Price, &product.UserID)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (d *Repository) GetProductByNameAndUser(productName string, userID int) (models.Product, error) {
	result := d.client.QueryRow("SELECT id, name, quantity, price, user_id FROM products where name = ? && user_id = ?", productName, userID)

	var product models.Product
	err := result.Scan(&product.ID, &product.Name, &product.Quantity, &product.Price)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (d *Repository) UpdateProductQuantity(quantity int, productID int) error {
	txn, err := d.client.Begin()
	if err != nil {
		return err
	}

	defer txn.Rollback()

	txn.Exec("UPDATE products SET quantity =? where product_id = ?", quantity, productID)
	txn.Commit()

	return nil
}

func (d *Repository) InvalidateProduct(productID int) error {
	txn, err := d.client.Begin()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	txn.Exec("UPDATE products SET valid = false where product_id = ?", productID)
	txn.Commit()

	return nil
}

func (d *Repository) BuyerProducts(buyerID int) ([]models.Product, error) {
	result, err := d.client.Query("SELECT p.id, p.name, p.quantity, p.price FROM products p JOIN users u ON p.user_id = u.id WHERE u.type = 'buyer' && p.valid = true && u.id = ?", buyerID)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var products []models.Product

	for result.Next() {
		var product models.Product
		err := result.Scan(&product.ID, &product.Name, &product.Quantity, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (d *Repository) SellerProducts(sellerID int) ([]models.Product, error) {
	result, err := d.client.Query("SELECT p.id, p.name, p.quantity, p.price FROM products p JOIN users u ON p.user_id = u.id WHERE u.type = 'seller' && p.valid = true && u.id = ?", sellerID)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var products []models.Product

	for result.Next() {
		var product models.Product
		err := result.Scan(&product.ID, &product.Name, &product.Quantity, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
