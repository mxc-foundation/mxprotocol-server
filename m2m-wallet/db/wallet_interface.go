package db

import (
	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

func DbCreateWalletTable() {
	pgDb.CreateWalletTable()
}

func DbInsertWallet(w pstgDb.Wallet) {
	pgDb.InsertWallet(w)
}

func DbGetWalletId(orgIdLora int) int {
	return pgDb.GetWalletId(orgIdLora)
}

func DbGetWallet(wp *pstgDb.Wallet, orgIdLora int) error {
	return pgDb.GetWallet(wp, orgIdLora)
}
