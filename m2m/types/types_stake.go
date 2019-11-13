package types

import "time"

type Stake struct {
	Id             int64       `db:"id"`
	FkWallet       int64       `db:"fk_wallet"`
	Amount         float64     `db:"amount"`
	Status         StakeStatus `db:"status"`
	StartStakeTime time.Time   `db:"start_stake_time"`
	UnstakeTime    time.Time   `db:"unstake_time"`
}

type StakeStatus string

const (
	STAKING_ACTIVE   StakeStatus = "ACTIVE"
	STAKING_UNSTAKED StakeStatus = "UNSTAKED"
)
