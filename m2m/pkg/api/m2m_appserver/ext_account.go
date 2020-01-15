package appserver

import (
	"context"
	log "github.com/sirupsen/logrus"
	api "github.com/mxc-foundation/mxprotocol-server/m2m/api/m2m_ui"
	"github.com/mxc-foundation/mxprotocol-server/m2m/db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/services/ext_account"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/services/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func (s *M2MServerAPI) ModifyMoneyAccount(ctx context.Context, req *api.ModifyMoneyAccountRequest) (*api.ModifyMoneyAccountResponse, error) {
	log.WithFields(log.Fields{
		"orgId":       req.OrgId,
		"moneyAbbr":   api.Money_name[int32(req.MoneyAbbr)],
		"accountAddr": strings.ToLower(req.CurrentAccount),
	}).Debug("grpc_api/ModifyMoneyAccount")

	if 0 == req.OrgId {
		return &api.ModifyMoneyAccountResponse{Status: false}, nil
	}

	err := ext_account.UpdateActiveExtAccount(req.OrgId, req.CurrentAccount, api.Money_name[int32(req.MoneyAbbr)])
	if err != nil {
		log.WithError(err).Error("grpc_api/ModifyMoneyAccount")
		return &api.ModifyMoneyAccountResponse{Status: false},
			status.Errorf(codes.InvalidArgument, "Duplicate or invalid format.")
	}
	return &api.ModifyMoneyAccountResponse{Status: true}, nil
}

func (s *M2MServerAPI) GetChangeMoneyAccountHistory(ctx context.Context, req *api.GetMoneyAccountChangeHistoryRequest) (*api.GetMoneyAccountChangeHistoryResponse, error) {
	log.WithFields(log.Fields{
		"orgId":     req.OrgId,
		"offset":    req.Offset,
		"limit":     req.Limit,
		"moneyAbbr": api.Money_name[int32(req.MoneyAbbr)],
	}).Debug("grpc_api/GetChangeMoneyAccountHistory")

	walletId, err := wallet.GetWalletId(req.OrgId)
	if err != nil {
		log.WithError(err).Error("grpc_api/GetChangeMoneyAccountHistory")
		return &api.GetMoneyAccountChangeHistoryResponse{}, nil
	}

	response := api.GetMoneyAccountChangeHistoryResponse{}
	ptr, err := db.ExtAccount.GetExtAcntHist(walletId, req.Offset*req.Limit, req.Limit)
	if err != nil {
		log.WithError(err).Error("grpc_api/GetChangeMoneyAccountHistory")
		return &api.GetMoneyAccountChangeHistoryResponse{}, nil
	}

	for _, v := range ptr {
		if v.ExtCurrencyAbv != api.Money_name[int32(req.MoneyAbbr)] {
			continue
		}
		history := api.MoneyAccountChangeHistory{}
		history.Addr = v.AccountAdr
		history.CreatedAt = v.InsertTime.String()
		history.Status = v.Status
		response.ChangeHistory = append(response.ChangeHistory, &history)
	}
	response.Count, err = db.ExtAccount.GetExtAcntHistRecCnt(walletId)

	return &response, nil
}

func (s *M2MServerAPI) GetActiveMoneyAccount(ctx context.Context, req *api.GetActiveMoneyAccountRequest) (*api.GetActiveMoneyAccountResponse, error) {
	log.WithFields(log.Fields{
		"orgId":     req.OrgId,
		"moneyAbbr": api.Money_name[int32(req.MoneyAbbr)],
	}).Debug("grpc_api/GetActiveMoneyAccount")

	accountAddr, err := ext_account.GetActiveExtAccount(req.OrgId, api.Money_name[int32(req.MoneyAbbr)])
	if err != nil {
		log.WithError(err).Error("grpc_api/GetActiveMoneyAccount")
		return &api.GetActiveMoneyAccountResponse{ActiveAccount: ""}, nil
	}

	return &api.GetActiveMoneyAccountResponse{ActiveAccount: accountAddr}, nil
}
