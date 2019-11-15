package ui

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/m2m_ui"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/auth"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var timeLayout = "2006-01-02 15:04:05"

type StakingServerAPI struct {
	serviceName string
}

// StakingServerAPI returns a new StakingServerAPI.
func NewStakingServerAPI() *StakingServerAPI {
	return &StakingServerAPI{serviceName: "staking"}
}

func (s *StakingServerAPI) GetStakingPercentage(ctx context.Context, req *api.StakingPercentageRequest) (*api.StakingPercentageResponse, error) {
	stakingPercentage := config.Cstruct.Staking.StakingPercentage
	return &api.StakingPercentageResponse{StakingPercentage: stakingPercentage}, nil
}

func (s *StakingServerAPI) Stake(ctx context.Context, req *api.StakeRequest) (*api.StakeResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.StakeResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:
		log.WithFields(log.Fields{
			"orgId": req.OrgId,
		}).Debug("grpc_api/Stake")
		if req.Amount <= 0 {
			return &api.StakeResponse{Status: "Staking amount must be more than 0.", UserProfile: &userProfile}, nil
		}

		walletId, err := db.Wallet.GetWalletIdFromOrgId(req.OrgId)
		if err != nil {
			log.WithError(err).Error("StakeAPI/Cannot get walletId from DB")
		}

		// check if person has enough balance for staking
		personalBalance, err := db.Wallet.GetWalletBalance(walletId)
		if req.Amount > personalBalance {
			return &api.StakeResponse{Status: "You do not have enough money.", UserProfile: &userProfile}, nil
		}

		stakeProf, err := db.Stake.GetActiveStake(walletId)
		if err != nil {
			log.WithError(err).Error("StakeAPI/Cannot get staking profile from DB")
		}

		//If this person has one staking in DB already, return
		var nilStake = types.Stake{}
		if stakeProf != nilStake {
			return &api.StakeResponse{Status: "There is already one active stake, you should do the unstake first.", UserProfile: &userProfile}, nil
		}

		//add the stake value to DB
		_, err = db.Stake.InsertStake(walletId, req.Amount)
		if err != nil {
			log.WithError(err).Error("StakeAPI/Cannot insert new stake to DB")
		}

		return &api.StakeResponse{Status: "Stake successful.", UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "Internal error")
}

func (s *StakingServerAPI) Unstake(ctx context.Context, req *api.UnstakeRequest) (*api.UnstakeResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.UnstakeResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:
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
			return &api.UnstakeResponse{Status: "There is no active stake.", UserProfile: &userProfile}, nil
		}

		//get the min day from config, and compare if already longer than the min day.
		minStakeDays := config.Cstruct.Staking.StakingMinDays

		//check if it's longer than minStakeDays
		period := time.Now().UTC().Sub(stakeProf.StartStakeTime).Hours()

		if (period / 24) < float64(minStakeDays) {
			status := fmt.Sprintf("The minimum unstake period is %v days", minStakeDays)
			return &api.UnstakeResponse{Status: status,
				UserProfile: &userProfile}, nil
		}

		//update unstake time and status to DB.
		err = db.Stake.Unstake(stakeProf.Id)
		if err != nil {
			log.WithError(err).Error("StakeAPI/Cannot update unstake to DB")
		}

		return &api.UnstakeResponse{Status: "Unstake successful.", UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "Internal error")
}

func (s *StakingServerAPI) GetActiveStakes(ctx context.Context, req *api.GetActiveStakesRequest) (*api.GetActiveStakesResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.GetActiveStakesResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:
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
			return &api.GetActiveStakesResponse{UserProfile: &userProfile}, nil
		}

		actStake := &api.ActiveStake{}
		actStake.Id = stakeProf.Id
		actStake.FkWallet = stakeProf.FkWallet
		actStake.Amount = stakeProf.Amount
		actStake.StakeStatus = string(stakeProf.Status)
		actStake.StartStakeTime = stakeProf.StartStakeTime.String()
		actStake.UnstakeTime = stakeProf.UnstakeTime.String()

		return &api.GetActiveStakesResponse{ActStake: actStake, UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "Internal error")
}

func (s *StakingServerAPI) GetStakingHistory(ctx context.Context, req *api.StakingHistoryRequest) (*api.StakingHistoryResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.StakingHistoryResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:
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
			stakeHist.End = v.UnstakeTime.String()
			//get the month from start time
			stakingRevMonth := v.StartStakeTime.Month().String()
			stakeHist.RevMonth = stakingRevMonth
			stakeHist.NetworkIncome = v.SuperNodeIncome
			stakeHist.MonthlyRate = v.IncomeToStakePortion * 100
			stakeHist.Revenue = v.RevenueAmount
			stakeHist.Balance = v.UpdatedBalance

			resp.StakingHist = append(resp.StakingHist, stakeHist)
		}

		resp.Count = totalHists
		resp.UserProfile = &userProfile

		return resp, nil
	}

	return nil, status.Errorf(codes.Unknown, "Internal error")
}
