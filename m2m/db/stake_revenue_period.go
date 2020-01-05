package db

import (
	"time"

	pg "github.com/mxc-foundation/mxprotocol-server/m2m/db/postgres_db"
)

type stakeRevenuePeriodDBInterface interface {
	CreateStakeRevenuePeriodTable() error
	InsertStakeRevenuePeriod(StakingPeriodStart time.Time, StakingPeriodEnd time.Time, SuperNodeIncome float64, IncomeToStakePortion float64) (insertIndex int64, err error)
	UpdateCompletedStakeRevenuePeriod(stakeRevenuePeriodId int64) error
}

var StakeRevenuePeriod = stakeRevenuePeriodDBInterface(&pg.PgStakeRevenuePeriod)
