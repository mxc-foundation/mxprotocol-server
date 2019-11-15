package ext_account

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services/wallet"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

func Setup() error {
	log.Info("Setup ext_account service")
	return nil
}

func UpdateActiveExtAccount(orgId int64, newAccount string, currencyAbbr string) error {
	walletId, err := wallet.GetWalletId(orgId)
	if err != nil {
		log.WithError(err).Error("service/UpdateActiveExtAccount")
		return err
	}

	_, err = db.ExtAccount.InsertExtAccount(walletId, newAccount, currencyAbbr)
	if err != nil {
		log.WithError(err).Error("service/UpdateActiveExtAccount")
		return err
	}

	return nil
}

func GetActiveExtAccount(orgId int64, currencyAbbr string) (string, error) {
	walletId, err := wallet.GetWalletId(orgId)
	if err != nil {
		log.WithError(err).Error("service/GetActiveExtAccount")
		return "", err
	}

	var accountAddr string
	if orgId == 0 {
		accountAddr, err = db.ExtAccount.GetSuperNodeExtAccountAdr(currencyAbbr)
	} else {
		accountAddr, err = db.ExtAccount.GetUserExtAccountAdr(walletId, currencyAbbr)
	}

	if err != nil {
		if strings.HasSuffix(err.Error(), types.DbError.NoRowQueryRes.Error()) {
			log.Warnf("service/GetActiveExtAccount: get account with walletId=%d, currency=%s", walletId, currencyAbbr)
			return "", nil
		}
		log.WithError(err).Error("service/GetActiveExtAccount")
		return "", err
	}

	return accountAddr, nil
}
