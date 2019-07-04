package db

func DbCreateWithdrawTable() error {
	return pgDb.CreateWithdrawTable()
}
