package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type stakeDBInterface interface {
	CreateStakeTable() error
	InsertStake(walletId int64, amount float64) (insertIndex int64, err error)
}

var Stake = stakeDBInterface(&pg.PgStake)
