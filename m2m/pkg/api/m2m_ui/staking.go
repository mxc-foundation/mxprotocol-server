package ui

import (
	"context"
	log "github.com/sirupsen/logrus"
	api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/m2m_ui"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
		//Todo: check if this person already has staking in DB.
		//getActiveStake(walletId) nil or not nil

		/*//Todo: get the min staking time
		minDay := config.MxpConfig{}.Staking.StakingMinDays*/

		//Todo: write the stake value to DB
		//addStake(walletId, amount)

		return &api.StakeResponse{UserProfile: &userProfile}, nil
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
		//Todo: get the start date from DB
		//getActiveStake()

		//Todo: get the min day from config, and compare if already longer than the min day.
		//year, month, today := time.Now().Date()

		//Todo: update unstake_time and status to DB.
		//unstake(stakeID)

		return &api.UnstakeResponse{UserProfile: &userProfile}, nil
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