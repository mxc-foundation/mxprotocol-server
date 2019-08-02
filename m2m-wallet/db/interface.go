package db

import "database/sql"

var db *DBHandler
type DBHandler struct {
	*sql.DB
}

func DB() *DBHandler {
	return db
}

type TxHandler struct {
	*sql.Tx
}

func (db *DBHandler) Begin() (*TxHandler, error) {
	tx, err := db.DB.Begin()
	return &TxHandler{tx}, err
}