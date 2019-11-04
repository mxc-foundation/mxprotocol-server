package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type stakeDBInterface interface {
	CreateStakeTable() error
}

var Stake = stakeDBInterface(&pg.PgStake)
