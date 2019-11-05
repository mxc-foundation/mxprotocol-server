package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type stakeRevenuePeriodDBInterface interface {
	CreateStakeRevenuePeriodTable() error
}

var StakeRevenuePeriod = stakeRevenuePeriodDBInterface(&pg.PgStakeRevenuePeriod)
