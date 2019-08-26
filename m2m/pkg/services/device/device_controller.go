package device

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func Setup() error {
	log.Info("Setup device service")
	return nil
}

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

func (s *DeviceServerAPI) SetDeviceMode (ctx context.Context, req *api.SetDeviceModeRequest) (*api.SetDeviceModeResponse, error) {
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

	}

	return nil, status.Errorf(codes.Unknown, "")
}