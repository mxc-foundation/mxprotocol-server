package db

import pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"

func DbCreateInternalTxTable() error {
	return pgDb.CreateInternalTxTable()
}

func DbInsertInternalTx(it pstgDb.InternalTx) (insertIndex int, err error) {
	return pgDb.InsertInternalTx(it)
}
