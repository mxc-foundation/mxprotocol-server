package db

import (
	pg "github.com/mxc-foundation/mxprotocol-server/m2m/db/postgres_db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/types"
)

type aggDeviceUsageDBInterface interface {
	CreateAggDvUsgTable() error
	InsertAggDvUsg(adu types.AggDvUsg) (insertIndex int64, err error)
}

var AggDeviceUsage = aggDeviceUsageDBInterface(&pg.PgAggDeviceUsage)
