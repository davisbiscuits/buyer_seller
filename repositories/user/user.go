package user

import (
	"database/sql"
	"marketplace/models"
	"marketplace/repositories"
	"marketplace/stores"
)

type Repository struct {
	repositories.IUserRepository
	client *sql.DB
}

// New initiates the user repository structure
func New(mysqlStore stores.Imysqldb) repositories.IUserRepository {
	mysql := mysqlStore.Connection()
	return &Repository{
		client: mysql,
	}
}

func (d *Repository) CreateUser(userType string, userEmail string, userName string) (int64, error) {
	txn, err := d.client.Begin()
	if err != nil {
		return -1, err
	}

	defer txn.Rollback()

	result, err := txn.Exec("INSERT INTO users (type, email, name) VALUES (?, ?, ?)", userType, userEmail, userName)
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

func (d *Repository) FindSellers(productName string, productQuantity int, productPrice int) ([]models.User, error) {

	result, err := d.client.Query(`SELECT u.id, u.name, u.email, u.type FROM users u join products p ON u.id = p.user_id
	where u.type = 'seller' AND p.name = ? AND p.price <= ? AND p.quantity >= ? ORDER BY p.price ASC`, productName, productPrice, productQuantity)
	if err != nil {
		return nil, err
	}

	var users []models.User

	for result.Next() {
		var user models.User
		err := result.Scan(&user.ID, &user.Name, &user.Email, &user.Type)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
