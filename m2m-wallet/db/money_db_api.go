package db

import "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"

type MoneyType api.Money

const (
	Ether = MoneyType(api.Money_Ether)
)

func DbMoneyUpdateAccountByWalletIdMoneyType(walletId int64, newAccount string, mType MoneyType) error {
	// select external account with: walletId and moneyType, update account_addr with newAccount
	return nil
}
