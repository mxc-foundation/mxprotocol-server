package db

func dbCreateInternalTxTable() error {
	return dbHandler.CreateInternalTxTable()
}

func DbInsertInternalTx(it InternalTx) (insertIndex int64, err error) {
	return dbHandler.InsertInternalTx(it)
}
