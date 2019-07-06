package supernode

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/auth"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/config"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/services/withdraw"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func Setup() error {
	//todo
	ticker_superAccount := time.NewTicker(time.Duration(config.Cstruct.SuperNode.CheckAccountSeconds) * time.Second)
	go func() {
		log.Info("start supernode goroutine")
		for range ticker_superAccount.C{
			//TODO: should change the super node address.
			checkTokenTx(config.Cstruct.SuperNode.ContractAddress,config.Cstruct.SuperNode.SuperNodeAddress)
		}
	}()

	ticker_checkPayment := time.NewTicker(time.Duration(config.Cstruct.SuperNode.CheckPaymentSecond) * time.Second)
	go func() {
		log.Info("start checkPay goroutine")
		for range ticker_checkPayment.C{
			//TODO: should change the super node address.
			reply, err := withdraw.CheckTxStatus(&config.Cstruct, 1)
			if err != nil{
				//TODO
			}
			if reply.TxPaymentStatusEnum == 1{
				//ToDo: resend the query
			}
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
