package db

import (
	"time"
)


type ExtAccount struct {
	Id                 int64     `db:"id"`
	FkWallet           int64     `db:"fk_wallet"`
	FkExtCurrency      int64     `db:"fk_ext_currency"`
	Account_adr        string    `db:"account_adr"`
	Insert_time        time.Time `db:"insert_time"`
	Status             string    `db:"status"`
	LatestCheckedBlock int64     `db:"latest_checked_block"`
}

type ExtAccountHistRet struct {
	AccountAdr     string
	InsertTime     time.Time
	Status         string
	ExtCurrencyAbv string
}

func dbCreateExtAccountTable() error {
	return db.CreateExtAccountTable()
}

func DBInsertExtAccount(walletId int64, newAccount string, currencyAbbr string) (insertIndex int64, err error) {
	extCurrencyId, err := DbGetExtCurrencyIdByAbbr(currencyAbbr)
	if err != nil {
		return extCurrencyId, err
	}

	ea := ExtAccount{
		FkWallet:      walletId,
		FkExtCurrency: extCurrencyId,
		Account_adr:   newAccount,
		Insert_time:   time.Now().UTC(),
	}

	extAcntId, errInsert := db.InsertExtAccount(ea)

	return extAcntId, errInsert
}

func DbGetSuperNodeExtAccountAdr(extCurrAbv string) (acntAdr string, err error) {
	return db.GetSuperNodeExtAccountAdr(extCurrAbv)
}

func DbGetSuperNodeExtAccountId(extCurrAbv string) (acntId int64, err error) {
	return db.GetSuperNodeExtAccountId(extCurrAbv)
}

func DbGetUserExtAccountAdr(walletId int64, extCurrAbv string) (acntAdr string, err error) {
	return db.GetUserExtAccountAdr(walletId, extCurrAbv)
}

func DbGetUserExtAccountId(walletId int64, extCurrAbv string) (int64, error) {
	return db.GetUserExtAccountId(walletId, extCurrAbv)
}

func DbGetLatestCheckedBlock(extAcntId int64) (int64, error) {
	return db.GetLatestCheckedBlock(extAcntId)
}

func DbUpdateLatestCheckedBlock(extAcntId int64, updatedBlockNum int64) error {
	return db.UpdateLatestCheckedBlock(extAcntId, updatedBlockNum)
}

func DbGetExtAccountIdByAdr(acntAdr string) (int64, error) {
	return db.GetExtAccountIdByAdr(acntAdr)
}

func DbGetExtAcntHist(walletId int64, offset int64, limit int64) ([]ExtAccountHistRet, error) {
	return db.GetExtAcntHist(walletId, offset, limit)
}
