package ui

import (
	"context"
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

type StakingServerAPI struct{
	serviceName string
}

// StakingServerAPI returns a new StakingServerAPI.
func NewStakingServerAPI() *StakingServerAPI {
	return &StakingServerAPI{serviceName: "staking"}
}

func (s *StakingServerAPI) Stake (ctx context.Context, req *api.StakeRequest) (*api.StakeResponse, error){
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
		walletId, err := db.Wallet.GetWalletIdFromOrgId(req.OrgId)
		if err != nil {
			log.WithError(err).Error("StakeAPI/Cannot get walletId from DB")
		}

		stakeProf, err := db.Stake.GetActiveStake(walletId)
		if err != nil {
			log.WithError(err).Error("StakeAPI/Cannot get staking profile from DB")
		}

		//If this person has one staking in DB already, return.
		var nilStake = types.Stake{}
		if stakeProf != nilStake {
			return &api.StakeResponse{Status:"There is already one active stake, you should do the unstake first.", UserProfile: &userProfile}, nil
		}

		//add the stake value to DB
		_, err = db.Stake.InsertStake(walletId, req.Amount)
		if err != nil {
			log.WithError(err).Error("StakeAPI/Cannot insert new stake to DB")
		}

		return &api.StakeResponse{Status:"Stake successful.", UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "Internal error")
}

func (s *StakingServerAPI) Unstake (ctx context.Context, req *api.UnstakeRequest) (*api.UnstakeResponse, error) {
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
			return &api.UnstakeResponse{Status:"There is no active stake.", UserProfile: &userProfile}, nil
		}

		startTime, err := time.Parse(timeLayout, stakeProf.StartStakeTime.String())
		if err != nil {
			log.WithError(err).Error("startTime time format error")
		}

		//get the min day from config, and compare if already longer than the min day.
		minStakeDays := config.MxpConfig{}.Staking.StakingMinDays

		now, err := time.Parse(timeLayout, time.Now().String())
		if err != nil {
			log.WithError(err).Error("time.now format error")
		}

		//check if it's longer than minStakeDays
		//Todo: Change the org
		period := now.Sub(startTime).Hours() //startTime.Sub(now)
		if (period/24) < float64(minStakeDays) {
			return &api.UnstakeResponse{Status:"The minimum unstake period is " + string(minStakeDays) + " days."}, nil
		}

		//update unstake time and status to DB.
		err = db.Stake.Unstake(stakeProf.Id)
		if err != nil {
			log.WithError(err).Error("StakeAPI/Cannot update unstake to DB")
		}

		return &api.UnstakeResponse{Status:"Unstake successful." ,UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "Internal error")
}

func (s *StakingServerAPI) GetStakingHistory (ctx context.Context, req *api.StakingHistoryRequest) (*api.StakingHistoryResponse, error){
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


		return &api.StakingHistoryResponse{UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "Internal error")
}