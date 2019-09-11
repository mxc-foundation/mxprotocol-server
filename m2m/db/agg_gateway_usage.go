package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	types "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

func DbCreateAggGwUsgTable() error {
	return pg.PgDB.CreateAggGwUsgTable()
}

func DbInsertAggGwUsg(agu types.AggGwUsg) (insertIndex int64, err error) {
	return pg.PgDB.InsertAggGwUsg(agu)
}
