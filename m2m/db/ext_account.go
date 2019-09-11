package db

import (
	"strings"
	"time"

	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type extAccountDBInterface interface {
	CreateExtAccountTable() error
	InsertExtAccount(ea pg.ExtAccount) (insertIndex int64, err error)
	GetSuperNodeExtAccountAdr(extCurrAbv string) (string, error)
	GetSuperNodeExtAccountId(extCurrAbv string) (int64, error)
	GetUserExtAccountAdr(walletId int64, extCurrAbv string) (string, error)
	GetUserExtAccountId(walletId int64, extCurrAbv string) (int64, error)
	GetExtAccountIdByAdr(acntAdr string, extCurrAbv string) (int64, error)
	GetLatestCheckedBlock(extAcntId int64) (int64, error)
	UpdateLatestCheckedBlock(extAcntId int64, updatedBlockNum int64) error
	GetExtAcntHist(walletId int64, offset int64, limit int64) ([]pg.ExtAccountHistRet, error)
	GetExtAcntHistRecCnt(walletId int64) (recCnt int64, err error)
}

var extAccount extAccountDBInterface

type ExtAccountHistRet pg.ExtAccountHistRet

func dbCreateExtAccountTable() error {
	extAccount = &pg.PgExtAccount
	return extAccount.CreateExtAccountTable()
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

	extAcntId, errInsert := extAccount.InsertExtAccount(ea)

	return extAcntId, errInsert
}

func DbGetSuperNodeExtAccountAdr(extCurrAbv string) (acntAdr string, err error) {
	return extAccount.GetSuperNodeExtAccountAdr(extCurrAbv)
}

func DbGetSuperNodeExtAccountId(extCurrAbv string) (acntId int64, err error) {
	return extAccount.GetSuperNodeExtAccountId(extCurrAbv)
}

func DbGetUserExtAccountAdr(walletId int64, extCurrAbv string) (acntAdr string, err error) {
	return extAccount.GetUserExtAccountAdr(walletId, extCurrAbv)
}

func DbGetUserExtAccountId(walletId int64, extCurrAbv string) (int64, error) {
	return extAccount.GetUserExtAccountId(walletId, extCurrAbv)
}

func DbGetLatestCheckedBlock(extAcntId int64) (int64, error) {
	return extAccount.GetLatestCheckedBlock(extAcntId)
}

func DbUpdateLatestCheckedBlock(extAcntId int64, updatedBlockNum int64) error {
	return extAccount.UpdateLatestCheckedBlock(extAcntId, updatedBlockNum)
}

func DbGetExtAccountIdByAdr(acntAdr string, extCurrAbv string) (int64, error) {
	return extAccount.GetExtAccountIdByAdr(strings.ToLower(acntAdr), extCurrAbv)
}

func castExtAccountHistRet(acntHist []pg.ExtAccountHistRet, err1 error) (castedVal []ExtAccountHistRet, err error) {
	for _, v := range acntHist {
		castedVal = append(castedVal, ExtAccountHistRet(v))
	}
	return castedVal, err1
}

func DbGetExtAcntHist(walletId int64, offset int64, limit int64) ([]ExtAccountHistRet, error) {
	return castExtAccountHistRet(extAccount.GetExtAcntHist(walletId, offset, limit))
}

func DbGetExtAcntHistRecCnt(walletId int64) (int64, error) {
	return extAccount.GetExtAcntHistRecCnt(walletId)
}
