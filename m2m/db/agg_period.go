package db

import (
	"time"

	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type aggPeriodDBInterface interface {
	CreateAggPeriodTable() error
	InsertAggPeriod(latestIdAccountedDlPkt int64, durationMinutes int64) (insertInd int64, err error)
	UpdateExecutedAggPeriod(aggPeriodId int64, execEndAt time.Time) (err error)
	GetLatestAccountedDlPktId() (latestAccountedDlPktId int64, err error)
}

var AggPeriod = aggPeriodDBInterface(&pg.PgAggPeriod)
