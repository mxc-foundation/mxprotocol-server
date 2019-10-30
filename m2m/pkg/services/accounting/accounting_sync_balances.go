package accounting

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
)

func syncTmpBalance(walletId int64) {

	oldTmpBalance, errGetBal := db.Wallet.GetWalletBalance(walletId)
	if errGetBal != nil {
		log.WithFields(log.Fields{
			"walletId": walletId,
		}).WithError(errGetBal).Warn("accounting/syncTmpBalance unable to get tmpBalance ")
	}

	updatedBalance, err := db.Wallet.SyncTmpBalance(walletId)

	if err != nil {
		log.WithFields(log.Fields{
			"walletId": walletId,
		}).WithError(err).Warn("accounting/syncTmpBalance ")
	}

	if updatedBalance != oldTmpBalance {
		log.WithFields(log.Fields{
			"walletId":    walletId,
			"balance":     updatedBalance,
			"tmp_balance": oldTmpBalance,
		}).Warn("accounting/syncTmpBalance tmp Balance and Balance differs!")
	}

	if updatedBalance < 0 {
		log.WithFields(log.Fields{
			"walletId": walletId,
			"balance":  updatedBalance,
		}).Warn("accounting/syncTmpBalance Account balance is NEGATIVE!")
	}

}
