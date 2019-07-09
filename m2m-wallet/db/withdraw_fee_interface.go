package db

import (
	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
	"time"
)

func DbCreateWithdrawFeeTable() error {
	return pgDb.CreateWithdrawFeeTable()
}

func DbInsertWithdrawFee(extCurrencyAbbr string, withdrawFee float64) (insertIndex int, err error) {
	id, err := DbGetExtCurrencyIdByAbbr(extCurrencyAbbr)
	if err != nil {
		return 0, err
	}
	w := pstgDb.WithdrawFee{
		FkExtCurr:  id,
		Fee:        withdrawFee,
		InsertTime: time.Now().UTC(),
		Status:     "ACTIVE",
	}
	return pgDb.InsertWithdrawFee(w)
}

func DbGetActiveWithdrawFee(extCurrAbv string) (withdrawFee float64, err error) {
	return pgDb.GetActiveWithdrawFee(extCurrAbv)
}
