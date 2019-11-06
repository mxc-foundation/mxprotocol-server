package types

import "time"

type StakeRevenuePeriod struct {
	Id                   int64                    `db:"id"`
	StakingPeriodStart   time.Time                `db:"staking_period_start"`
	StakingPeriodEnd     time.Time                `db:"staking_period_end"`
	SuperNodeIncome      float64                  `db:"supernode_income"`
	IncomeToStakePortion float64                  `db:"income_to_stake_portion"`
	ExecStartTime        time.Time                `db:"exec_start_time"`
	ExecEndTime          time.Time                `db:"exec_end_time"`
	Status               StakeRevenuePeriodStatus `db:"status"`
}

type StakeRevenuePeriodStatus string

const (
	STAKE_REVENUE_IN_PROCESS StakeRevenuePeriodStatus = "IN_PROCESS"
	STAKE_REVENUE_COMPLETED  StakeRevenuePeriodStatus = "COMPLETED"
)
