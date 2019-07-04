package supernode

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
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		log.Info("start supernode goroutine")
		for range ticker.C{
			checkTokenTx("0x5Ca381bBfb58f0092df149bD3D243b08B9a8386e","0x8a96E17d85Bd897a88B547718865de990D2Fcb80")
		}
	}()
	log.Info("setup supernode service")
	return nil
}

type SupernodeServerAPI struct {
	serviceName string
}

func NewSupernodeServerAPI() *SupernodeServerAPI {
	return &SupernodeServerAPI{serviceName: "supernode"}
}

func (s *SupernodeServerAPI) GetSuperNodeActiveMoneyAccount(ctx context.Context, req *api.GetSuperNodeActiveMoneyAccountRequest) (*api.GetSuperNodeActiveMoneyAccountResponse, error) {
	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	//Todo: userInfo should be the information of users eg.id,name,org,etc. Use it to get data from DB.
	fmt.Println("username = ", userProfile.User.Username)
	return &api.GetSuperNodeActiveMoneyAccountResponse{SupernodeActiveAccount: "supernode_account", Error: ""}, nil
}
