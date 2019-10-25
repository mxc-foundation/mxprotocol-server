package db

import (
	"time"

	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type aggPeriodDBInterface interface {
	CreateAggPeriodTable() error
	InsertAggPeriod(durationMinutes int64) (insertInd int64, latestIdAccountedDlPkt int64, err error)
	UpdateExecutedAggPeriod(aggPeriodId int64, execEndAt time.Time) (err error)
}

var AggPeriod = aggPeriodDBInterface(&pg.PgAggPeriod)
