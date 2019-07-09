package db

import pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"

func DbCreateWithdrawFeeTable() error {
	return pgDb.CreateWithdrawFeeTable()
}

func DbInsertWithdrawFee(w pstgDb.WithdrawFee) (insertIndex int64, err error) {
	return pgDb.InsertWithdrawFee(w)
}

func DbGetActiveWithdrawFee(extCurrAbv string) (withdrawFee float64, err error) {
	return pgDb.GetActiveWithdrawFee(extCurrAbv)
}

func DbGetActiveWithdrawFeeId(extCurrAbv string) (withdrawFeeId int64, err error) {
	return pgDb.GetActiveWithdrawFeeId(extCurrAbv)
}
