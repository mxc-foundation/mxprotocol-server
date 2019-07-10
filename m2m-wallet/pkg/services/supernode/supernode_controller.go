package supernode

import (
	"context"
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
	ticker := time.NewTicker(time.Duration(config.Cstruct.SuperNode.CheckAccountSeconds) * time.Second)
	go func() {
		log.Info("start supernode goroutine")
		for range ticker.C {
			supernodeAccount, err := db.DbGetSuperNodeExtAccountAdr(config.Cstruct.SuperNode.ExtCurrAbv)
			if err != nil {
				log.WithError(err).Warning("service/supernode")
				continue
			}

			err = checkTokenTx(config.Cstruct.SuperNode.ContractAddress, supernodeAccount, config.Cstruct.SuperNode.ExtCurrAbv)
			if err != nil {
				log.Warning("Restarting...")
				continue
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

func (s *SupernodeServerAPI) AddSuperNodeMoneyAccount(ctx context.Context, in *api.AddSuperNodeMoneyAccountRequest) (*api.AddSuperNodeMoneyAccountResponse, error) {
	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	walletId, err := db.DbInsertWallet(0, db.SUPER_ADMIN)
	if err != nil {
		return &api.AddSuperNodeMoneyAccountResponse{Status: false, UserProfile: &userProfile}, err
	}

	_, err = db.DBInsertExtAccount(walletId, in.AccountAddr, api.Money_name[int32(in.MoneyAbbr)])
	if err != nil {
		return &api.AddSuperNodeMoneyAccountResponse{Status: false, UserProfile: &userProfile}, err
	}

	return &api.AddSuperNodeMoneyAccountResponse{Status: true, UserProfile: &userProfile}, nil
}

func (s *SupernodeServerAPI) GetSuperNodeActiveMoneyAccount(ctx context.Context, req *api.GetSuperNodeActiveMoneyAccountRequest) (*api.GetSuperNodeActiveMoneyAccountResponse, error) {
	_, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	superNodeAddr, err := db.DbGetSuperNodeExtAccountAdr(config.Cstruct.SuperNode.ExtCurrAbv)
	if err != nil {
		return nil, status.Errorf(codes.DataLoss, "storage: Cannot get supernode account from DB: %s", err)
	}

	return &api.GetSuperNodeActiveMoneyAccountResponse{SupernodeActiveAccount: superNodeAddr}, nil
}
