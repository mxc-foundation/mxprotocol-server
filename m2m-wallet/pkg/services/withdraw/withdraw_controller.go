package withdraw

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	"time"
)

func Setup() error {
	//todo
	//check database and update ctxWithdraw.withdrawFee
	log.Info(fmt.Sprintf("setup withdraw service( current withdraw fee = %f", ctxWithdraw.withdrawFee))
	return nil
}

type WithdrawServerAPI struct {
	//todo
}

func NewWithdrawServerAPI() *WithdrawServerAPI {
	return &WithdrawServerAPI{}
}

func (s *WithdrawServerAPI) GetWithdrawFee(ctx context.Context, req *api.GetWithdrawFeeRequest) (*api.GetWithdrawFeeResponse, error) {
	//todo
	ctxWithdraw.withdrawFee += 2.0
	return &api.GetWithdrawFeeResponse{WithdrawFee: ctxWithdraw.withdrawFee, Error: ""}, nil
}

func (s *WithdrawServerAPI) GetWithdrawHistory(context.Context, *api.GetWithdrawHistoryRequest) (*api.GetWithdrawHistoryResponse, error) {
	var count = int64(6)
	history_list := api.GetWithdrawHistoryResponse{
		Count: count,
	}

	for i := 0; i < int(count); i++ {
		item := api.WithdrawHistory{
			From:      "a",
			To:        "b",
			MoneyType: "Ether",
			Amount:    12.333,
			CreatedAt: time.Now().UTC().String(),
		}

		history_list.WithdrawHistory = append(history_list.WithdrawHistory, &item)
	}

	return &history_list, nil
}

func (s *WithdrawServerAPI) WithdrawReq(context.Context, *api.WithdrawReqRequest) (*api.WithdrawReqResponse, error) {
	//todo
	return &api.WithdrawReqResponse{Status: true, Error: ""}, nil
}
