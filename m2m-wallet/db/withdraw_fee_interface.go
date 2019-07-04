package db

import pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"

func DbCreateWithdrawFeeTable() error {
	return pgDb.CreateWithdrawFeeTable()
}

func DbInsertWithdrawFee(w *pstgDb.WithdrawFee) error {
	return pgDb.InsertWithdrawFee(w)
}
