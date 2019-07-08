package supernode

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

func Setup() error {
	//todo
	ticker_superAccount := time.NewTicker(time.Duration(config.Cstruct.SuperNode.CheckAccountSeconds) * time.Second)
	go func() {
		log.Info("start supernode goroutine")
		for range ticker_superAccount.C {
			//TODO: should change the currAbv

			supernodeAccount, err := db.DbGetSuperNodeExtAccountAdr(config.Cstruct.SuperNode.ExtCurrAbv)
			if err != nil {
				log.Error("Cannot get supernode account: ", err)
			}

			checkTokenTx(config.Cstruct.SuperNode.ContractAddress, supernodeAccount)
		}
	}()

	/*ticker_checkPayment := time.NewTicker(time.Duration(config.Cstruct.SuperNode.CheckPaymentSecond) * time.Second)
	go func(reqID_paymentservice, withdrawId) {
		log.Info("start checkPay goroutine")
		for range ticker_checkPayment.C {

			reply, err := withdraw.CheckTxStatus(&config.Cstruct, reqID_paymentservice)
			if err != nil {
				//TODO
			}

			if reply.Error != "" {
				log.Error("CheckTxStatusReply Error: ", reply.Error)
			}

			if reply.TxPaymentStatusEnum != 2 {
				//ToDo: save it into db and check in next round
			} else {
				//ToDo: save it into db
				reply.TxHash
				reply.TxSentTime
				reply.TxPaymentStatusEnum
				db.DbUpdateWithdrawSuccessful(withdrawId)
				return
			}
		}
	}()*/

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
