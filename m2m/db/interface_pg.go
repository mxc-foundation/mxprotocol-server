package db

import (
	"database/sql"
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type PostgresHandler struct {
	*sql.DB
}

func (h *PostgresHandler) Begin() (*TxHandler, error) {
	tx, err := h.DB.Begin()
	return &TxHandler{tx}, err
}

func (h *PostgresHandler) AddDB(d *sql.DB) {
	h.DB = d
	pg.PgDB = d
}
