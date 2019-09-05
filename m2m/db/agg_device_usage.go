package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type aggDeviceUsageDBInterface interface {
	CreateAggDvUsgTable() error
	InsertAggDvUsg(adu types.AggDvUsg) (insertIndex int64, err error)
}

var aggDeviceUsage aggDeviceUsageDBInterface

func DbCreateAggDvUsgTable() error {
	aggDeviceUsage = &pg.PgAggDeviceUsage
	return aggDeviceUsage.CreateAggDvUsgTable()
}

func DbInsertAggDvUsg(adu types.AggDvUsg) (insertIndex int64, err error) {
	return aggDeviceUsage.InsertAggDvUsg(adu)
}
