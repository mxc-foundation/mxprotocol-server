package staking

import (
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"time"
)

func Setup(conf config.MxpConfig) error {
	log.Info("cron task begin...")
	c := cron.New()
	//first day of the month at 1pm
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
	//Todo: get income from DB, since is one month ago.
	//superNodeIncome (since time.Time, until time.Time)
	income := 12345.1232

	//ToDo: call the query to start the rev and get the ID
	//insertStakeReveneuPeriod(sumIncome, stakingIncomePercentage, stakingPeriodStart, stakingperiodEnd) (stakeRevPerId)

	//Todo: get amount of stakes from DB
	//getActStakes()
	stakes := []*int{}

	var totalPortion float64

	for j := range stakes {
		//j.amount * stakingTimePortion (everyone is different)
		//		totalPortion := stakingValue * stakingTimePortion
	}

	for i := range stakes {
		//Todo: get staking value from DB
		stakingValue := 123.56789
		log.WithError().Fatal()
		log.Fatal()

		//Todo: calculate the real days, if staking_days >= staking_revenue_days, stakingTimePortion = 1
		//get the staking start day from staking struct
		stakingTimePortion := float64(28/30)

		//Todo: finish the fun
		revenue := (income * conf.Staking.StakingPercentage) * (stakingValue * float64(stakingTimePortion)) / totalPortion

		//Todo: update revenue to DB.
		//insertStakeReveneu (stakingId, stakeReveneuPeriodId, revenue)
	}

	//Todo: when all the process finished, give the time to DB.
	//updateStakeReveneuPeriodComplete(stakeReveneuPeriodId)

	return nil
}
