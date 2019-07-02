package wallet

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

func Setup() error {
	//todo
	log.Info("setup wallet service")
	return nil
}

type WalletServerAPI struct {
	//todo
}

func NewWalletServerAPI() *WalletServerAPI {
	return &WalletServerAPI{}
}

func (s *WalletServerAPI) GetWalletBalance(ctx context.Context, req *api.GetWalletBalanceRequest) (*api.GetWalletBalanceResponse, error) {
	//Todo: change url if using diff API, if only one API, put it into conf.
	//url := "http://172.19.0.5:8080/api/jwtvalidate/"
	url := "http://appserver:8080/api/internal/profile"

	info, err := auth.TokenMiddleware(ctx, url)
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	errInfo := auth.Err{}
	err = json.Unmarshal(*info, &errInfo)
	if err != nil{
		fmt.Println(err)
	}

	if errInfo.Error != ""{
		return nil, grpc.Errorf(codes.Unauthenticated, "authentication failed: %s", errInfo.Error)
	}

	userInfo := auth.ProfileResponse{}
	err = json.Unmarshal(*info, &userInfo)
	if err != nil {
		fmt.Println(err)
	}

	//Todo: userInfo should be the information of users eg.id,name,org,etc. Use it to get data from DB.
	fmt.Println("userName: ", userInfo.User.Username)

	return &api.GetWalletBalanceResponse{Balance: 54321.1212, Error: ""}, nil
}

func (s *WalletServerAPI) GetVmxcTxHistory(context.Context, *api.GetVmxcTxHistoryRequest) (*api.GetVmxcTxHistoryResponse, error) {
	var count = int64(6)
	history_list := api.GetVmxcTxHistoryResponse{
		Count: count,
	}

	for i := 0; i < int(count); i++ {
		item := api.VmxcTxHistory{
			From:      "a",
			To:        "b",
			TxType:    "subscription",
			Amount:    12.333,
			CreatedAt: time.Now().UTC().String(),
		}

		history_list.TxHistory = append(history_list.TxHistory, &item)
	}

	return &history_list, nil
}
