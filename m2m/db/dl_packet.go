package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	types "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

func DbCreateDlPktTable() error {
	return pg.PgDB.CreateDlPktTable()
}

func DbInsertDlPkt(dlp types.DlPkt) (insertIndex int64, err error) {
	return pg.PgDB.InsertDlPkt(dlp)
}
