package db

import (
	"time"

	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type stakeRevenuePeriodDBInterface interface {
	CreateStakeRevenuePeriodTable() error
	InsertStakeRevenuePeriod(StakingPeriodStart time.Time, StakingPeriodEnd time.Time, SuperNodeIncome float64, IncomeToStakePortion float64) (insertIndex int64, err error)
	UpdateCompletedStakeReveneuPeriod(stakeReveneuPeriodId int64) error
}

var StakeRevenuePeriod = stakeRevenuePeriodDBInterface(&pg.PgStakeRevenuePeriod)
