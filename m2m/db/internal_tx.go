package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type internalTxDBInterface interface {
	CreateInternalTxTable() error
	InsertInternalTx(it pg.InternalTx) (insertIndex int64, err error)
}

var internalTx internalTxDBInterface

func dbCreateInternalTxTable() error {
	internalTx = &pg.PgInternalTx
	return internalTx.CreateInternalTxTable()
}

func DbInsertInternalTx(it pg.InternalTx) (insertIndex int64, err error) {
	return internalTx.InsertInternalTx(it)
}
