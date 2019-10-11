package accounting

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
)

func Setup(conf config.MxpConfig) error {

	var aggDurationMinutes int64 = 2 * 60 // To Do: should be received from the config file

	tickerAccounting := time.NewTicker(time.Duration(aggDurationMinutes) * time.Minute)
	go func() {
		for range tickerAccounting.C {
			execTime := time.Now().UTC()
			log.Info("Accounting routine started")
			var dlPrice float64 = 0.01 // To Do: should be received from the config file
			if err := performAccounting(execTime, aggDurationMinutes, dlPrice); err != nil {
				log.WithError(err).Error(" Accounting Failed! ")
			}
		}
	}()

	log.Info("setup accounting service")

	return nil
}
