package db

func dbCreateInternalTxTable() error {
	return db.CreateInternalTxTable()
}

func DbInsertInternalTx(it InternalTx) (insertIndex int64, err error) {
	return db.InsertInternalTx(it)
}
