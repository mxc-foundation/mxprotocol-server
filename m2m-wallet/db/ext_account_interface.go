package db

import pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"

func DbCreateExtAccountTable() error {
	return pgDb.CreateExtAccountTable()
}

func DBInsertExtAccount(ea pstgDb.ExtAccount) (insertIndex int, err error) {
	return pgDb.InsertExtAccount(ea)
}

func DbGetSuperNodeExtAccountAdr(extCurrAbv string) (acntAdr string, err error) {
	return pgDb.GetSuperNodeExtAccountAdr(extCurrAbv)
}

func DbGetUserExtAccountAdr(walletId int, extCurrAbv string) (acntAdr string, err error) {
	return pgDb.GetUserExtAccountAdr(walletId, extCurrAbv)
}

func DbGetLatestCheckedBlock(extAcntId int) (int, error) {
	return pgDb.GetLatestCheckedBlock(extAcntId)
}

func DbUpdateLatestCheckedBlock(extAcntId int, updatedBlockNum int) error {
	return pgDb.UpdateLatestCheckedBlock(extAcntId, updatedBlockNum)
}
