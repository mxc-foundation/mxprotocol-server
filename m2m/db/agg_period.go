package db

import (
	pg "github.com/mxc-foundation/mxprotocol-server/m2m/db/postgres_db"
)

type aggPeriodDBInterface interface {
	CreateAggPeriodTable() error
	InsertAggPeriod(durationMinutes int64) (insertInd int64, latestIdAccountedDlPkt int64, err error)
	UpdateSuccessfulExecutedAggPeriod(aggPeriodId int64, latestIdAccountedDlPkt int64) (err error)
}

var AggPeriod = aggPeriodDBInterface(&pg.PgAggPeriod)
