package supernode

import (
	"context"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/auth"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/config"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/services/ext_account"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Setup() error {
	ticker_superAccount := time.NewTicker(time.Duration(config.Cstruct.SuperNode.CheckAccountSeconds) * time.Second)
	go func() {
		log.Info("Start supernode goroutine")
		for range ticker_superAccount.C {
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
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, 0)

	switch res.Type {
	case auth.JsonParseError:
		fallthrough
	case auth.ErrorInfoNotNull:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.AddSuperNodeMoneyAccountResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")
	case auth.OK:
		log.WithFields(log.Fields{
			"moneyAbbr":   api.Money_name[int32(in.MoneyAbbr)],
			"accountAddr": strings.ToLower(in.AccountAddr),
		}).Debug("grpc_api/AddSuperNodeMoneyAccount")

		err := ext_account.UpdateActiveExtAccount(0, strings.ToLower(in.AccountAddr), api.Money_name[int32(in.MoneyAbbr)])
		if err != nil {
			log.WithError(err).Error("grpc_api/AddSuperNodeMoneyAccount")
			return &api.AddSuperNodeMoneyAccountResponse{Status: false, UserProfile: &userProfile}, nil
		}

		return &api.AddSuperNodeMoneyAccountResponse{Status: true, UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *SupernodeServerAPI) GetSuperNodeActiveMoneyAccount(ctx context.Context, req *api.GetSuperNodeActiveMoneyAccountRequest) (*api.GetSuperNodeActiveMoneyAccountResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, 0)

	switch res.Type {
	case auth.JsonParseError:
		fallthrough
	case auth.ErrorInfoNotNull:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.GetSuperNodeActiveMoneyAccountResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:

		log.WithFields(log.Fields{
			"moneyAbbr": api.Money_name[int32(req.MoneyAbbr)],
		}).Debug("grpc_api/GetSuperNodeActiveMoneyAccount")

		accountAddr, err := db.DbGetSuperNodeExtAccountAdr(api.Money_name[int32(req.MoneyAbbr)])
		if err != nil {
			log.WithError(err).Error("grpc_api/GetSuperNodeActiveMoneyAccount")
			return &api.GetSuperNodeActiveMoneyAccountResponse{SupernodeActiveAccount: "", UserProfile: &userProfile}, nil
		}

		return &api.GetSuperNodeActiveMoneyAccountResponse{SupernodeActiveAccount: accountAddr, UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}
