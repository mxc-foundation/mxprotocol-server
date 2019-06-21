package money

import (
	"context"
	log "github.com/sirupsen/logrus"
	"mxprotocol-server/m2m-wallet/api"
	"time"
)

func Setup() error {
	//todo
	log.Info("setup money service")
	return nil
}

type MoneyServerAPI struct {
	//todo
}

func NewMoneyServerAPI() *MoneyServerAPI {
	return &MoneyServerAPI{}
}

func (s *MoneyServerAPI) ModifyMoneyAccount(context.Context, *api.ModifyMoneyAccountRequest) (*api.ModifyMoneyAccountResponse, error) {
	return &api.ModifyMoneyAccountResponse{Error: "test", Status: true}, nil
}

func (s *MoneyServerAPI) GetChangeMoneyAccountHistory(context.Context, *api.GetMoneyAccountChangeHistoryRequest) (*api.GetMoneyAccountChangeHistoryResponse, error) {
	var count = int64(3)
	history_list := api.GetMoneyAccountChangeHistoryResponse{
		Count: count,
	}

	for i := 0; i < int(count); i++ {
		item := api.MoneyAccountChangeHistory{
			From:"alice",
			To:"bob",
			CreatedAt: time.Now().UTC().String(),
		}
		history_list.ChangeHistory = append(history_list.ChangeHistory, &item)
	}

	return &history_list, nil
}

func (s *MoneyServerAPI) GetActiveMoneyAccount(context.Context,	*api.GetActiveMoneyAccountRequest) (*api.GetActiveMoneyAccountResponse, error) {
	return &api.GetActiveMoneyAccountResponse{Error:"", ActiveAccount:"",}, nil
}
