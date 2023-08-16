package mysqldb

import (
	"database/sql"
	"fmt"
	"log"
	"marketplace/stores"

	_ "github.com/go-sql-driver/mysql"
)

// Store stores the structure of the store for the postgres structure
type Store struct {
	stores.Imysqldb
	Client *sql.DB
}

type MysqlConf struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

// New function is responsible for creating a new instance of postgres store
func New() stores.Imysqldb {

	cfg := MysqlConf{
		Username: "root",
		Password: "test1234",
		Host:     "db",
		Port:     "3306",
		Database: "",
	}

	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	// create db if not present
	db, err := sql.Open("mysql", sqlInfo)

	if err != nil {
		log.Fatal("unable to connect to db, error: ", err)
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS marketplace;")
	if err != nil {
		log.Fatal("unable to create db, error: ", err)
	}
	db.Close()

	// connect to db
	cfg.Database = "marketplace"
	sqlInfo = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err = sql.Open("mysql", sqlInfo)

	if err != nil {
		log.Fatal("unable to connect to db, error: ", err)
	}

	// ping database
	err = db.Ping()
	if err != nil {
		log.Fatal("unable to ping db, error: ", err)
	}

	return &Store{
		Client: db,
	}
}

// Connection returns the live connection to postgres DB
func (s *Store) Connection() *sql.DB {
	return s.Client
}
