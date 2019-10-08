package db

import (
	"time"

	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type aggPeriodDBInterface interface {
	CreateAggPeriodTable() error
	InsertAggPeriod(startAt time.Time, durationMinutes int64) (insertInd int64, err error)
}

var AggPeriod = aggPeriodDBInterface(&pg.PgAggPeriod)
