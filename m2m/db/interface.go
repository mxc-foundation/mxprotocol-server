package db

import (
	"database/sql"
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

var i DBInterface

type DBInterface interface {
	GetDB(string) *sql.DB
	AddDB(*sql.DB)
}

func addDB(inter DBInterface, db *sql.DB) DBHandler {
	if _, ok := inter.(*pg.PGHandler); ok {
		i.AddDB(db)
		return DBHandler{db}
	}

	return DBHandler{nil}
}

type TxHandler struct {
	*sql.Tx
}

type DBHandler struct {
	*sql.DB
}

func (db *DBHandler) Begin() (*TxHandler, error) {
	tx, err := db.DB.Begin()
	return &TxHandler{tx}, err
}
