package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type aggDeviceUsageDBInterface interface {
	CreateAggDvUsgTable() error
	InsertAggDvUsg(adu types.AggDvUsg) (insertIndex int64, err error)
}

var AggDeviceUsage = aggDeviceUsageDBInterface(&pg.PgAggDeviceUsage)
