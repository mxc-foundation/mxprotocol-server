package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

func dbCreateInternalTxTable() error {
	return pg.CreateInternalTxTable()
}

func DbInsertInternalTx(it pg.InternalTx) (insertIndex int64, err error) {
	return pg.InsertInternalTx(it)
}
