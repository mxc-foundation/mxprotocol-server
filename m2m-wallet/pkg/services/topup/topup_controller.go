package topup

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func Setup() error {
	log.Info("Setup top_up service")
	return nil
}

type TopUpServerAPI struct {
	serviceName string
}

func NewTopUpServerAPI() *TopUpServerAPI {
	return &TopUpServerAPI{serviceName: "top up"}
}

func (s *TopUpServerAPI) GetTopUpHistory(ctx context.Context, req *api.GetTopUpHistoryRequest) (*api.GetTopUpHistoryResponse, error) {
	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	var count = int64(4)
	history_list := []*api.TopUpHistory{}

	for i := 0; i < int(count); i++ {
		item := api.TopUpHistory{
			From:      "a",
			To:        "b",
			Amount:    12.333,
			CreatedAt: time.Now().UTC().String(),
		}

		history_list = append(history_list, &item)
	}

	return &api.GetTopUpHistoryResponse{Count: count, TopupHistory: history_list, UserProfile: &userProfile}, nil
}
