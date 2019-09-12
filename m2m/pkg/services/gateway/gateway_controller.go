package gateway

import (
	"context"
	log "github.com/sirupsen/logrus"
	api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/m2m"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/auth"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Setup() error {
	log.Info("Setup gateway service")
	return nil
}

type GatewayServerAPI struct {
	serviceName string
}

func NewGatewayServerAPI() *GatewayServerAPI {
	return &GatewayServerAPI{serviceName: "gateway"}
}

func (s *GatewayServerAPI) GetGatewayList(ctx context.Context, req *api.GetGatewayListRequest) (*api.GetGatewayListResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)
	case auth.OrganizationIdRearranged:
		return &api.GetGatewayListResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")
	case auth.OK:
		log.WithFields(log.Fields{
			"orgId":     req.OrgId,
			"offset":    req.Offset,
			"limit":     req.Limit,
			"wallet_id": req.WalletId,
		}).Debug("grpc_api/GetGatewayList")

		gwList, err := db.Gateway.GetGatewayListOfWallet(req.WalletId, req.Offset, req.Limit)
		if err != nil {
			log.WithError(err).Error("grpc_api/GetGatewayList")
			return &api.GetGatewayListResponse{UserProfile: &userProfile}, err
		}

		resp := api.GetGatewayListResponse{UserProfile: &userProfile}
		for _, v := range gwList {
			gwProfile := api.GatewayProfile{}
			gwProfile.Id = v.Id
			gwProfile.Mac = v.Mac
			gwProfile.FkGwNs = v.FkGatewayNs
			gwProfile.FkWallet = v.FkWallet
			gwProfile.Mode = string(v.Mode)
			gwProfile.CreateAt = v.CreatedAt.String()
			gwProfile.LastSeenAt = v.LastSeenAt.String()
			gwProfile.OrgId = v.OrgId
			gwProfile.Description = v.Description
			gwProfile.Name = v.Name

			resp.GwProfile = append(resp.GwProfile, &gwProfile)
		}

		return &resp, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *GatewayServerAPI) GetGatewayProfile(ctx context.Context, req *api.GetGatewayProfileRequest) (*api.GetGatewayProfileResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)
	case auth.OrganizationIdRearranged:
		return &api.GetGatewayProfileResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")
	case auth.OK:
		log.WithFields(log.Fields{
			"orgId":  req.OrgId,
			"gwId":   req.GwId,
			"offset": req.Offset,
			"limit":  req.Limit,
		}).Debug("grpc_api/GetGatewayProfile")

		gwProfile, err := db.Gateway.GetGatewayProfile(req.GwId)
		if err != nil {
			log.WithError(err).Error("grpc_api/GetGatewayProfile")
			return &api.GetGatewayProfileResponse{UserProfile: &userProfile}, err
		}

		resp := api.GatewayProfile{
			Id:          gwProfile.Id,
			Mac:         gwProfile.Mac,
			FkGwNs:      gwProfile.FkGatewayNs,
			FkWallet:    gwProfile.FkWallet,
			Mode:        string(gwProfile.Mode),
			CreateAt:    gwProfile.CreatedAt.String(),
			LastSeenAt:  gwProfile.LastSeenAt.String(),
			OrgId:       gwProfile.OrgId,
			Description: gwProfile.Description,
			Name:        gwProfile.Name,
		}

		return &api.GetGatewayProfileResponse{UserProfile: &userProfile, GwProfile: &resp}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *GatewayServerAPI) GetGatewayHistory(ctx context.Context, req *api.GetGatewayHistoryRequest) (*api.GetGatewayHistoryResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)
	case auth.OrganizationIdRearranged:
		return &api.GetGatewayHistoryResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")
	case auth.OK:

	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *GatewayServerAPI) SetGatewayMode(ctx context.Context, req *api.SetGatewayModeRequest) (*api.SetGatewayModeResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)
	case auth.OrganizationIdRearranged:
		return &api.SetGatewayModeResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")
	case auth.OK:
		log.WithFields(log.Fields{
			"orgId": req.OrgId,
			"gwID":  req.GwId,
			"gwMod": req.GwMode,
		}).Debug("grpc_api/SetDeviceMode")

		switch req.GwMode.String() {
		case "GW_INACTIVE":
			if err := db.Gateway.SetGatewayMode(req.GwId, types.GW_INACTIVE); err != nil {
				log.WithError(err).Error("grpc_api/SetDeviceMode")
				return &api.SetGatewayModeResponse{Status: false, UserProfile: &userProfile}, err
			}
		case "GW_FREE_GATEWAYS_LIMITED":
			if err := db.Gateway.SetGatewayMode(req.GwId, types.GW_FREE_GATEWAYS_LIMITED); err != nil {
				log.WithError(err).Error("grpc_api/SetDeviceMode")
				return &api.SetGatewayModeResponse{Status: false, UserProfile: &userProfile}, err
			}
		case "GW_WHOLE_NETWORK":
			if err := db.Gateway.SetGatewayMode(req.GwId, types.GW_WHOLE_NETWORK); err != nil {
				log.WithError(err).Error("grpc_api/SetDeviceMode")
				return &api.SetGatewayModeResponse{Status: false, UserProfile: &userProfile}, err
			}
		case "GW_DELETED":
			if err := db.Gateway.SetGatewayMode(req.GwId, types.GW_DELETED); err != nil {
				log.WithError(err).Error("grpc_api/SetDeviceMode")
				return &api.SetGatewayModeResponse{Status: false, UserProfile: &userProfile}, err
			}
		}

		return &api.SetGatewayModeResponse{UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}
