package db

import pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"

func dbCreateInternalTxTable() error {
	return pgDb.CreateInternalTxTable()
}

func DbInsertInternalTx(it pstgDb.InternalTx) (insertIndex int64, err error) {
	return pgDb.InsertInternalTx(it)
}
