package staking

import (
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
)

func Setup(conf config.MxpConfig) error {
	log.Info("cron task begin...")
	c := cron.New()
	err := c.AddFunc("0011*?", func() {
		log.Info("Start stakingReveneuExec")
		err := stakingReveneuExec(conf)
		if err != nil {
			log.WithError(err).Error("StakingReveneu Error")
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	c.Start()

	return nil
}

func stakingReveneuExec(conf config.MxpConfig) error {
	//Todo: get income from DB
	income := 12345.1232

	//Todo: get amount of stakes from DB
	people := []*int{}
	for i := range people {
		//Todo: get staking value from DB
		stakingValue := 123.56789

		//Todo: calculate the real days
		stakingTimePortion := 28/30

		//Todo: finish the fun
		revenue := (income * conf.Staking.StakingPercentage) * (stakingValue * float64(stakingTimePortion))
	}

	return nil
}
