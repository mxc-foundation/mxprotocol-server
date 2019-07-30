package db

type WalletType string

const (
	USER        WalletType = "USER"
	SUPER_ADMIN WalletType = "SUPER_ADMIN"
)

func dbCreateWalletTable() error {
	return dbHandler.CreateWalletTable()
}

func DbInsertWallet(orgId int64, walletType WalletType) (insertIndex int64, err error) {
	w := Wallet{
		FkOrgLa: orgId,
		TypeW:   string(walletType),
		Balance: 0.0,
	}
	return dbHandler.InsertWallet(w)
}

func DbGetWalletIdFromOrgId(orgIdLora int64) (int64, error) {
	return dbHandler.GetWalletIdFromOrgId(orgIdLora)
}

func DbGetWalletBalance(walletId int64) (float64, error) {
	return dbHandler.GetWalletBalance(walletId)
}

func DbUpdateBalanceByWalletId(walletId int64, newBalance float64) error {
	return dbHandler.UpdateBalanceByWalletId(walletId, newBalance)
}

func DbGetWalletIdSuperNode() (walletId int64, err error) {
	return dbHandler.GetWalletIdSuperNode()
}
