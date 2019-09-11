package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type aggWalletUsageDBInterface interface {
	CreateAggWltUsgTable() error
	InsertAggWltUsg(awu types.AggWltUsg) (insertIndex int64, err error)
}
var aggWalletUsage aggWalletUsageDBInterface

func DbCreateAggWltUsgTable() error {
	aggWalletUsage = &pg.PgAggWalletUsage
	return aggWalletUsage.CreateAggWltUsgTable()
}

func DbInsertAggWltUsg(awu types.AggWltUsg) (insertIndex int64, err error) {
	return aggWalletUsage.InsertAggWltUsg(awu)
}
