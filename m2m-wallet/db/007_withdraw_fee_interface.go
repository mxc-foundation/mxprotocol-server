package db

import (
	"time"
)

func dbCreateWithdrawFeeTable() error {
	return dbHandler.CreateWithdrawFeeTable()
}

func DbInsertWithdrawFee(extCurrencyAbbr string, withdrawFee float64) (insertIndex int64, err error) {
	id, err := DbGetExtCurrencyIdByAbbr(extCurrencyAbbr)
	if err != nil {
		return id, err
	}
	w := WithdrawFee{
		FkExtCurr:  id,
		Fee:        withdrawFee,
		InsertTime: time.Now().UTC(),
		Status:     "ACTIVE",
	}
	return dbHandler.InsertWithdrawFee(w)
}

func DbGetActiveWithdrawFee(extCurrAbv string) (withdrawFee float64, err error) {
	return dbHandler.GetActiveWithdrawFee(extCurrAbv)
}