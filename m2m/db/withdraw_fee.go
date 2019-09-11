package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	"time"
)

type withdrawFeeDBInterface interface {
	CreateWithdrawFeeTable() error
	InsertWithdrawFee(wf pg.WithdrawFee) (insertIndex int64, err error)
	GetActiveWithdrawFee(extCurrAbv string) (withdrawFee float64, err error)
	GetActiveWithdrawFeeId(extCurrAbv string) (withdrawFee int64, err error)
}

var withdrawFee withdrawFeeDBInterface

func dbCreateWithdrawFeeTable() error {
	withdrawFee = &pg.PgWithdrawFee
	return withdrawFee.CreateWithdrawFeeTable()
}

func DbInsertWithdrawFee(extCurrencyAbbr string, wdFee float64) (insertIndex int64, err error) {
	id, err := DbGetExtCurrencyIdByAbbr(extCurrencyAbbr)
	if err != nil {
		return id, err
	}
	w := pg.WithdrawFee{
		FkExtCurr:  id,
		Fee:        wdFee,
		InsertTime: time.Now().UTC(),
		Status:     "ACTIVE",
	}
	return withdrawFee.InsertWithdrawFee(w)
}

func DbGetActiveWithdrawFee(extCurrAbv string) (wdFee float64, err error) {
	return withdrawFee.GetActiveWithdrawFee(extCurrAbv)
}
