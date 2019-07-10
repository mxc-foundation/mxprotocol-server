package money

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/auth"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/services/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func Setup() error {
	log.Info("setup money service")
	return nil
}

func updateActiveMoneyAccount(orgId int64, newAccount string, currencyAbbr string) error {
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

type MoneyServerAPI struct {
	serviceName string
}

func NewMoneyServerAPI() *MoneyServerAPI {
	return &MoneyServerAPI{serviceName: "money"}
}

func (s *MoneyServerAPI) ModifyMoneyAccount(ctx context.Context, req *api.ModifyMoneyAccountRequest) (*api.ModifyMoneyAccountResponse, error) {
	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return &api.ModifyMoneyAccountResponse{Status: false, UserProfile: &userProfile},
			status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	err = updateActiveMoneyAccount(req.OrgId, req.CurrentAccount, api.Money_name[int32(req.MoneyAbbr)])
	if err != nil {
		return &api.ModifyMoneyAccountResponse{Status: false, UserProfile: &userProfile}, err
	}

	return &api.ModifyMoneyAccountResponse{Status: true, UserProfile: &userProfile}, nil
}

func (s *MoneyServerAPI) GetChangeMoneyAccountHistory(ctx context.Context, req *api.GetMoneyAccountChangeHistoryRequest) (*api.GetMoneyAccountChangeHistoryResponse, error) {
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

func (s *MoneyServerAPI) GetActiveMoneyAccount(ctx context.Context, req *api.GetActiveMoneyAccountRequest) (*api.GetActiveMoneyAccountResponse, error) {
	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	return &api.GetActiveMoneyAccountResponse{ActiveAccount: "", UserProfile: &userProfile}, nil
}
