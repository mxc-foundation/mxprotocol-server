package withdraw

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"mxprotocol-server/m2m-wallet/api"
)

func Setup() error {
	//check database and update ctxWithdraw.withdrawFee
	log.Info(fmt.Sprintf("setup withdraw service( current withdraw fee = %f", ctxWithdraw.withdrawFee))
	return nil
}

type WithdrawServerAPI struct {
}

func NewWithdrawServerAPI() *WithdrawServerAPI {
	return &WithdrawServerAPI{}
}

func (s *WithdrawServerAPI) GetWithdrawFee(ctx context.Context, req *api.GetWithdrawFeeRequest) (*api.GetWithdrawFeeResponse, error) {
	ctxWithdraw.withdrawFee += 2.0
	return &api.GetWithdrawFeeResponse{WithdrawFee: ctxWithdraw.withdrawFee, Error: ""}, nil
}
