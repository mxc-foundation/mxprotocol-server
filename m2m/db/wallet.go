package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

type WalletType string

const (
	USER        WalletType = "USER"
	SUPER_ADMIN WalletType = "SUPER_ADMIN"
)

func dbCreateWalletTable() error {
	return pg.PgDB.CreateWalletTable()
}

func DbInsertWallet(orgId int64, walletType WalletType) (insertIndex int64, err error) {
	w := pg.Wallet{
		FkOrgLa: orgId,
		TypeW:   string(walletType),
		Balance: 0.0,
	}
	return pg.PgDB.InsertWallet(w)
}

func DbGetWalletIdFromOrgId(orgIdLora int64) (int64, error) {
	return pg.PgDB.GetWalletIdFromOrgId(orgIdLora)
}

func DbGetWalletBalance(walletId int64) (float64, error) {
	return pg.PgDB.GetWalletBalance(walletId)
}

func DbUpdateBalanceByWalletId(walletId int64, newBalance float64) error {
	return pg.PgDB.UpdateBalanceByWalletId(walletId, newBalance)
}

func DbGetWalletIdSuperNode() (walletId int64, err error) {
	return pg.PgDB.GetWalletIdSuperNode()
}
