package m2m

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

type DeviceServerAPI struct {
	serviceName string
}

func NewDeviceServerAPI() *DeviceServerAPI {
	return &DeviceServerAPI{serviceName: "device"}
}

func (s *DeviceServerAPI) GetDeviceList(ctx context.Context, req *api.GetDeviceListRequest) (*api.GetDeviceListResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)
	case auth.OrganizationIdRearranged:
		return &api.GetDeviceListResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")
	case auth.OK:
		log.WithFields(log.Fields{
			"orgId":  req.OrgId,
			"offset": req.Offset,
			"limit":  req.Limit,
		}).Debug("grpc_api/GetDeviceList")

		dvList, err := db.Device.GetDeviceListOfWallet(req.OrgId, req.Offset, req.Limit)
		if err != nil {
			log.WithError(err).Error("grpc_api/GetDeviceList")
			return &api.GetDeviceListResponse{UserProfile: &userProfile}, err
		}

		resp := api.GetDeviceListResponse{UserProfile: &userProfile}
		for _, v := range dvList {
			dvProfile := api.DeviceProfile{}
			dvProfile.Id = v.Id
			dvProfile.DevEui = v.DevEui
			dvProfile.FkWallet = v.FkWallet
			dvProfile.Mode = string(v.Mode)
			dvProfile.CreatedAt = v.CreatedAt.String()
			dvProfile.LastSeenAt = v.LastSeenAt.String()
			dvProfile.ApplicationId = v.ApplicationId
			dvProfile.Name = v.Name

			resp.DevProfile = append(resp.DevProfile, &dvProfile)
		}

		return &resp, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *DeviceServerAPI) GetDeviceProfile(ctx context.Context, req *api.GetDeviceProfileRequest) (*api.GetDeviceProfileResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)
	case auth.OrganizationIdRearranged:
		return &api.GetDeviceProfileResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")
	case auth.OK:
		log.WithFields(log.Fields{
			"orgId": req.OrgId,
			"devId": req.DevId,
		}).Debug("grpc_api/GetDeviceProfile")

		devProfile, err := db.Device.GetDeviceProfile(req.DevId)
		if err != nil {
			log.WithError(err).Error("grpc_api/GetDeviceProfile")
			return &api.GetDeviceProfileResponse{UserProfile: &userProfile}, err
		}

		resp := api.DeviceProfile{
			Id:            devProfile.Id,
			DevEui:        devProfile.DevEui,
			FkWallet:      devProfile.FkWallet,
			Mode:          string(devProfile.Mode),
			CreatedAt:     devProfile.CreatedAt.String(),
			LastSeenAt:    devProfile.LastSeenAt.String(),
			ApplicationId: devProfile.ApplicationId,
			Name:          devProfile.Name,
		}

		return &api.GetDeviceProfileResponse{UserProfile: &userProfile, DevProfile: &resp}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *DeviceServerAPI) GetDeviceHistory(ctx context.Context, req *api.GetDeviceHistoryRequest) (*api.GetDeviceHistoryResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)
	case auth.OrganizationIdRearranged:
		return &api.GetDeviceHistoryResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")
	case auth.OK:

	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *DeviceServerAPI) SetDeviceMode(ctx context.Context, req *api.SetDeviceModeRequest) (*api.SetDeviceModeResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)
	case auth.OrganizationIdRearranged:
		return &api.SetDeviceModeResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")
	case auth.OK:
		log.WithFields(log.Fields{
			"orgId":  req.OrgId,
			"devID":  req.DevId,
			"devMod": req.DevMode,
		}).Debug("grpc_api/SetDeviceMode")

		switch req.DevMode.String() {
		case "DV_INACTIVE":
			if err := db.Device.SetDeviceMode(req.DevId, types.DV_INACTIVE); err != nil {
				log.WithError(err).Error("grpc_api/SetDeviceMode")
				return &api.SetDeviceModeResponse{Status: false, UserProfile: &userProfile}, err
			}
		case "DV_FREE_GATEWAYS_LIMITED":
			if err := db.Device.SetDeviceMode(req.DevId, types.DV_FREE_GATEWAYS_LIMITED); err != nil {
				log.WithError(err).Error("grpc_api/SetDeviceMode")
				return &api.SetDeviceModeResponse{Status: false, UserProfile: &userProfile}, err
			}
		case "DV_WHOLE_NETWORK":
			if err := db.Device.SetDeviceMode(req.DevId, types.DV_WHOLE_NETWORK); err != nil {
				log.WithError(err).Error("grpc_api/SetDeviceMode")
				return &api.SetDeviceModeResponse{Status: false, UserProfile: &userProfile}, err
			}
		case "DV_DELETED":
			if err := db.Device.SetDeviceMode(req.DevId, types.DV_DELETED); err != nil {
				log.WithError(err).Error("grpc_api/SetDeviceMode")
				return &api.SetDeviceModeResponse{Status: false, UserProfile: &userProfile}, err
			}
		}

		/*if err := db.DbSetDeviceMode(req.DevId, req.DevMode); err != nil {
			log.WithError(err).Error("grpc_api/SetDeviceMode")
			return &api.SetDeviceModeResponse{Status: false, UserProfile: &userProfile}, err
		}*/

		return &api.SetDeviceModeResponse{UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}
