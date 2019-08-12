package topup

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
	log.Info("Setup top_up service")
	return nil
}

type TopUpServerAPI struct {
	serviceName string
}

func NewTopUpServerAPI() *TopUpServerAPI {
	return &TopUpServerAPI{serviceName: "top up"}
}

func (s *TopUpServerAPI) GetTopUpHistory(ctx context.Context, req *api.GetTopUpHistoryRequest) (*api.GetTopUpHistoryResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.GetTopUpHistoryResponse{UserProfile: &userProfile},
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
			return &api.GetTopUpHistoryResponse{UserProfile: &userProfile}, nil
		}

		response := api.GetTopUpHistoryResponse{UserProfile: &userProfile}
		ptr, err := db.DbGetTopupHist(walletId, req.Offset*req.Limit, req.Limit)
		if err != nil {
			log.WithError(err).Error("grpc_api/GetTopUpHistory")
			return &api.GetTopUpHistoryResponse{UserProfile: &userProfile}, nil
		}

		for _, v := range ptr {
			history := api.TopUpHistory{}
			history.From = v.AcntSender
			history.To = v.AcntRcvr
			history.Amount = v.Value
			history.CreatedAt = v.TxAprvdTime.String()
			history.MoneyType = v.ExtCurrency
			history.TxHash = v.TxHash

			response.TopupHistory = append(response.TopupHistory, &history)
		}
		response.Count, err = db.DbGetTopupHistRecCnt(walletId)

		return &response, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}
