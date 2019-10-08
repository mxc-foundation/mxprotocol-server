package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type aggWalletUsageDBInterface interface {
	CreateAggWltUsgTable() error
	InsertAggWltUsg(awu types.AggWltUsg) (insertIndex int64, err error)
	GetWalletUsageHist(walletId int64, offset int64, limit int64) ([]types.AggWltUsg, error)
	GetWalletUsageHistCnt(walletId int64) (recCnt int64, err error)
}

var AggWalletUsage = aggWalletUsageDBInterface(&pg.PgAggWalletUsage)
