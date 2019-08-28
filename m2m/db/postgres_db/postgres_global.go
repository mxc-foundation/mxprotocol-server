package postgres_db

import (
	"database/sql"
)

var timeLayout string = "2006-01-02T15:04:05.000000Z"

type PGHandler struct{
	*sql.DB
}

var PgDB PGHandler

func (s PGHandler) GetDB(dbname string) *sql.DB{
	return PgDB.DB
}

func (s PGHandler) AddDB(db *sql.DB) {
	PgDB.DB = db
}