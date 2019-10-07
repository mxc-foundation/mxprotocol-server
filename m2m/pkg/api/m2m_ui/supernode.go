package ui

import (
	"context"
	log "github.com/sirupsen/logrus"
	api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/m2m_ui"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/auth"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services/ext_account"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type SupernodeServerAPI struct {
	serviceName string
}

func NewSupernodeServerAPI() *SupernodeServerAPI {
	return &SupernodeServerAPI{serviceName: "supernode"}
}

func (s *SupernodeServerAPI) AddSuperNodeMoneyAccount(ctx context.Context, in *api.AddSuperNodeMoneyAccountRequest) (*api.AddSuperNodeMoneyAccountResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, 0)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.AddSuperNodeMoneyAccountResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")
	case auth.OK:
		log.WithFields(log.Fields{
			"moneyAbbr":   api.Money_name[int32(in.MoneyAbbr)],
			"accountAddr": strings.ToLower(in.AccountAddr),
		}).Debug("grpc_api/AddSuperNodeMoneyAccount")

		err := ext_account.UpdateActiveExtAccount(0, in.AccountAddr, api.Money_name[int32(in.MoneyAbbr)])
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
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.GetSuperNodeActiveMoneyAccountResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:

		log.WithFields(log.Fields{
			"moneyAbbr": api.Money_name[int32(req.MoneyAbbr)],
		}).Debug("grpc_api/GetSuperNodeActiveMoneyAccount")

		accountAddr, err := ext_account.GetActiveExtAccount(0, api.Money_name[int32(req.MoneyAbbr)])
		if err != nil {
			log.WithError(err).Error("grpc_api/GetSuperNodeActiveMoneyAccount")
			return &api.GetSuperNodeActiveMoneyAccountResponse{SupernodeActiveAccount: "", UserProfile: &userProfile}, nil
		}

		return &api.GetSuperNodeActiveMoneyAccountResponse{SupernodeActiveAccount: accountAddr, UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}
