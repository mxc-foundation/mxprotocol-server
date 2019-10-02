package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type aggWalletUsageDBInterface interface {
	CreateAggWltUsgTable() error
	InsertAggWltUsg(awu types.AggWltUsg) (insertIndex int64, err error)
	GetWalletUsageHist(rogId int64) ([]types.AggWltUsg, error)
}

var AggWalletUsage = aggWalletUsageDBInterface(&pg.PgAggWalletUsage)
