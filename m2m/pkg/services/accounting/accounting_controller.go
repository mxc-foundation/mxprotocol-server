package accounting

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
)

func Setup(conf config.MxpConfig) error {

	log.Info("*********  Aggregation started log")
	log.Error("*********  Aggregation started error")
	log.Warn("*********  Aggregation started warning")
	// log.WithError(err).Warning("service/supernode")

	var aggDurationMinutes int64 = 2 * 60 //48 * 60 // To Do: should be received from the config file

	// calling accounting routine based on time trigger called here

	testAccounting(time.Now().UTC(), 60*72, 0.1) ///@@ delte/ for test

	tickerAccounting := time.NewTicker(time.Duration(aggDurationMinutes) * time.Minute) // @@ change to minute
	go func() {
		for range tickerAccounting.C {
			execTime := time.Now().UTC()
			log.Info("Accounting routine started")
			var dlPrice float64 = 0.01 // To Do: should be received from the config file
			// super node  // func (*walletInterface) GetWalletIdSuperNode() (walletId int64, err error) {
			testAccounting(execTime, aggDurationMinutes, dlPrice)

		}
	}()

	log.Info("setup accounting service")

	return nil
}
