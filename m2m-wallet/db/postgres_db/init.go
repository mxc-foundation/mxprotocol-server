package postgres_db

import "database/sql"

var timeLayout string = "2006-01-02T15:04:05.000000Z"

type DbSpec struct {
	Db         *sql.DB
	DriverName string
	Dburl      string
}
