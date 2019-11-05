package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type stakeRevenueDBInterface interface {
	CreateStakeRevenueTable() error
	InsertStakeRevenue(stakeId int64, stakeReveneuPeriodId int64, revenueAmount float64) (insertIndex int64, err error)
}

var StakeRevenue = stakeRevenueDBInterface(&pg.PgStakeRevenue)
