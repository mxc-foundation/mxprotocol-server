package postgres_db

import "database/sql"

type DbSpec struct {
	Db         *sql.DB
	DriverName string
	Dburl      string
}
