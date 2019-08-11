package wallet

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/auth"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func Setup(conf config.MxpConfig) error {
	log.Info("setup wallet service")

	return nil
}

//  return option 1: 0, true        --> no wallet created yet
//  return option 2: 0, false       --> sql error
//  return option 3: walletId, true --> get walletId successfully
func userHasWallet(orgId int64) (int64, bool) {
	walletId, err := db.DbGetWalletIdFromOrgId(orgId)
	if err != nil {
		if strings.HasSuffix(err.Error(), db.DbError.NoRowQueryRes.Error()) {
			return 0, true
		}

		return 0, false
	}

	return walletId, true
}

func createWallet(orgId int64) (walletId int64, err error) {
	if 0 == orgId {
		walletId, err = db.DbInsertWallet(orgId, db.SUPER_ADMIN)
	} else {
		walletId, err = db.DbInsertWallet(orgId, db.USER)
	}
	if err != nil {
		return walletId, err
	}

	return walletId, nil
}

func GetWalletId(orgId int64) (walletId int64, err error) {
	var res bool

	walletId, res = userHasWallet(orgId)
	if true == res && 0 == walletId {
		if walletId, err = createWallet(orgId); err != nil {
			return 0, err
		}
	} else if false == res {
		err = errors.New("Failed to get walletId.")
		log.WithError(err).Error("pkg/wallet/GetWalletId")
		return 0, err
	}

	return walletId, nil
}

func GetBalance(orgId int64) (float64, error) {
	walletId, err := GetWalletId(orgId)
	if err != nil {
		return 0, err
	}

	balance, err := db.DbGetWalletBalance(walletId)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func UpdateBalance(orgId int64, oper PaymentCategory, deviceType DeviceType, amount float64) error {
	walletId, err := GetWalletId(orgId)
	if err != nil {
		return err
	}

	balance, err := db.DbGetWalletBalance(walletId)
	if err != nil {
		return err
	}

	for _, v := range operMap {
		if v.pc == oper && v.dt == deviceType {
			balance = v.operation(balance, amount)
		}
	}

	err = db.DbUpdateBalanceByWalletId(walletId, balance)
	if err != nil {
		return err
	}

	return nil
}

type WalletServerAPI struct {
	serviceName string
}

func NewWalletServerAPI() *WalletServerAPI {
	return &WalletServerAPI{serviceName: "wallet"}
}

func (s *WalletServerAPI) GetWalletBalance(ctx context.Context, req *api.GetWalletBalanceRequest) (*api.GetWalletBalanceResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.GetWalletBalanceResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:

		log.WithFields(log.Fields{
			"orgId": req.OrgId,
		}).Debug("grpc_api/GetWalletBalance")

		balance, err := GetBalance(req.OrgId)
		if err != nil {
			log.WithError(err).Error("grpc_api/GetWalletBalance")
			return &api.GetWalletBalanceResponse{UserProfile: &userProfile}, nil
		}

		return &api.GetWalletBalanceResponse{Balance: balance, UserProfile: &userProfile}, nil

	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *WalletServerAPI) GetVmxcTxHistory(ctx context.Context, req *api.GetVmxcTxHistoryRequest) (*api.GetVmxcTxHistoryResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.GetVmxcTxHistoryResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:

		log.WithFields(log.Fields{
			"orgId":  req.OrgId,
			"offset": req.Offset,
			"limit":  req.Limit,
		})

		return &api.GetVmxcTxHistoryResponse{UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}
