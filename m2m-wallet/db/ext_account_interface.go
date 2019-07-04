package db

import pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"

func DbCreateExtAccountTable() error {
	return pgDb.CreateExtAccountTable()
}

func DBInsertExtAccount(ea *pstgDb.ExtAccount) error {
	return pgDb.InsertExtAccount(ea)
}
