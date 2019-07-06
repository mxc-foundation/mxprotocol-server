package db

import (
	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
	"time"
)

func DbCreateExtAccountTable() error {
	return pgDb.CreateExtAccountTable()
}

func DBInsertExtAccount(walletId int64, newAccount string, currencyAbbr string) (insertIndex int, err error) {
	// get extCurrencyId from currencyAbbr
	extCurrencyId, err := DbGetExtCurrencyIdByAbbr(currencyAbbr)
	if err != nil {
		return 0, err
	}

	ea := pstgDb.ExtAccount{
		FkWallet:      walletId,
		FkExtCurrency: extCurrencyId,
		Account_adr:   newAccount,
		Insert_time:   time.Now().UTC(),
	}
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
