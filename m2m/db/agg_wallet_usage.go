package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	types "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

func DbCreateAggWltUsgTable() error {
	return pg.PgDB.CreateAggWltUsgTable()
}

func DbInsertAggWltUsg(awu types.AggWltUsg) (insertIndex int64, err error) {
	return pg.PgDB.InsertAggWltUsg(awu)
}
