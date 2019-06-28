package wallet

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	"time"
)

func Setup() error {
	//todo
	log.Info("setup wallet service")
	return nil
}

type WalletServerAPI struct {
	//todo
}

func NewWalletServerAPI() *WalletServerAPI {
	return &WalletServerAPI{}
}

func (s *WalletServerAPI) GetWalletBalance(context.Context, *api.GetWalletBalanceRequest) (*api.GetWalletBalanceResponse, error) {
	return &api.GetWalletBalanceResponse{Balance: 12345.1212, Error: ""}, nil
}

func (s *WalletServerAPI) GetVmxcTxHistory(context.Context, *api.GetVmxcTxHistoryRequest) (*api.GetVmxcTxHistoryResponse, error) {
	var count = int64(6)
	history_list := api.GetVmxcTxHistoryResponse{
		Count: count,
	}

	for i := 0; i < int(count); i++ {
		item := api.VmxcTxHistory{
			From:      "a",
			To:        "b",
			TxType:    "subscription",
			Amount:    12.333,
			CreatedAt: time.Now().UTC().String(),
		}

		history_list.TxHistory = append(history_list.TxHistory, &item)
	}

	return &history_list, nil
}
