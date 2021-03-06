package db

import (
	pg "github.com/mxc-foundation/mxprotocol-server/m2m/db/postgres_db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/types"
)

type aggWalletUsageDBInterface interface {
	CreateAggWltUsgTable() error
	InsertAggWltUsg(awu types.AggWltUsg) (insertIndex int64, err error)
	GetWalletUsageHist(walletId int64, offset int64, limit int64) ([]types.AggWltUsg, error)
	GetWalletUsageHistCnt(walletId int64) (recCnt int64, err error)
	CreateAggWltUsgFunctions() error
	ExecAggWltUsgPayments(internalTx types.InternalTx, superNodeIncomeVal float64) (updatedBalance float64, err error)
}

var AggWalletUsage = aggWalletUsageDBInterface(&pg.PgAggWalletUsage)
