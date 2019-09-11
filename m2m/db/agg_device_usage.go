package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	types "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

func DbCreateAggDvUsgTable() error {
	return pg.PgDB.CreateAggDvUsgTable()
}

func DbInsertAggDvUsg(adu types.AggDvUsg) (insertIndex int64, err error) {
	return pg.PgDB.InsertAggDvUsg(adu)
}
