package appserver

import (
	"context"
	"strconv"

	log "github.com/sirupsen/logrus"
	api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/appserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (s *M2MServerAPI) GetSettings(ctx context.Context, in *api.GetSettingsRequest) (*api.GetSettingsResponse, error) {
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

func (s *M2MServerAPI) ModifySettings(ctx context.Context, in *api.ModifySettingsRequest) (*api.ModifySettingsResponse, error) {
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
