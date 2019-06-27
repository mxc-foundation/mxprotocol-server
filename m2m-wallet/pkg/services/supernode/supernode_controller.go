package supernode

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
)

func Setup() error {
	//todo
	log.Info("setup supernode service")
	return nil
}

type SupernodeServerAPI struct {
	//todo
}

func NewSupernodeServerAPI() *SupernodeServerAPI {
	return &SupernodeServerAPI{}
}

func (s *SupernodeServerAPI) GetSuperNodeActiveMoneyAccount(context.Context, *api.GetSuperNodeActiveMoneyAccountRequest) (*api.GetSuperNodeActiveMoneyAccountResponse, error) {
	return &api.GetSuperNodeActiveMoneyAccountResponse{SupernodeActiveAccount: "supernode_account", Error: ""}, nil
}
