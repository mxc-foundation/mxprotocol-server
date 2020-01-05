package db

import (
	pg "github.com/mxc-foundation/mxprotocol-server/m2m/db/postgres_db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/types"
)

type stakeRevenueDBInterface interface {
	CreateStakeRevenueTable() error
	CreateStakeRevenueFunctions() error
	InsertStakeRevenue(stakeId int64, stakeRevenuePeriodId int64, revenueAmount float64) (insertIndex int64, err error)
	GetStakeRevenueHistory(walletId int64, offset int64, limit int64) (stakeRevenueHists []types.StakeRevenueHist, err error)
	GetStakeRevenueHistoryCnt(walletId int64) (recCnt int64, err error)
}

var StakeRevenue = stakeRevenueDBInterface(&pg.PgStakeRevenue)
