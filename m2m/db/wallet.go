package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type walletDBInterface interface {
	CreateWalletTable() error
	InsertWallet(orgId int64, walletType types.WalletType) (insertIndex int64, err error)
	GetWalletIdFromOrgId(orgIdLora int64) (int64, error)
	GetWalletBalance(walletId int64) (float64, error)
	GetWalletIdofActiveAcnt(acntAdr string, externalCur string) (walletId int64, err error)
	GetWalletIdSuperNode() (walletId int64, err error)
	UpdateBalanceByWalletId(walletId int64, newBalance float64) error
}

var Wallet = walletDBInterface(&pg.PgWallet)
