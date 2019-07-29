package db

import (
	"time"

	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

func dbCreateExtAccountTable() error {
	return pgDb.CreateExtAccountTable()
}

func DBInsertExtAccount(walletId int64, newAccount string, currencyAbbr string) (insertIndex int64, err error) {
	extCurrencyId, err := DbGetExtCurrencyIdByAbbr(currencyAbbr)
	if err != nil {
		return extCurrencyId, err
	}

	ea := pstgDb.ExtAccount{
		FkWallet:      walletId,
		FkExtCurrency: extCurrencyId,
		Account_adr:   newAccount,
		Insert_time:   time.Now().UTC(),
	}

	extAcntId, errInsert := pgDb.InsertExtAccount(ea)

	return extAcntId, errInsert
}

func DbGetSuperNodeExtAccountAdr(extCurrAbv string) (acntAdr string, err error) {
	return pgDb.GetSuperNodeExtAccountAdr(extCurrAbv)
}

func DbGetSuperNodeExtAccountId(extCurrAbv string) (acntId int64, err error) {
	return pgDb.GetSuperNodeExtAccountId(extCurrAbv)
}

func DbGetUserExtAccountAdr(walletId int64, extCurrAbv string) (acntAdr string, err error) {
	return pgDb.GetUserExtAccountAdr(walletId, extCurrAbv)
}

func DbGetUserExtAccountId(walletId int64, extCurrAbv string) (int64, error) {
	return pgDb.GetUserExtAccountId(walletId, extCurrAbv)
}

func DbGetLatestCheckedBlock(extAcntId int64) (int64, error) {
	return pgDb.GetLatestCheckedBlock(extAcntId)
}

func DbUpdateLatestCheckedBlock(extAcntId int64, updatedBlockNum int64) error {
	return pgDb.UpdateLatestCheckedBlock(extAcntId, updatedBlockNum)
}

func DbGetExtAccountIdByAdr(acntAdr string) (int64, error) {
	return pgDb.GetExtAccountIdByAdr(acntAdr)
}

func DbGetExtAcntHist(walletId int64, offset int64, limit int64) ([]pstgDb.ExtAccountHistRet, error) {
	return pgDb.GetExtAcntHist(walletId, offset, limit)
}
