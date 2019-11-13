package wallet

import (
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/m2m_ui"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

func Setup(conf config.MxpConfig) error {
	log.Info("setup wallet service")
	log.Info("initialize supported extCurrency")

	if err := initExtCurrencyTable(); err != nil {
		return err
	}

	if err := initControllingWallets(); err != nil {
		return err
	}

	return nil
}

var CurrencyList = []types.ExtCurrency{}

func initExtCurrencyTable() error {
	currency := types.ExtCurrency{}
	for _, v := range api.Money_name {
		if _, err := db.ExtCurrency.GetExtCurrencyIdByAbbr(v); err == nil {
			continue
		}

		currency.Id = 0
		currency.Name = v
		currency.Abv = v
		CurrencyList = append(CurrencyList, currency)
	}

	for _, element := range CurrencyList {
		if _, err := db.ExtCurrency.InsertExtCurr(element); err != nil {
			return err
		}
	}
	return nil
}

func initControllingWallets() error {

	// SUPER_ADMIN wallet is separately initialized
	if _, errIns := db.Wallet.InsertNodeIncomeWallet(); errIns != nil {
		return errors.Wrap(errIns, "wallet_controller/initControllingWallets/ ")
	}

	if _, errIns := db.Wallet.InsertStakeStorageWallet(); errIns != nil {
		return errors.Wrap(errIns, "wallet_controller/initControllingWallets/ ")
	}

	return nil
}

//  return option 1: 0, true        --> no wallet created yet
//  return option 2: 0, false       --> sql error
//  return option 3: walletId, true --> get walletId successfully
func userHasWallet(orgId int64) (int64, bool) {
	walletId, err := db.Wallet.GetWalletIdFromOrgId(orgId)
	if err != nil {
		if strings.HasSuffix(err.Error(), types.DbError.NoRowQueryRes.Error()) {
			return 0, true
		}

		return 0, false
	}

	return walletId, true
}

func createWallet(orgId int64) (walletId int64, err error) {

	if 0 == orgId {
		walletId, err = db.Wallet.InsertWallet(orgId, types.SUPER_ADMIN)
	} else {
		walletId, err = db.Wallet.InsertWallet(orgId, types.USER)
	}
	if err != nil {
		return walletId, err
	}

	return walletId, nil
}

func GetWalletId(orgId int64) (walletId int64, err error) {
	var res bool

	walletId, res = userHasWallet(orgId)
	if true == res && 0 == walletId {
		if walletId, err = createWallet(orgId); err != nil {
			return 0, err
		}
	} else if false == res {
		err = errors.New("Failed to get walletId.")
		log.WithError(err).Error("pkg/wallet/GetWalletId")
		return 0, err
	}

	return walletId, nil
}

func GetBalance(orgId int64) (float64, error) {
	walletId, err := GetWalletId(orgId)
	if err != nil {
		return 0, err
	}

	balance, err := db.Wallet.GetWalletBalance(walletId)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func UpdateBalance(orgId int64, oper PaymentCategory, deviceType DeviceType, amount float64) error {
	walletId, err := GetWalletId(orgId)
	if err != nil {
		return err
	}

	balance, err := db.Wallet.GetWalletBalance(walletId)
	if err != nil {
		return err
	}

	for _, v := range operMap {
		if v.pc == oper && v.dt == deviceType {
			balance = v.operation(balance, amount)
		}
	}

	// err = db.Wallet.UpdateBalanceByWalletId(walletId, balance)
	if err != nil {
		return err
	}

	return nil
}
