package appserver

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	api "github.com/mxc-foundation/mxprotocol-server/m2m/api/m2m_ui"
	"github.com/mxc-foundation/mxprotocol-server/m2m/db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/config"
	"github.com/mxc-foundation/mxprotocol-server/m2m/types"
	"time"
)

var timeLayout = "2006-01-02 15:04:05"

func (s *M2MServerAPI) GetStakingPercentage(ctx context.Context, req *api.StakingPercentageRequest) (*api.StakingPercentageResponse, error) {
	stakingPercentage := config.Cstruct.Staking.StakingPercentage
	return &api.StakingPercentageResponse{StakingPercentage: stakingPercentage}, nil
}

func (s *M2MServerAPI) Stake(ctx context.Context, req *api.StakeRequest) (*api.StakeResponse, error) {
	log.WithFields(log.Fields{
		"orgId": req.OrgId,
	}).Debug("grpc_api/Stake")
	if req.Amount <= 0 {
		return &api.StakeResponse{Status: "Staking amount must be more than 0."}, nil
	}

	walletId, err := db.Wallet.GetWalletIdFromOrgId(req.OrgId)
	if err != nil {
		log.WithError(err).Error("StakeAPI/Cannot get walletId from DB")
	}

	// check if person has enough balance for staking
	personalBalance, err := db.Wallet.GetWalletBalance(walletId)
	if req.Amount > personalBalance {
		return &api.StakeResponse{Status: "You do not have enough money."}, nil
	}

	stakeProf, err := db.Stake.GetActiveStake(walletId)
	if err != nil {
		log.WithError(err).Error("StakeAPI/Cannot get staking profile from DB")
	}

	//If this person has one staking in DB already, return
	var nilStake = types.Stake{}
	if stakeProf != nilStake {
		return &api.StakeResponse{Status: "There is already one active stake, you should do the unstake first."}, nil
	}

	//add the stake value to DB
	_, err = db.Stake.InsertStake(walletId, req.Amount)
	if err != nil {
		log.WithError(err).Error("StakeAPI/Cannot insert new stake to DB")
	}

	return &api.StakeResponse{Status: "Stake successful."}, nil
}

func (s *M2MServerAPI) Unstake(ctx context.Context, req *api.UnstakeRequest) (*api.UnstakeResponse, error) {
	log.WithFields(log.Fields{
		"orgId": req.OrgId,
	}).Debug("grpc_api/Unstake")
	walletID, err := db.Wallet.GetWalletIdFromOrgId(req.OrgId)
	if err != nil {
		log.WithError(err).Error("UnstakeAPI/Cannot get walletID from DB")
	}

	//get the start date from stakeProf
	stakeProf, err := db.Stake.GetActiveStake(walletID)

	//If this person has one staking in DB already, return.
	var nilStake = types.Stake{}
	if stakeProf == nilStake {
		return &api.UnstakeResponse{Status: "There is no active stake."}, nil
	}

	//get the min day from config, and compare if already longer than the min day.
	minStakeDays := config.Cstruct.Staking.StakingMinDays

	//check if it's longer than minStakeDays
	period := time.Now().UTC().Sub(stakeProf.StartStakeTime).Hours()

	if (period / 24) < float64(minStakeDays) {
		status := fmt.Sprintf("The minimum unstake period is %v days", minStakeDays)
		return &api.UnstakeResponse{Status: status}, nil
	}

	//update unstake time and status to DB.
	err = db.Stake.Unstake(stakeProf.Id)
	if err != nil {
		log.WithError(err).Error("StakeAPI/Cannot update unstake to DB")
	}

	return &api.UnstakeResponse{Status: "Unstake successful."}, nil
}

func (s *M2MServerAPI) GetActiveStakes(ctx context.Context, req *api.GetActiveStakesRequest) (*api.GetActiveStakesResponse, error) {
	log.WithFields(log.Fields{
		"orgId": req.OrgId,
	}).Debug("grpc_api/GetActiveStakes")
	walletId, err := db.Wallet.GetWalletIdFromOrgId(req.OrgId)
	if err != nil {
		log.WithError(err).Error("GetActiveStakes/Cannot get walletID from DB")
	}

	stakeProf, err := db.Stake.GetActiveStake(walletId)
	if err != nil {
		log.WithError(err).Error("GetActiveStakes/Cannot get staking profile from DB")
	}

	//var nilStake = types.Stake{}
	if stakeProf.Status == "" {
		return &api.GetActiveStakesResponse{}, nil
	}

	actStake := &api.ActiveStake{}
	actStake.Id = stakeProf.Id
	actStake.FkWallet = stakeProf.FkWallet
	actStake.Amount = stakeProf.Amount
	actStake.StakeStatus = string(stakeProf.Status)
	actStake.StartStakeTime = stakeProf.StartStakeTime.String()
	actStake.UnstakeTime = stakeProf.UnstakeTime.String()

	return &api.GetActiveStakesResponse{ActStake: actStake}, nil
}

func (s *M2MServerAPI) GetStakingHistory(ctx context.Context, req *api.StakingHistoryRequest) (*api.StakingHistoryResponse, error) {
	log.WithFields(log.Fields{
		"orgId": req.OrgId,
	}).Debug("grpc_api/GetStakingHistory")
	walletId, err := db.Wallet.GetWalletIdFromOrgId(req.OrgId)
	if err != nil {
		log.WithError(err).Error("GetStakingHistory/Cannot get walletID from DB")
	}

	stakingHists, err := db.StakeRevenue.GetStakeRevenueHistory(walletId, req.Offset, req.Limit)
	if err != nil {
		log.WithError(err).Error("GetStakingHistory/Cannot get histories from DB")
	}

	totalHists, err := db.StakeRevenue.GetStakeRevenueHistoryCnt(walletId)
	if err != nil {
		log.WithError(err).Error("GetStakingHistory/Cannot get total numbers of histories from DB")
	}

	resp := &api.StakingHistoryResponse{}

	for _, v := range stakingHists {
		stakeHist := &api.GetStakingHistory{}
		stakeHist.StakeAmount = v.StakeAmount
		stakeHist.Start = v.StartStakeTime.String()

		if v.UnstakeTime.String() == "0001-01-01 00:00:00 +0000 +0000" {
			stakeHist.End = "--:--"
		} else {
			stakeHist.End = v.UnstakeTime.String()
		}

		if v.StakingPeriodStart.Month() != v.StakingPeriodEnd.Month() {
			log.WithError(err).Error("GetStakingHistory/StakingPeriodStart and End in different month.")
		}

		if v.StakingPeriodStart.String() == "0001-01-01 00:00:00 +0000 UTC" || v.StakingPeriodEnd.String() == "0001-01-01 00:00:00 +0000 UTC" {
			stakeHist.RevMonth = "No Revenue"
		} else {
			stakeHist.RevMonth = v.StakingPeriodStart.Month().String()
		}

		stakeHist.NetworkIncome = v.SuperNodeIncome
		stakeHist.MonthlyRate = v.IncomeToStakePortion * 100
		stakeHist.Revenue = v.RevenueAmount
		stakeHist.Balance = v.UpdatedBalance

		resp.StakingHist = append(resp.StakingHist, stakeHist)
	}

	resp.Count = totalHists

	return resp, nil
}
