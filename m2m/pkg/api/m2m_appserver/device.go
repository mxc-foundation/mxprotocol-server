package appserver

import (
	"context"

	log "github.com/sirupsen/logrus"
	api "github.com/mxc-foundation/mxprotocol-server/m2m/api/m2m_ui"
	"github.com/mxc-foundation/mxprotocol-server/m2m/db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/types"
)

func (s *M2MServerAPI) GetDeviceList(ctx context.Context, req *api.GetDeviceListRequest) (*api.GetDeviceListResponse, error) {
	log.WithFields(log.Fields{
		"orgId":  req.OrgId,
		"offset": req.Offset,
		"limit":  req.Limit,
	}).Debug("grpc_api/GetDeviceList")

	walletId, err := db.Wallet.GetWalletIdFromOrgId(req.OrgId)
	if err != nil {
		log.WithError(err).Error("grpc_api/GetDeviceList")
		return &api.GetDeviceListResponse{}, err
	}

	totalDev, err := db.Device.GetDeviceRecCnt(walletId)
	if err != nil {
		log.WithError(err).Error("grpc_api/GetDeviceList")
		return &api.GetDeviceListResponse{}, err
	}

	offset := req.Offset * req.Limit

	dvList, err := db.Device.GetDeviceListOfWallet(walletId, offset, req.Limit)
	if err != nil {
		log.WithError(err).Error("grpc_api/GetDeviceList")
		return &api.GetDeviceListResponse{}, err
	}

	resp := api.GetDeviceListResponse{
		Count: totalDev,
	}
	for _, v := range dvList {
		dvProfile := api.DeviceProfile{}
		dvProfile.Id = v.Id
		dvProfile.DevEui = v.DevEui
		dvProfile.FkWallet = v.FkWallet
		dvMode := api.DeviceMode(api.DeviceMode_value[string(v.Mode)])
		dvProfile.Mode = dvMode
		dvProfile.CreatedAt = v.CreatedAt.String()
		dvProfile.LastSeenAt = v.LastSeenAt.String()
		dvProfile.ApplicationId = v.ApplicationId
		dvProfile.Name = v.Name

		resp.DevProfile = append(resp.DevProfile, &dvProfile)
	}

	return &resp, nil
}

func (s *M2MServerAPI) GetDeviceProfile(ctx context.Context, req *api.GetDeviceProfileRequest) (*api.GetDeviceProfileResponse, error) {
	log.WithFields(log.Fields{
		"orgId": req.OrgId,
		"devId": req.DevId,
	}).Debug("grpc_api/GetDeviceProfile")

	devProfile, err := db.Device.GetDeviceProfile(req.DevId)
	if err != nil {
		log.WithError(err).Error("grpc_api/GetDeviceProfile")
		return &api.GetDeviceProfileResponse{}, err
	}

	dvMode := api.DeviceMode(api.DeviceMode_value[string(devProfile.Mode)])

	resp := api.DeviceProfile{
		Id:            devProfile.Id,
		DevEui:        devProfile.DevEui,
		FkWallet:      devProfile.FkWallet,
		Mode:          dvMode,
		CreatedAt:     devProfile.CreatedAt.String(),
		LastSeenAt:    devProfile.LastSeenAt.String(),
		ApplicationId: devProfile.ApplicationId,
		Name:          devProfile.Name,
	}

	return &api.GetDeviceProfileResponse{DevProfile: &resp}, nil
}

func (s *M2MServerAPI) GetDeviceHistory(ctx context.Context, req *api.GetDeviceHistoryRequest) (*api.GetDeviceHistoryResponse, error) {
	return &api.GetDeviceHistoryResponse{}, nil
}

func (s *M2MServerAPI) SetDeviceMode(ctx context.Context, req *api.SetDeviceModeRequest) (*api.SetDeviceModeResponse, error) {
	log.WithFields(log.Fields{
		"orgId":  req.OrgId,
		"devID":  req.DevId,
		"devMod": req.DevMode,
	}).Debug("grpc_api/SetDeviceMode")

	switch req.DevMode.String() {
	case api.DeviceMode_name[int32(api.DeviceMode_DV_INACTIVE)]:
		if err := db.Device.SetDeviceMode(req.DevId, types.DV_INACTIVE); err != nil {
			log.WithError(err).Error("grpc_api/SetDeviceMode")
			return &api.SetDeviceModeResponse{Status: false}, err
		}
	case api.DeviceMode_name[int32(api.DeviceMode_DV_FREE_GATEWAYS_LIMITED)]:
		if err := db.Device.SetDeviceMode(req.DevId, types.DV_FREE_GATEWAYS_LIMITED); err != nil {
			log.WithError(err).Error("grpc_api/SetDeviceMode")
			return &api.SetDeviceModeResponse{Status: false}, err
		}
	case api.DeviceMode_name[int32(api.DeviceMode_DV_WHOLE_NETWORK)]:
		if err := db.Device.SetDeviceMode(req.DevId, types.DV_WHOLE_NETWORK); err != nil {
			log.WithError(err).Error("grpc_api/SetDeviceMode")
			return &api.SetDeviceModeResponse{Status: false}, err
		}
	case api.DeviceMode_name[int32(api.DeviceMode_DV_DELETED)]:
		if err := db.Device.SetDeviceMode(req.DevId, types.DV_DELETED); err != nil {
			log.WithError(err).Error("grpc_api/SetDeviceMode")
			return &api.SetDeviceModeResponse{Status: false}, err
		}
	}

	/*if err := db.DbSetDeviceMode(req.DevId, req.DevMode); err != nil {
		log.WithError(err).Error("grpc_api/SetDeviceMode")
		return &api.SetDeviceModeResponse{Status: false, UserProfile: &userProfile}, err
	}*/

	return &api.SetDeviceModeResponse{}, nil
}
