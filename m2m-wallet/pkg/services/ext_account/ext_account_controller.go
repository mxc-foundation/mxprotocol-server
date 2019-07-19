package ext_account

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/auth"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/services/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Setup() error {
	log.Info("Setup ext_account service")
	return nil
}

func UpdateActiveExtAccount(orgId int64, newAccount string, currencyAbbr string) error {
	walletId, err := wallet.GetWalletId(orgId)
	if err != nil {
		return err
	}

	_, err = db.DBInsertExtAccount(walletId, newAccount, currencyAbbr)
	if err != nil {
		return err
	}

	return nil
}

type ExtAccountServerAPI struct {
	serviceName string
}

func NewMoneyServerAPI() *ExtAccountServerAPI {
	return &ExtAccountServerAPI{serviceName: "ext_account"}
}

func (s *ExtAccountServerAPI) ModifyMoneyAccount(ctx context.Context, req *api.ModifyMoneyAccountRequest) (*api.ModifyMoneyAccountResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.JsonParseError:
	case auth.ErrorInfoNotNull:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdDeleted:
		return &api.ModifyMoneyAccountResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:

		log.WithFields(log.Fields{
			"orgId": req.OrgId,
			"moneyAbbr": api.Money_name[int32(req.MoneyAbbr)],
			"accountAddr": req.CurrentAccount,
		}).Debug("grpc_api/ModifyMoneyAccount")

		err := UpdateActiveExtAccount(req.OrgId, req.CurrentAccount, api.Money_name[int32(req.MoneyAbbr)])
		if err != nil {
			log.WithError(err).Error("grpc_api/ModifyMoneyAccount")
			return &api.ModifyMoneyAccountResponse{Status: false, UserProfile: &userProfile}, nil
		}
		return &api.ModifyMoneyAccountResponse{Status: true, UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *ExtAccountServerAPI) GetChangeMoneyAccountHistory(ctx context.Context, req *api.GetMoneyAccountChangeHistoryRequest) (*api.GetMoneyAccountChangeHistoryResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.JsonParseError:
	case auth.ErrorInfoNotNull:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdDeleted:
		return &api.GetMoneyAccountChangeHistoryResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:

		log.WithFields(log.Fields{
			"orgId": req.OrgId,
			"offset": req.Offset,
			"limit": req.Limit,
			"moneyAbbr": api.Money_name[int32(req.MoneyAbbr)],
		}).Debug("grpc_api/GetChangeMoneyAccountHistory")

		return &api.GetMoneyAccountChangeHistoryResponse{UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *ExtAccountServerAPI) GetActiveMoneyAccount(ctx context.Context, req *api.GetActiveMoneyAccountRequest) (*api.GetActiveMoneyAccountResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.JsonParseError:
	case auth.ErrorInfoNotNull:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdDeleted:
		return &api.GetActiveMoneyAccountResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:

		log.WithFields(log.Fields{
			"orgId": req.OrgId,
			"moneyAbbr": api.Money_name[int32(req.MoneyAbbr)],
		}).Debug("grpc_api/GetActiveMoneyAccount")

		walletId, err := db.DbGetWalletIdFromOrgId(req.OrgId)
		if err != nil {
			log.WithError(err).Error("grpc_api/GetActiveMoneyAccount")
			return &api.GetActiveMoneyAccountResponse{ActiveAccount: "", UserProfile: &userProfile}, nil
		}

		accountAddr, err := db.DbGetUserExtAccountAdr(walletId, api.Money_name[int32(req.MoneyAbbr)])
		if err != nil {
			log.WithError(err).Error("grpc_api/GetActiveMoneyAccount")
			return &api.GetActiveMoneyAccountResponse{ActiveAccount: "", UserProfile: &userProfile}, nil
		}

		return &api.GetActiveMoneyAccountResponse{ActiveAccount: accountAddr, UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}
