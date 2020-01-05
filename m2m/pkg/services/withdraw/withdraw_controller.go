package withdraw

import (
	log "github.com/sirupsen/logrus"
	"github.com/mxc-foundation/mxprotocol-server/m2m/db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/config"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/services/wallet"
)

var WithdrawFee map[string]float64

func Setup(conf config.MxpConfig) error {
	log.Info("Setup withdraw service")

	WithdrawFee = make(map[string]float64)
	for _, v := range wallet.CurrencyList {
		WithdrawFee[v.Abv] = 20
	}

	PaymentServiceAvailable = paymentServiceAvailable(conf)
	if false == PaymentServiceAvailable {
		log.Warning("service/withdraw")
	}

	for _, v := range wallet.CurrencyList {
		withdrawFee, err := db.WithdrawFee.GetActiveWithdrawFee(v.Abv)
		if err != nil {
			if _, err := db.WithdrawFee.InsertWithdrawFee(v.Abv, WithdrawFee[v.Abv]); err != nil {
				log.WithError(err).Error("service/withdraw")
				return err
			}
		} else {
			WithdrawFee[v.Abv] = withdrawFee
		}
	}

	return nil
}
