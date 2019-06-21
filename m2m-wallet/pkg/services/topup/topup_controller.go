package topup

import (
	"context"
	log "github.com/sirupsen/logrus"
	"mxprotocol-server/m2m-wallet/api"
	"time"
)

func Setup() error {
	//todo
	log.Info("setup top_up service")
	return nil
}

type TopUpServerAPI struct {
	//todo
}

func NewTopUpServerAPI() *TopUpServerAPI {
	return &TopUpServerAPI{}
}

func (s *TopUpServerAPI) GetTopUpHistory(context.Context, *api.GetTopUpHistoryRequest) (*api.GetTopUpHistoryResponse, error) {
	var count = int64(4)
	history_list := api.GetTopUpHistoryResponse{
		Count:count,
	}

	for i := 0; i < int(count); i++ {
		item := api.TopUpHistory{
			From:"a",
			To:"b",
			Amount:12.333,
			CreatedAt: time.Now().UTC().String(),
		}

		history_list.TopupHistory = append(history_list.TopupHistory, &item)
	}

	return &history_list, nil
}