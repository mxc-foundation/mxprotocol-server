package ext_account

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/m2m"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/auth"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func Setup() error {
	log.Info("Setup ext_account service")
	return nil
}

func UpdateActiveExtAccount(orgId int64, newAccount string, currencyAbbr string) error {
	walletId, err := wallet.GetWalletId(orgId)
	if err != nil {
		log.WithError(err).Error("service/UpdateActiveExtAccount")
		return err
	}

	_, err = db.ExtAccount.InsertExtAccount(walletId, newAccount, currencyAbbr)
	if err != nil {
		log.WithError(err).Error("service/UpdateActiveExtAccount")
		return err
	}

	return nil
}

func GetActiveExtAccount(orgId int64, currencyAbbr string) (string, error) {
	walletId, err := wallet.GetWalletId(orgId)
	if err != nil {
		log.WithError(err).Error("service/GetActiveExtAccount")
		return "", err
	}

	var accountAddr string
	if orgId == 0 {
		accountAddr, err = db.ExtAccount.GetSuperNodeExtAccountAdr(currencyAbbr)
	} else {
		accountAddr, err = db.ExtAccount.GetUserExtAccountAdr(walletId, currencyAbbr)
	}

	if err != nil {
		if strings.HasSuffix(err.Error(), db.DbError.NoRowQueryRes.Error()) {
			log.Warnf("service/GetActiveExtAccount: get account with walletId=%d, currency=%s", walletId, currencyAbbr)
			return "", nil
		}
		log.WithError(err).Error("service/GetActiveExtAccount")
		return "", err
	}

	return accountAddr, nil
}

type ExtAccountServerAPI struct {
	serviceName string
}

func NewMoneyServerAPI() *ExtAccountServerAPI {
	return &ExtAccountServerAPI{serviceName: "ext_account"}
}

func (s *ExtAccountServerAPI) ModifyMoneyAccount(ctx context.Context, req *m2m.ModifyMoneyAccountRequest) (*m2m.ModifyMoneyAccountResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &m2m.ModifyMoneyAccountResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:
		log.WithFields(log.Fields{
			"orgId":       req.OrgId,
			"moneyAbbr":   m2m.Money_name[int32(req.MoneyAbbr)],
			"accountAddr": strings.ToLower(req.CurrentAccount),
		}).Debug("grpc_api/ModifyMoneyAccount")

		if 0 == req.OrgId {
			return &m2m.ModifyMoneyAccountResponse{Status: false, UserProfile: &userProfile}, nil
		}

		err := UpdateActiveExtAccount(req.OrgId, req.CurrentAccount, m2m.Money_name[int32(req.MoneyAbbr)])
		if err != nil {
			log.WithError(err).Error("grpc_api/ModifyMoneyAccount")
			return &m2m.ModifyMoneyAccountResponse{Status: false, UserProfile: &userProfile},
				status.Errorf(codes.InvalidArgument, "Duplicate or invalid format.")
		}
		return &m2m.ModifyMoneyAccountResponse{Status: true, UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *ExtAccountServerAPI) GetChangeMoneyAccountHistory(ctx context.Context, req *m2m.GetMoneyAccountChangeHistoryRequest) (*m2m.GetMoneyAccountChangeHistoryResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &m2m.GetMoneyAccountChangeHistoryResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:

		log.WithFields(log.Fields{
			"orgId":     req.OrgId,
			"offset":    req.Offset,
			"limit":     req.Limit,
			"moneyAbbr": m2m.Money_name[int32(req.MoneyAbbr)],
		}).Debug("grpc_api/GetChangeMoneyAccountHistory")

		walletId, err := wallet.GetWalletId(req.OrgId)
		if err != nil {
			log.WithError(err).Error("grpc_api/GetChangeMoneyAccountHistory")
			return &m2m.GetMoneyAccountChangeHistoryResponse{UserProfile: &userProfile}, nil
		}

		response := m2m.GetMoneyAccountChangeHistoryResponse{UserProfile: &userProfile}
		ptr, err := db.ExtAccount.GetExtAcntHist(walletId, req.Offset*req.Limit, req.Limit)
		if err != nil {
			log.WithError(err).Error("grpc_api/GetChangeMoneyAccountHistory")
			return &m2m.GetMoneyAccountChangeHistoryResponse{UserProfile: &userProfile}, nil
		}

		for _, v := range ptr {
			if v.ExtCurrencyAbv != m2m.Money_name[int32(req.MoneyAbbr)] {
				continue
			}
			history := m2m.MoneyAccountChangeHistory{}
			history.Addr = v.AccountAdr
			history.CreatedAt = v.InsertTime.String()
			history.Status = v.Status
			response.ChangeHistory = append(response.ChangeHistory, &history)
		}
		response.Count, err = db.ExtAccount.GetExtAcntHistRecCnt(walletId)

		return &response, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *ExtAccountServerAPI) GetActiveMoneyAccount(ctx context.Context, req *m2m.GetActiveMoneyAccountRequest) (*m2m.GetActiveMoneyAccountResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &m2m.GetActiveMoneyAccountResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:

		log.WithFields(log.Fields{
			"orgId":     req.OrgId,
			"moneyAbbr": m2m.Money_name[int32(req.MoneyAbbr)],
		}).Debug("grpc_api/GetActiveMoneyAccount")

		accountAddr, err := GetActiveExtAccount(req.OrgId, m2m.Money_name[int32(req.MoneyAbbr)])
		if err != nil {
			log.WithError(err).Error("grpc_api/GetActiveMoneyAccount")
			return &m2m.GetActiveMoneyAccountResponse{ActiveAccount: "", UserProfile: &userProfile}, nil
		}

		return &m2m.GetActiveMoneyAccountResponse{ActiveAccount: accountAddr, UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}
