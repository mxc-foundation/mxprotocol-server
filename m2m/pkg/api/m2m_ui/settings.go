package ui

import (
	"context"

	log "github.com/sirupsen/logrus"
	api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/m2m_ui"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	pdb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SettingsServerAPI struct {
	serviceName string
}

func NewSettingsServerAPI() *SettingsServerAPI {
	return &SettingsServerAPI{serviceName: "settings"}
}

func (s *SettingsServerAPI) GetSettings(ctx context.Context, in *api.GetSettingsRequest) (*api.GetSettingsResponse, error) {
	log.Info("GetSettings")
	_, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, 0)

	switch res.Type {
	case auth.AuthFailed, auth.JsonParseError, auth.OrganizationIdMisMatch, auth.OrganizationIdRearranged:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)
	case auth.OK:
	default:
		return nil, status.Error(codes.Unknown, "")
	}

	config, err := db.ConfigTable.Get()

	if err != nil {
		log.WithError(err).Error("grpc_api/GetSettings")
		return nil, status.Error(codes.Internal, "")
	}

	log.Info(config)

	result := &api.GetSettingsResponse{
		DownlinkFee:                int64(*config.DownlinkFee),
		LowBalanceWarning:          int64(*config.LowBalanceWarning),
		TransactionPercentageShare: int64(*config.TransactionPercentageShare),
	}

	return result, nil
}

func (s *SettingsServerAPI) ModifySettings(ctx context.Context, in *api.ModifySettingsRequest) (*api.ModifySettingsResponse, error) {
	_, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, 0)

	switch res.Type {
	case auth.AuthFailed, auth.JsonParseError, auth.OrganizationIdMisMatch, auth.OrganizationIdRearranged:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)
	case auth.OK:
	default:
		return nil, status.Error(codes.Unknown, "")
	}

	if in.TransactionPercentageShare == nil && in.LowBalanceWarning == nil && in.DownlinkFee == nil {
		return nil, status.Error(codes.InvalidArgument, "")
	}

	data := &pdb.Config{}

	if in.TransactionPercentageShare != nil {
		data.TransactionPercentageShare = intPnt(in.TransactionPercentageShare.Value)
	}

	if in.LowBalanceWarning != nil {
		data.LowBalanceWarning = intPnt(in.LowBalanceWarning.Value)
	}

	if in.DownlinkFee != nil {
		data.DownlinkFee = intPnt(in.DownlinkFee.Value)
	}

	err := db.ConfigTable.Update(data)
	if err != nil {
		log.WithError(err).Error("grpc_api/ModifySettings")
		return nil, status.Error(codes.Internal, "")
	}

	return &api.ModifySettingsResponse{
		Status: true,
	}, nil
}

func intPnt(data int64) *int {
	value := int(data)
	return &value
}
