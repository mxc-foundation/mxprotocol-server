package accounting

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
)

func Setup(conf config.MxpConfig) error {

	go func() {
		// set the first accounting time to be run at the closest hh:00'
		now := time.Now().UTC()
		nextTickTimeDiff := now.Truncate(time.Hour).Add(time.Hour).Sub(now)
		t := time.NewTicker(nextTickTimeDiff)
		superNodePktSentIncomeRatio := 0.08 // TODO: should be received from configs
		for {
			// Wait for tick from the ticker
			<-t.C

			var dlPrice float64 = conf.SuperNode.DlPrice
			var aggDurationMinutes int64 = conf.Accounting.IntervalMin

			if err := performAccounting(aggDurationMinutes, dlPrice, superNodePktSentIncomeRatio); err != nil {
				log.WithError(err).Error(" Accounting Failed! ")
			}

			now := time.Now().UTC()
			if aggDurationMinutes%60 == 0 {
				nextTickTimeDiff = now.Truncate(time.Hour).Add(time.Minute * time.Duration(aggDurationMinutes)).Sub(now)
			} else {
				nextTickTimeDiff = time.Minute * time.Duration(aggDurationMinutes)
			}
			t = time.NewTicker(nextTickTimeDiff)

		}
	}()

	log.Info("setup accounting service")

	return nil
}
