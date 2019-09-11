package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type WalletType string

const (
	USER        WalletType = "USER"
	SUPER_ADMIN WalletType = "SUPER_ADMIN"
)

type walletDBInterface interface {
	CreateWalletTable() error
	InsertWallet(w pg.Wallet) (insertIndex int64, err error)
	GetWalletIdFromOrgId(orgIdLora int64) (int64, error)
	GetWalletBalance(walletId int64) (float64, error)
	GetWalletIdofActiveAcnt(acntAdr string, externalCur string) (walletId int64, err error)
	GetWalletIdSuperNode() (walletId int64, err error)
	UpdateBalanceByWalletId(walletId int64, newBalance float64) error
}

var wallet walletDBInterface

func dbCreateWalletTable() error {
	wallet = &pg.PgWallet
	return wallet.CreateWalletTable()
}

func DbInsertWallet(orgId int64, walletType WalletType) (insertIndex int64, err error) {
	w := pg.Wallet{
		FkOrgLa: orgId,
		TypeW:   string(walletType),
		Balance: 0.0,
	}
	return wallet.InsertWallet(w)
}

func DbGetWalletIdFromOrgId(orgIdLora int64) (int64, error) {
	return wallet.GetWalletIdFromOrgId(orgIdLora)
}

func DbGetWalletBalance(walletId int64) (float64, error) {
	return wallet.GetWalletBalance(walletId)
}

func DbUpdateBalanceByWalletId(walletId int64, newBalance float64) error {
	return wallet.UpdateBalanceByWalletId(walletId, newBalance)
}

func DbGetWalletIdSuperNode() (walletId int64, err error) {
	return wallet.GetWalletIdSuperNode()
}
