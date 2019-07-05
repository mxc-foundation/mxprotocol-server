package db

import (
	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

func DbCreateWalletTable() error {
	return pgDb.CreateWalletTable()
}

func DbInsertWallet(w pstgDb.Wallet) (insertIndex int, err error) {
	return pgDb.InsertWallet(w)
}

func DbGetWalletIdFromOrgId(orgIdLora int) (int, error) {
	return pgDb.GetWalletIdFromOrgId(orgIdLora)
}

func DbGetWallet(wp *pstgDb.Wallet, walletId int) error {
	return pgDb.GetWallet(wp, walletId)
}

func DbGetWalletBalance(walletId int) (float64, error) {
	return pgDb.GetWalletBalance(walletId)
}
