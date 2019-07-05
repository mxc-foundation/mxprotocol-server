package db

import (
	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

func DbCreateTopupTable() error {
	return pgDb.CreateTopupTable()
}

func DbInsertTopup(tu pstgDb.Topup) (insertIndex int, err error) {
	return pgDb.InsertTopup(tu)
}
