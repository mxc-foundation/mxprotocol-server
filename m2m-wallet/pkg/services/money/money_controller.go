package money

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func Setup() error {
	//todo
	log.Info("setup money service")
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
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	//Todo: userInfo should be the information of users eg.id,name,org,etc. Use it to get data from DB.
	fmt.Println("username = ", userProfile.User.Username)

	return &api.ModifyMoneyAccountResponse{Error: "test", Status: true, UserProfile: &userProfile}, nil
}

func (s *MoneyServerAPI) GetChangeMoneyAccountHistory(ctx context.Context, req *api.GetMoneyAccountChangeHistoryRequest) (*api.GetMoneyAccountChangeHistoryResponse, error) {
	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	//Todo: userInfo should be the information of users eg.id,name,org,etc. Use it to get data from DB.
	fmt.Println("username = ", userProfile.User.Username)
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

	return &api.GetMoneyAccountChangeHistoryResponse{Error:"", Count: count, ChangeHistory: history_list, UserProfile: &userProfile}, nil
}

func (s *MoneyServerAPI) GetActiveMoneyAccount(ctx context.Context, req *api.GetActiveMoneyAccountRequest) (*api.GetActiveMoneyAccountResponse, error) {
	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	//Todo: userInfo should be the information of users eg.id,name,org,etc. Use it to get data from DB.
	fmt.Println("username = ", userProfile.User.Username)
	return &api.GetActiveMoneyAccountResponse{Error: "", ActiveAccount: "", UserProfile: &userProfile}, nil
}
