package wallet

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/auth"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func Setup(conf config.MxpConfig) error {
	//todo
	log.Info("setup wallet service")

	return nil
}

func userHasWallet(orgId int64) bool {
	// check from table wallet with: username, orgId, wallet_type=normal_user
	return false
}

func createWallet(orgId int64) (int64, error) {
	// create wallet with: username, orgId, wallet type = normal user
	return 0, nil
}

func GetWalletId(orgId int64) (walletId int64, err error) {
	if false == userHasWallet(orgId) {
		if walletId, err = createWallet(orgId); err != nil {
			return 0, err
		}
	}

	return walletId, nil
}

func GetBalance(orgId int64) (float64, error) {
	walletId, err := GetWalletId(orgId)
	if err != nil {
		return 0, err
	}

	res, err := db.DbWalletGetBalanceByWalletId(walletId)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func UpdateBalance(orgId int64, oper PaymentCategory, deviceType DeviceType, amount float64) error {
	walletId, err := GetWalletId(orgId)
	if err != nil {
		return err
	}

	balance, err := db.DbWalletGetBalanceByWalletId(walletId)
	if err != nil {
		return err
	}

	for _, v := range operMap {
		if v.pc == oper && v.dt == deviceType {
			balance = v.operation(balance, amount)
		}
	}

	err = db.DbWalletUpdateBalanceByWalletId(walletId, balance)
	if err != nil {
		return err
	}

	return nil
}

// grpc APIs

type WalletServerAPI struct {
	serviceName string
}

func NewWalletServerAPI() *WalletServerAPI {
	return &WalletServerAPI{serviceName: "wallet"}
}

func (s *WalletServerAPI) GetWalletBalance(ctx context.Context, req *api.GetWalletBalanceRequest) (*api.GetWalletBalanceResponse, error) {
	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return &api.GetWalletBalanceResponse{}, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	walletId, err := GetWalletId(req.OrgId)
	if err != nil {
		return &api.GetWalletBalanceResponse{}, err
	}

	balance, err := db.DbWalletGetBalanceByWalletId(walletId)
	if err != nil {
		return &api.GetWalletBalanceResponse{}, err
	}

	return &api.GetWalletBalanceResponse{Balance: balance, Error: "", UserProfile: &userProfile}, nil
}

func (s *WalletServerAPI) GetVmxcTxHistory(ctx context.Context, req *api.GetVmxcTxHistoryRequest) (*api.GetVmxcTxHistoryResponse, error) {
	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	//Todo: userInfo should be the information of users eg.id,name,org,etc. Use it to get data from DB.
	fmt.Println("username = ", userProfile.User.Username)

	var count = int64(6)
	history_list := []*api.VmxcTxHistory{}
	for i := 0; i < int(count); i++ {
		item := api.VmxcTxHistory{
			From:      "a",
			To:        "b",
			TxType:    "subscription",
			Amount:    12.333,
			CreatedAt: time.Now().UTC().String(),
		}

		history_list = append(history_list, &item)
	}

	return &api.GetVmxcTxHistoryResponse{Error: "", Count: count, TxHistory: history_list, UserProfile: &userProfile}, nil
}
