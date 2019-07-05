package db

func DbWalletGetBalanceByWalletId(walletId int64) (float64, error) {
	// get wallet balance with: wallet id
	// only normal user calls DbWalletGetBalanceByWalletId
	// supernode owners call api from supernode pkg
	return 123456.023, nil
}

func DbWalletGetBalanceByOrgId(orgId int64) (float64, error) {
	// get wallet balance with: username, orgId
	// only normal user calls DbWalletGetBalanceByUsernameOrgId
	// supernode owners call api from supernode pkg
	return 123456.023, nil
}

func DbWalletUpdateBalanceByWalletId(walletId int64, newBalance float64) error {
	// select wallet with walletId, update balance with newBalance
	return nil
}
