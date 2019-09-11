package topup

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/m2m"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/auth"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Setup() error {
	log.Info("Setup top_up service")
	return nil
}

type TopUpServerAPI struct {
	serviceName string
}

func NewTopUpServerAPI() *TopUpServerAPI {
	return &TopUpServerAPI{serviceName: "top up"}
}

func (s *TopUpServerAPI) GetTopUpHistory(ctx context.Context, req *m2m.GetTopUpHistoryRequest) (*m2m.GetTopUpHistoryResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &m2m.GetTopUpHistoryResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:

		log.WithFields(log.Fields{
			"orgId":  req.OrgId,
			"offset": req.Offset,
			"limit":  req.Limit,
		}).Debug("grpc_api/GetTopUpHistory")

		walletId, err := wallet.GetWalletId(req.OrgId)
		if err != nil {
			log.WithError(err).Error("grpc_api/GetTopUpHistory")
			return &m2m.GetTopUpHistoryResponse{UserProfile: &userProfile}, nil
		}

		response := m2m.GetTopUpHistoryResponse{UserProfile: &userProfile}
		ptr, err := db.Topup.GetTopupHist(walletId, req.Offset*req.Limit, req.Limit)
		if err != nil {
			log.WithError(err).Error("grpc_api/GetTopUpHistory")
			return &m2m.GetTopUpHistoryResponse{UserProfile: &userProfile}, nil
		}

		for _, v := range ptr {
			history := m2m.TopUpHistory{}
			history.From = v.AcntSender
			history.To = v.AcntRcvr
			history.Amount = v.Value
			history.CreatedAt = v.TxAprvdTime.String()
			history.MoneyType = v.ExtCurrency
			history.TxHash = v.TxHash

			response.TopupHistory = append(response.TopupHistory, &history)
		}
		response.Count, err = db.Topup.GetTopupHistRecCnt(walletId)

		return &response, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *TopUpServerAPI) GetTopUpDestination(ctx context.Context, req *m2m.GetTopUpDestinationRequest) (*m2m.GetTopUpDestinationResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &m2m.GetTopUpDestinationResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:

		log.WithFields(log.Fields{
			"orgId":     req.OrgId,
			"moneyType": m2m.Money_name[int32(req.MoneyAbbr)],
		}).Debug("grpc_api/GetTopUpDestination")

		supernode_account, err := db.ExtAccount.GetSuperNodeExtAccountAdr(m2m.Money_name[int32(req.MoneyAbbr)])
		if err != nil {
			log.WithError(err).Error("grpc_api/GetTopUpDestination")
			return &m2m.GetTopUpDestinationResponse{ActiveAccount: "", UserProfile: &userProfile}, err
		}

		return &m2m.GetTopUpDestinationResponse{ActiveAccount: supernode_account, UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}
