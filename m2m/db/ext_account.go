package db

import (
	"strings"
	"time"

	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type ExtAccountHistRet pg.ExtAccountHistRet

func dbCreateExtAccountTable() error {
	return pg.CreateExtAccountTable()
}

func DBInsertExtAccount(walletId int64, newAccount string, currencyAbbr string) (insertIndex int64, err error) {
	extCurrencyId, err := DbGetExtCurrencyIdByAbbr(currencyAbbr)
	if err != nil {
		return extCurrencyId, err
	}

	ea := pg.ExtAccount{
		FkWallet:      walletId,
		FkExtCurrency: extCurrencyId,
		AccountAdr:    strings.ToLower(newAccount),
		InsertTime:    time.Now().UTC(),
	}

	extAcntId, errInsert := pg.InsertExtAccount(ea)

	return extAcntId, errInsert
}

func DbGetSuperNodeExtAccountAdr(extCurrAbv string) (acntAdr string, err error) {
	return pg.GetSuperNodeExtAccountAdr(extCurrAbv)
}

func DbGetSuperNodeExtAccountId(extCurrAbv string) (acntId int64, err error) {
	return pg.GetSuperNodeExtAccountId(extCurrAbv)
}

func DbGetUserExtAccountAdr(walletId int64, extCurrAbv string) (acntAdr string, err error) {
	return pg.GetUserExtAccountAdr(walletId, extCurrAbv)
}

func DbGetUserExtAccountId(walletId int64, extCurrAbv string) (int64, error) {
	return pg.GetUserExtAccountId(walletId, extCurrAbv)
}

func DbGetLatestCheckedBlock(extAcntId int64) (int64, error) {
	return pg.GetLatestCheckedBlock(extAcntId)
}

func DbUpdateLatestCheckedBlock(extAcntId int64, updatedBlockNum int64) error {
	return pg.UpdateLatestCheckedBlock(extAcntId, updatedBlockNum)
}

func DbGetExtAccountIdByAdr(acntAdr string, extCurrAbv string) (int64, error) {
	return pg.GetExtAccountIdByAdr(strings.ToLower(acntAdr), extCurrAbv)
}

func castExtAccountHistRet(acntHist []pg.ExtAccountHistRet, err1 error) (castedVal []ExtAccountHistRet, err error) {
	for _, v := range acntHist {
		castedVal = append(castedVal, ExtAccountHistRet(v))
	}
	return castedVal, err1
}

func DbGetExtAcntHist(walletId int64, offset int64, limit int64) ([]ExtAccountHistRet, error) {
	return castExtAccountHistRet(pg.GetExtAcntHist(walletId, offset, limit))
}

func DbGetExtAcntHistRecCnt(walletId int64) (int64, error) {
	return pg.GetExtAcntHistRecCnt(walletId)
}
