package ext_account

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/auth"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/services/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Setup() error {
	log.Info("Setup ext_account service")
	return nil
}

func updateActiveExtAccount(orgId int64, newAccount string, currencyAbbr string) error {
	walletId, err := wallet.GetWalletId(orgId)
	if err != nil {
		return err
	}

	_, err = db.DBInsertExtAccount(walletId, newAccount, currencyAbbr)
	if err != nil {
		return err
	}

	return nil
}

type ExtAccountServerAPI struct {
	serviceName string
}

func NewMoneyServerAPI() *ExtAccountServerAPI {
	return &ExtAccountServerAPI{serviceName: "ext_account"}
}

func (s *ExtAccountServerAPI) ModifyMoneyAccount(ctx context.Context, req *api.ModifyMoneyAccountRequest) (*api.ModifyMoneyAccountResponse, error) {
	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	log.WithFields(log.Fields{
		"orgID":      req.OrgId,
		"moneyAbbr":  api.Money_name[int32(req.MoneyAbbr)],
		"newAccount": req.CurrentAccount,
	}).Debug("grpc_api/ModifyMoneyAccount")

	err = updateActiveExtAccount(req.OrgId, req.CurrentAccount, api.Money_name[int32(req.MoneyAbbr)])
	if err != nil {
		return &api.ModifyMoneyAccountResponse{Status: false, UserProfile: &userProfile}, err
	}

	return &api.ModifyMoneyAccountResponse{Status: true, UserProfile: &userProfile}, nil
}

func (s *ExtAccountServerAPI) GetChangeMoneyAccountHistory(ctx context.Context, req *api.GetMoneyAccountChangeHistoryRequest) (*api.GetMoneyAccountChangeHistoryResponse, error) {
	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	var count = int64(3)

	history_list := []*api.MoneyAccountChangeHistory{}
	for i := 0; i < int(count); i++ {
		item := api.MoneyAccountChangeHistory{
			From:      "alice",
			To:        "bob",
			CreatedAt: time.Now().UTC().String(),
		}
		history_list = append(history_list, &item)
	}

	return &api.GetMoneyAccountChangeHistoryResponse{Count: count, ChangeHistory: history_list, UserProfile: &userProfile}, nil
}

func (s *ExtAccountServerAPI) GetActiveMoneyAccount(ctx context.Context, req *api.GetActiveMoneyAccountRequest) (*api.GetActiveMoneyAccountResponse, error) {
	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	log.WithFields(log.Fields{
		"orgId":     req.OrgId,
		"moneyAbbr": api.Money_name[int32(req.MoneyAbbr)],
	}).Debug("grpc_api/GetActiveMoneyAccount")

	walletId, err := db.DbGetWalletIdFromOrgId(req.OrgId)
	if err != nil {
		return &api.GetActiveMoneyAccountResponse{ActiveAccount: "", UserProfile: &userProfile}, err
	}

	accountAddr, err := db.DbGetUserExtAccountAdr(walletId, api.Money_name[int32(req.MoneyAbbr)])
	if err != nil {
		return &api.GetActiveMoneyAccountResponse{ActiveAccount: "", UserProfile: &userProfile}, err
	}

	return &api.GetActiveMoneyAccountResponse{ActiveAccount: accountAddr, UserProfile: &userProfile}, nil
}
