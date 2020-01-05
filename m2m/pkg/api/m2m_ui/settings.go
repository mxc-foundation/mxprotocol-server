package ui

import (
	"context"
	"strconv"

	log "github.com/sirupsen/logrus"
	api "github.com/mxc-foundation/mxprotocol-server/m2m/api/m2m_ui"
	"github.com/mxc-foundation/mxprotocol-server/m2m/db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/auth"
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

	configs, err := db.ConfigTable.GetConfigs([]string{
		"downlink_fee",
		"transaction_percentage_share",
		"low_balance_warning",
	})

	if err != nil {
		log.WithError(err).Error("grpc_api/GetSettings")
		return nil, status.Error(codes.Internal, "")
	}

	result := &api.GetSettingsResponse{}

	for _, c := range configs {
		value, err := strconv.Atoi(c.Value.(string))
		if err != nil {
			log.WithError(err).Error("grpc_api/GetSettings")
			return nil, status.Error(codes.Internal, "")
		}
		switch true {
		case c.Key == "downlink_fee":
			result.DownlinkFee = int64(value)
		case c.Key == "transaction_percentage_share":
			result.TransactionPercentageShare = int64(value)
		case c.Key == "low_balance_warning":
			result.LowBalanceWarning = int64(value)
		}
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

	data := map[string]interface{}{}

	if in.TransactionPercentageShare != nil {
		data["transaction_percentage_share"] = in.TransactionPercentageShare.Value
	}

	if in.LowBalanceWarning != nil {
		data["low_balance_warning"] = in.LowBalanceWarning.Value
	}

	if in.DownlinkFee != nil {
		data["downlink_fee"] = in.DownlinkFee.Value
	}

	err := db.ConfigTable.UpdateConfigs(data)
	if err != nil {
		log.WithError(err).Error("grpc_api/ModifySettings")
		return nil, status.Error(codes.Internal, "")
	}

	return &api.ModifySettingsResponse{
		Status: true,
	}, nil
}
