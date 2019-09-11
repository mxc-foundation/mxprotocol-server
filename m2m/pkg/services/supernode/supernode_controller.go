package supernode

import (
	"context"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/m2m"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/auth"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services/ext_account"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Setup(conf config.MxpConfig) error {
	supernode_account, err := ext_account.GetActiveExtAccount(0, conf.SuperNode.ExtCurrAbv)
	if err != nil {
		return err
	} else if supernode_account == "" {
		err := ext_account.UpdateActiveExtAccount(0, conf.SuperNode.SuperNodeAddress, conf.SuperNode.ExtCurrAbv)
		if err != nil {
			return err
		}
	}

	ticker_superAccount := time.NewTicker(time.Duration(config.Cstruct.SuperNode.CheckAccountSeconds) * time.Second)
	go func() {
		log.Info("Start supernode goroutine")
		for range ticker_superAccount.C {
			supernodeAccount, err := ext_account.GetActiveExtAccount(0, conf.SuperNode.ExtCurrAbv)
			if err != nil {
				log.WithError(err).Warning("service/supernode")
				continue
			}

			err = checkTokenTx(conf.SuperNode.ContractAddress, supernodeAccount, conf.SuperNode.ExtCurrAbv)
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

func (s *SupernodeServerAPI) AddSuperNodeMoneyAccount(ctx context.Context, in *m2m.AddSuperNodeMoneyAccountRequest) (*m2m.AddSuperNodeMoneyAccountResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, 0)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &m2m.AddSuperNodeMoneyAccountResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")
	case auth.OK:
		log.WithFields(log.Fields{
			"moneyAbbr":   m2m.Money_name[int32(in.MoneyAbbr)],
			"accountAddr": strings.ToLower(in.AccountAddr),
		}).Debug("grpc_api/AddSuperNodeMoneyAccount")

		err := ext_account.UpdateActiveExtAccount(0, in.AccountAddr, m2m.Money_name[int32(in.MoneyAbbr)])
		if err != nil {
			log.WithError(err).Error("grpc_api/AddSuperNodeMoneyAccount")
			return &m2m.AddSuperNodeMoneyAccountResponse{Status: false, UserProfile: &userProfile}, nil
		}

		return &m2m.AddSuperNodeMoneyAccountResponse{Status: true, UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *SupernodeServerAPI) GetSuperNodeActiveMoneyAccount(ctx context.Context, req *m2m.GetSuperNodeActiveMoneyAccountRequest) (*m2m.GetSuperNodeActiveMoneyAccountResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, 0)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &m2m.GetSuperNodeActiveMoneyAccountResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:

		log.WithFields(log.Fields{
			"moneyAbbr": m2m.Money_name[int32(req.MoneyAbbr)],
		}).Debug("grpc_api/GetSuperNodeActiveMoneyAccount")

		accountAddr, err := ext_account.GetActiveExtAccount(0, m2m.Money_name[int32(req.MoneyAbbr)])
		if err != nil {
			log.WithError(err).Error("grpc_api/GetSuperNodeActiveMoneyAccount")
			return &m2m.GetSuperNodeActiveMoneyAccountResponse{SupernodeActiveAccount: "", UserProfile: &userProfile}, nil
		}

		return &m2m.GetSuperNodeActiveMoneyAccountResponse{SupernodeActiveAccount: accountAddr, UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}
