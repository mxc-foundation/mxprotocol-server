package types

import "time"

type StakeRevenue struct {
	Id                   int64   `db:"id"`
	FkStakeRevenuePeriod int64   `db:"fk_stake_revenue_period"`
	FkStake              int64   `db:"fk_stake"`
	FkWallet             int64   `db:"r_fk_wallet"`
	RevenueAmount        float64 `db:"revenue_amount"`
	UpdatedBalance       float64 `db:"updated_balance"`
}

type StakeRevenueHist struct {
	WalletId             int64
	StakeAmount          float64
	StartStakeTime       time.Time
	UnstakeTime          time.Time
	StakingPeriodStart   time.Time
	StakingPeriodEnd     time.Time
	SuperNodeIncome      float64
	IncomeToStakePortion float64
	RevenueAmount        float64
	UpdatedBalance       float64
}
