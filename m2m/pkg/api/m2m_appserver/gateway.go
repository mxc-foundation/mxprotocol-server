package appserver

import (
	"context"

	log "github.com/sirupsen/logrus"
	api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/appserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

func (s *M2MServerAPI) GetGatewayList(ctx context.Context, req *api.GetGatewayListRequest) (*api.GetGatewayListResponse, error) {
	log.WithFields(log.Fields{
		"orgId":  req.OrgId,
		"offset": req.Offset,
		"limit":  req.Limit,
	}).Debug("grpc_api/GetGatewayList")

	walletId, err := db.Wallet.GetWalletIdFromOrgId(req.OrgId)
	if err != nil {
		log.WithError(err).Error("grpc_api/GetGatewayList")
		return &api.GetGatewayListResponse{}, err
	}

	totalGw, err := db.Gateway.GetGatewayRecCnt(walletId)
	if err != nil {
		log.WithError(err).Error("grpc_api/GetGatewayList")
		return &api.GetGatewayListResponse{}, err
	}

	offset := req.Offset * req.Limit

	gwList, err := db.Gateway.GetGatewayListOfWallet(walletId, offset, req.Limit)
	if err != nil {
		log.WithError(err).Error("grpc_api/GetGatewayList")
		return &api.GetGatewayListResponse{}, err
	}

	resp := api.GetGatewayListResponse{
		Count:       totalGw,
	}
	for _, v := range gwList {
		gwProfile := api.GatewayProfile{}
		gwProfile.Id = v.Id
		gwProfile.Mac = v.Mac
		gwProfile.FkGwNs = v.FkGatewayNs
		gwProfile.FkWallet = v.FkWallet
		gwMode := api.GatewayMode(api.GatewayMode_value[string(v.Mode)])
		gwProfile.Mode = gwMode
		gwProfile.CreateAt = v.CreatedAt.String()
		gwProfile.LastSeenAt = v.LastSeenAt.String()
		gwProfile.OrgId = v.OrgId
		gwProfile.Description = v.Description
		gwProfile.Name = v.Name

		resp.GwProfile = append(resp.GwProfile, &gwProfile)
	}

	return &resp, nil
}

func (s *M2MServerAPI) GetGatewayProfile(ctx context.Context, req *api.GetGatewayProfileRequest) (*api.GetGatewayProfileResponse, error) {
	log.WithFields(log.Fields{
		"orgId":  req.OrgId,
		"gwId":   req.GwId,
		"offset": req.Offset,
		"limit":  req.Limit,
	}).Debug("grpc_api/GetGatewayProfile")

	gwProfile, err := db.Gateway.GetGatewayProfile(req.GwId)
	if err != nil {
		log.WithError(err).Error("grpc_api/GetGatewayProfile")
		return &api.GetGatewayProfileResponse{}, err
	}

	gwMode := api.GatewayMode(api.GatewayMode_value[string(gwProfile.Mode)])

	resp := api.GatewayProfile{
		Id:          gwProfile.Id,
		Mac:         gwProfile.Mac,
		FkGwNs:      gwProfile.FkGatewayNs,
		FkWallet:    gwProfile.FkWallet,
		Mode:        gwMode,
		CreateAt:    gwProfile.CreatedAt.String(),
		LastSeenAt:  gwProfile.LastSeenAt.String(),
		OrgId:       gwProfile.OrgId,
		Description: gwProfile.Description,
		Name:        gwProfile.Name,
	}

	return &api.GetGatewayProfileResponse{GwProfile: &resp}, nil
}

func (s *M2MServerAPI) GetGatewayHistory(ctx context.Context, req *api.GetGatewayHistoryRequest) (*api.GetGatewayHistoryResponse, error) {
	return &api.GetGatewayHistoryResponse{}, nil
}

func (s *M2MServerAPI) SetGatewayMode(ctx context.Context, req *api.SetGatewayModeRequest) (*api.SetGatewayModeResponse, error) {
	log.WithFields(log.Fields{
		"orgId": req.OrgId,
		"gwID":  req.GwId,
		"gwMod": req.GwMode,
	}).Debug("grpc_api/SetDeviceMode")

	switch req.GwMode.String() {
	case api.GatewayMode_name[int32(api.GatewayMode_GW_INACTIVE)]:
		if err := db.Gateway.SetGatewayMode(req.GwId, types.GW_INACTIVE); err != nil {
			log.WithError(err).Error("grpc_api/SetDeviceMode")
			return &api.SetGatewayModeResponse{Status: false}, err
		}
	case api.GatewayMode_name[int32(api.GatewayMode_GW_FREE_GATEWAYS_LIMITED)]:
		if err := db.Gateway.SetGatewayMode(req.GwId, types.GW_FREE_GATEWAYS_LIMITED); err != nil {
			log.WithError(err).Error("grpc_api/SetDeviceMode")
			return &api.SetGatewayModeResponse{Status: false}, err
		}
	case api.GatewayMode_name[int32(api.GatewayMode_GW_WHOLE_NETWORK)]:
		if err := db.Gateway.SetGatewayMode(req.GwId, types.GW_WHOLE_NETWORK); err != nil {
			log.WithError(err).Error("grpc_api/SetDeviceMode")
			return &api.SetGatewayModeResponse{Status: false}, err
		}
	case api.GatewayMode_name[int32(api.GatewayMode_GW_DELETED)]:
		if err := db.Gateway.SetGatewayMode(req.GwId, types.GW_DELETED); err != nil {
			log.WithError(err).Error("grpc_api/SetDeviceMode")
			return &api.SetGatewayModeResponse{Status: false}, err
		}
	}

	return &api.SetGatewayModeResponse{}, nil
}
