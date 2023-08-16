package stores

import "database/sql"

type Imysqldb interface {
	Connection() *sql.DB
}
