package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type extAccountDBInterface interface {
	CreateExtAccountTable() error
	InsertExtAccount(walletId int64, newAccount string, currencyAbbr string) (insertIndex int64, err error)
	GetSuperNodeExtAccountAdr(extCurrAbv string) (string, error)
	GetSuperNodeExtAccountId(extCurrAbv string) (int64, error)
	GetUserExtAccountAdr(walletId int64, extCurrAbv string) (string, error)
	GetUserExtAccountId(walletId int64, extCurrAbv string) (int64, error)
	GetExtAccountIdByAdr(acntAdr string, extCurrAbv string) (int64, error)
	GetLatestCheckedBlock(extAcntId int64) (int64, error)
	UpdateLatestCheckedBlock(extAcntId int64, updatedBlockNum int64) error
	GetExtAcntHist(walletId int64, offset int64, limit int64) ([]types.ExtAccountHistRet, error)
	GetExtAcntHistRecCnt(walletId int64) (recCnt int64, err error)
}

var ExtAccount = extAccountDBInterface(&pg.PgExtAccount)
