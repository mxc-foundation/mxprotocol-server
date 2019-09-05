package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	"time"
)

func dbCreateWithdrawFeeTable() error {
	return pg.CreateWithdrawFeeTable()
}

func DbInsertWithdrawFee(extCurrencyAbbr string, withdrawFee float64) (insertIndex int64, err error) {
	id, err := DbGetExtCurrencyIdByAbbr(extCurrencyAbbr)
	if err != nil {
		return id, err
	}
	w := pg.WithdrawFee{
		FkExtCurr:  id,
		Fee:        withdrawFee,
		InsertTime: time.Now().UTC(),
		Status:     "ACTIVE",
	}
	return pg.InsertWithdrawFee(w)
}

func DbGetActiveWithdrawFee(extCurrAbv string) (withdrawFee float64, err error) {
	return pg.GetActiveWithdrawFee(extCurrAbv)
}
