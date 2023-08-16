package order

import (
	"database/sql"
	"marketplace/dtos"
	"marketplace/repositories"
	"marketplace/stores"
)

type Repository struct {
	repositories.IOrderRepository
	client *sql.DB
}

// New initiates the signal repository structure
func New(mysqlStore stores.Imysqldb) repositories.IOrderRepository {
	mysql := mysqlStore.Connection()
	return &Repository{
		client: mysql,
	}
}

func (d *Repository) AddOrder(productID int, sellerID int, buyerID int) (int64, error) {
	txn, err := d.client.Begin()
	if err != nil {
		return -1, err
	}

	defer txn.Rollback()

	result, err := txn.Exec("INSERT INTO orders (seller_id, buyer_id, product_id, status, valid) VALUES (?, ?, ?, ?, ?)", sellerID, buyerID, productID, "proposed", true)
	if err != nil {
		return -1, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	err = txn.Commit()
	return insertedID, err
}

func (d *Repository) AcceptOrder(orderID int) error {
	txn, err := d.client.Begin()
	if err != nil {
		return err
	}

	defer txn.Rollback()

	txn.Exec("UPDATE orders SET status = 'accepted' where id = ?", orderID)

	err = txn.Commit()
	return err
}

func (d *Repository) RejectOrder(orderID int) error {
	txn, err := d.client.Begin()
	if err != nil {
		return err
	}

	defer txn.Rollback()

	txn.Exec("UPDATE orders SET status = ? where id = ?", "reject", orderID)

	err = txn.Commit()
	return err
}

func (d *Repository) GetOrder(orderID int) (dtos.OrderResponse, error) {

	result := d.client.QueryRow(`SELECT o.id, o.status, s.id as seller_id, s.name as seller_name, s.email as seller_email,
										b.id as buyer_id, b.name as buyer_name, b.email as buyer_email,
								 	    p.id as product_id, p.name as product_name, p.quantity as product_quantity, p.price as product_price FROM orders o 
								 JOIN users s ON o.seller_id = s.id AND s.type = 'seller'
								 JOIN users b ON o.buyer_id = s.id AND s.type = 'buyer'
								 JOIN products p ON o.product_id = p.id AND p.valid = true where o.orderID = ?`, orderID)

	var order dtos.OrderResponse
	err := result.Scan(&order.ID, &order.Status,
		&order.Seller.ID, &order.Seller.Name, &order.Seller.Email,
		&order.Buyer.ID, &order.Buyer.Name, &order.Buyer.Email,
		&order.Product.ID, &order.Product.Name, &order.Product.Quantity, &order.Product.Price,
	)

	return order, err
}

func (d *Repository) BuyerOrders(buyerID int) ([]dtos.BuyerOrdersResponse, error) {

	result, err := d.client.Query(`SELECT o.id, o.status, u.id as seller_id, u.name as seller_name, u.email as seller_email, 
										  p.id as product_id, p.name as product_name, p.quantity as product_quantity, p.price as product_price FROM orders o 
								   JOIN users u ON o.seller_id = u.id AND u.type = 'seller'
								   JOIN products p ON o.product_id = p.id AND p.valid = true where o.buyer_id = ?`, buyerID)
	if err != nil {
		return nil, err
	}

	var buyerOrders []dtos.BuyerOrdersResponse
	for result.Next() {
		var buyerOrder dtos.BuyerOrdersResponse
		err := result.Scan(&buyerOrder.ID, &buyerOrder.Status,
			&buyerOrder.Seller.ID, &buyerOrder.Seller.Name, &buyerOrder.Seller.Email,
			&buyerOrder.Requirement.ID, &buyerOrder.Requirement.Name, &buyerOrder.Requirement.Quantity, &buyerOrder.Requirement.Price)

		if err != nil {
			return nil, err
		}
		buyerOrders = append(buyerOrders, buyerOrder)
	}
	if err = result.Err(); err != nil {
		return nil, err
	}
	return buyerOrders, nil
}

func (d *Repository) SellerOrders(sellerID int) ([]dtos.SellerOrdersResponse, error) {

	result, err := d.client.Query(`SELECT o.id, o.status, u.id as buyer_id, u.name as buyer_name, u.email as buyer_email, 
										  p.id as product_id, p.name as product_name, p.quantity as product_quantity, p.price as product_price FROM orders o 
								   JOIN users u ON o.buyer_id = u.id AND u.type = 'buyer'
								   JOIN products p ON o.product_id = p.id AND p.valid = true where o.seller_id = ?`, sellerID)

	var sellerOrders []dtos.SellerOrdersResponse
	for result.Next() {
		var sellerOrder dtos.SellerOrdersResponse
		err := result.Scan(&sellerOrder.ID, &sellerOrder.Status,
			&sellerOrder.Buyer.ID, &sellerOrder.Buyer.Name, &sellerOrder.Buyer.Email,
			&sellerOrder.Product.ID, &sellerOrder.Product.Name, &sellerOrder.Product.Quantity, &sellerOrder.Product.Price)

		if err != nil {
			return nil, err
		}
		sellerOrders = append(sellerOrders, sellerOrder)
	}

	return sellerOrders, err
}
