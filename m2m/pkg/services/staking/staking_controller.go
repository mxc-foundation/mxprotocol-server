package staking

import (
	"fmt"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"time"
)

func Setup(conf config.MxpConfig) error {
	log.Info("cron task begin...")
	c := cron.New()
	//first day of the month at 12:00 (24-hour)
	err := c.AddFunc("0 0 12 1 * ?", func() {
		log.Info("Start stakingRevenueExec")
		go func() {
			err := stakingRevenueExec(conf)
			if err != nil {
				log.WithError(err).Error("StakingRevenue Error")
			}
		}()
	})
	if err != nil {
		log.Fatal(err)
	}

	go c.Start()

	return nil
}

func stakingRevenueExec(conf config.MxpConfig) error {
	//get income from DB, since is one month ago.
	//superNodeIncome (since time.Time, until time.Time)
	//Todo: edit this prat later
	income := 10000.00

	t := time.Now()
	//first date of month 00:00:00
	startTime := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	//last date of month 23:59:59
	endTime := startTime.AddDate(0, 1, 0).Add(time.Second * -1)

	// how many days in this month
	lastDate := startTime.AddDate(0, 1, 0)
	stakingRevDays := lastDate.Sub(startTime).Hours() / 24

	//call the query to start the rev and get the stake revenue ID
	stakeRevPeriodId, err := db.StakeRevenuePeriod.InsertStakeRevenuePeriod(startTime, endTime, income, conf.Staking.StakingPercentage)
	if err != nil {
		log.WithError(err).Error("stakingRevenueExec/Cannot get stakeRevPeriodId from DB")
	}

	//get amount of stakes from DB
	stakes, err := db.Stake.GetActiveStakes()
	if err != nil {
		log.WithError(err).Error("stakingRevenueExec/Cannot get stakes from DB")
	}

	var totalPortion float64

	for _, i := range stakes {

		var stakingTimePortion float64
		//how many hours from startTime until now
		stakingHours := time.Now().Sub(i.StartStakeTime).Hours()
		if (stakingHours / 24) >= stakingRevDays {
			stakingTimePortion = 1
		} else {
			stakingTimePortion = (stakingHours / 24) / stakingRevDays
		}
		//sum up denominator
		totalPortion += i.Amount * stakingTimePortion
	}

	for _, j := range stakes {
		var stakingTimePortion float64
		//how many hours from startTime until now
		stakingHours := time.Now().Sub(j.StartStakeTime).Hours()
		if (stakingHours / 24) >= stakingRevDays {
			stakingTimePortion = 1
		} else {
			stakingTimePortion = (stakingHours / 24) / stakingRevDays
		}

		//how much revenue this person should get.
		revenueAmount := (income * conf.Staking.StakingPercentage) * (j.Amount * stakingTimePortion) / totalPortion

		//update revenue to DB.
		_, err := db.StakeRevenue.InsertStakeRevenue(j.Id, stakeRevPeriodId, revenueAmount)
		if err != nil {
			log.WithError(err).Error("stakingRevenueExec/Cannot update revenue to DB")
		}

		//Todo: only for the testing
		fmt.Println("stakingDays: ", stakingHours/24)
		fmt.Println("revenueAmount: ", revenueAmount)
	}

	//when all the process finished, give the time to DB.
	if err := db.StakeRevenuePeriod.UpdateCompletedStakeReveneuPeriod(stakeRevPeriodId); err != nil {
		log.WithError(err).Error("stakingRevenueExec/Cannot update revenueTime to DB")
	}

	//Todo: only for the testing
	fmt.Println("startTime: ", startTime)
	fmt.Println("endTime: ", endTime)
	fmt.Println("lastDate: ", lastDate)
	fmt.Println("stakingRevDays: ", stakingRevDays)
	fmt.Println("TotalProtaion: ", totalPortion)

	return nil
}
