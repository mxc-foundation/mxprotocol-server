package accounting

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
)

func Setup(conf config.MxpConfig) error {

	var aggDurationMinutes int64 = conf.Accounting.IntervalMin

	performAccounting(time.Now().UTC(), 60*72, 0.1) ///@@ for test

	tickerAccounting := time.NewTicker(time.Duration(aggDurationMinutes) * time.Minute)
	go func() {
		for range tickerAccounting.C {
			execTime := time.Now().UTC()
			log.Info("Accounting routine started")
			var dlPrice float64 = conf.SuperNode.DlPrice
			fmt.Println("dlPrice:", dlPrice)

			if err := performAccounting(execTime, aggDurationMinutes, dlPrice); err != nil {
				log.WithError(err).Error(" Accounting Failed! ")
			}
		}
	}()

	log.Info("setup accounting service")

	return nil
}
