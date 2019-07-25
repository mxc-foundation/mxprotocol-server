package db

import (
	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

type WalletType string // db:wallet_type

const (
	USER        WalletType = "USER"
	SUPER_ADMIN WalletType = "SUPER_ADMIN"
)

func dbCreateWalletTable() error {
	return pgDb.CreateWalletTable()
}

func DbInsertWallet(orgId int64, walletType WalletType) (insertIndex int64, err error) {
	w := pstgDb.Wallet{
		FkOrgLa: orgId,
		TypeW:   string(walletType),
		Balance: 0.0,
	}
	return pgDb.InsertWallet(w)
}

func DbGetWalletIdFromOrgId(orgIdLora int64) (int64, error) {
	return pgDb.GetWalletIdFromOrgId(orgIdLora)
}

func DbGetWalletBalance(walletId int64) (float64, error) {
	return pgDb.GetWalletBalance(walletId)
}

func DbUpdateBalanceByWalletId(walletId int64, newBalance float64) error {

	return pgDb.UpdateBalanceByWalletId(walletId, newBalance)
}

func DbGetWalletIdSuperNode() (walletId int64, err error) {
	return pgDb.GetWalletIdSuperNode()
}
