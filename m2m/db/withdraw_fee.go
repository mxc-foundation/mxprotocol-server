package db

import (
	pg "github.com/mxc-foundation/mxprotocol-server/m2m/db/postgres_db"
)

type withdrawFeeDBInterface interface {
	CreateWithdrawFeeTable() error
	InsertWithdrawFee(extCurrencyAbbr string, wdFee float64) (insertIndex int64, err error)
	GetActiveWithdrawFee(extCurrAbv string) (withdrawFee float64, err error)
	GetActiveWithdrawFeeId(extCurrAbv string) (withdrawFee int64, err error)
}

var WithdrawFee = withdrawFeeDBInterface(&pg.PgWithdrawFee)
