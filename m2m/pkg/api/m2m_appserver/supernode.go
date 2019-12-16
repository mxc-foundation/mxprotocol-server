package appserver

import (
	"context"
	log "github.com/sirupsen/logrus"
	api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/m2m_ui"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services/ext_account"
	"strings"
)

func (s *M2MServerAPI) AddSuperNodeMoneyAccount(ctx context.Context, in *api.AddSuperNodeMoneyAccountRequest) (*api.AddSuperNodeMoneyAccountResponse, error) {
	log.WithFields(log.Fields{
		"moneyAbbr":   api.Money_name[int32(in.MoneyAbbr)],
		"accountAddr": strings.ToLower(in.AccountAddr),
	}).Debug("grpc_api/AddSuperNodeMoneyAccount")

	err := ext_account.UpdateActiveExtAccount(0, in.AccountAddr, api.Money_name[int32(in.MoneyAbbr)])
	if err != nil {
		log.WithError(err).Error("grpc_api/AddSuperNodeMoneyAccount")
		return &api.AddSuperNodeMoneyAccountResponse{Status: false}, nil
	}

	return &api.AddSuperNodeMoneyAccountResponse{Status: true}, nil

}

func (s *M2MServerAPI) GetSuperNodeActiveMoneyAccount(ctx context.Context, req *api.GetSuperNodeActiveMoneyAccountRequest) (*api.GetSuperNodeActiveMoneyAccountResponse, error) {
	log.WithFields(log.Fields{
		"moneyAbbr": api.Money_name[int32(req.MoneyAbbr)],
	}).Debug("grpc_api/GetSuperNodeActiveMoneyAccount")

	accountAddr, err := ext_account.GetActiveExtAccount(0, api.Money_name[int32(req.MoneyAbbr)])
	if err != nil {
		log.WithError(err).Error("grpc_api/GetSuperNodeActiveMoneyAccount")
		return &api.GetSuperNodeActiveMoneyAccountResponse{SupernodeActiveAccount: ""}, nil
	}

	return &api.GetSuperNodeActiveMoneyAccountResponse{SupernodeActiveAccount: accountAddr}, nil
}
