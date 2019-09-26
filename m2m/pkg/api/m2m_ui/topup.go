package ui

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/m2m_ui"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/auth"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TopUpServerAPI struct {
	serviceName string
}

func NewTopUpServerAPI() *TopUpServerAPI {
	return &TopUpServerAPI{serviceName: "top up"}
}

func (s *TopUpServerAPI) GetTransactionsHistory(ctx context.Context, req *api.GetTransactionsHistoryRequest) (*api.GetTransactionsHistoryResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.GetTransactionsHistoryResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:

		log.WithFields(log.Fields{
			"orgId":  req.OrgId,
			"offset": req.Offset,
			"limit":  req.Limit,
		}).Debug("grpc_api/GetTransactionsHistory")

		walletId, err := wallet.GetWalletId(req.OrgId)
		if err != nil {
			log.WithError(err).Error("grpc_api/GetTransactionsHistory")
			return &api.GetTransactionsHistoryResponse{UserProfile: &userProfile}, nil
		}

		response := api.GetTransactionsHistoryResponse{UserProfile: &userProfile}
		topupCount, err := db.Topup.GetTopupHistRecCnt(walletId)
		withdrawCount, err := db.Withdraw.GetWithdrawHistRecCnt(walletId)
		response.Count = topupCount + withdrawCount

		topupArray, err := db.Topup.GetTopupHist(walletId, 0, req.Limit*(req.Offset+1))
		if err != nil {
			log.WithError(err).Error("grpc_api/GetTransactionsHistory")
			return &api.GetTransactionsHistoryResponse{UserProfile: &userProfile}, nil
		}
		withdrawArray, err := db.Withdraw.GetWithdrawHist(walletId, 0, req.Limit*(req.Offset+1))
		if err != nil {
			log.WithError(err).Error("grpc_api/GetTransactionsHistory")
			return &api.GetTransactionsHistoryResponse{UserProfile: &userProfile}, nil
		}

		if len(topupArray)+len(withdrawArray) == 0 {
			response.UserProfile = &userProfile
			return &response, nil
		}

		sumArray := make([]*api.TransactionsHistory, len(topupArray)+len(withdrawArray))

		m := 0
		for i := 0; i < len(topupArray); i++ {
			tmp := api.TransactionsHistory{}
			tmp.From = topupArray[i].AcntSender
			tmp.To = topupArray[i].AcntRcvr
			tmp.MoneyAbbr = topupArray[i].ExtCurrency
			tmp.Amount = topupArray[i].Value
			tmp.Status = "SUCCESSFUL"
			tmp.TxHash = topupArray[i].TxHash
			tmp.TransactionType = "Topup"
			timestampConv, _ := ptypes.TimestampProto(topupArray[i].TxAprvdTime)
			tmp.LastUpdateTime = timestampConv

			sumArray[m] = &tmp
			m++
		}

		for i := 0; i < len(withdrawArray); i++ {
			tmp := api.TransactionsHistory{}
			tmp.From = withdrawArray[i].AcntSender
			tmp.To = withdrawArray[i].AcntRcvr
			tmp.MoneyAbbr = withdrawArray[i].ExtCurrency
			tmp.Amount = withdrawArray[i].Value
			tmp.Status = withdrawArray[i].TxStatus
			tmp.TxHash = withdrawArray[i].TxHash
			tmp.TransactionType = "Withdraw"
			timestampConv, _ := ptypes.TimestampProto(withdrawArray[i].TxAprvdTime)
			tmp.LastUpdateTime = timestampConv

			sumArray[m] = &tmp
			m++
		}

		// sort sumArray by LastUpdateTime, from latest to earliest
		var swapped bool
		for l := 0; l < m; l++ {
			swapped = false

			for k := 0; k < m-l-1; k++ {
				if sumArray[k].LastUpdateTime.Seconds < sumArray[k+1].LastUpdateTime.Seconds {
					sumArray[k], sumArray[k+1] = sumArray[k+1], sumArray[k]
					swapped = true
				}
			}

			if swapped == false {
				break
			}
		}

		response.TransactionHistory = sumArray[(req.Offset * req.Limit):(req.Offset*req.Limit + req.Limit)]

		return &response, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
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
		ptr, err := db.Topup.GetTopupHist(walletId, req.Offset*req.Limit, req.Limit)
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
		response.Count, err = db.Topup.GetTopupHistRecCnt(walletId)

		return &response, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *TopUpServerAPI) GetTopUpDestination(ctx context.Context, req *api.GetTopUpDestinationRequest) (*api.GetTopUpDestinationResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.GetTopUpDestinationResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:

		log.WithFields(log.Fields{
			"orgId":     req.OrgId,
			"moneyType": api.Money_name[int32(req.MoneyAbbr)],
		}).Debug("grpc_api/GetTopUpDestination")

		supernode_account, err := db.ExtAccount.GetSuperNodeExtAccountAdr(api.Money_name[int32(req.MoneyAbbr)])
		if err != nil {
			log.WithError(err).Error("grpc_api/GetTopUpDestination")
			return &api.GetTopUpDestinationResponse{ActiveAccount: "", UserProfile: &userProfile}, err
		}

		return &api.GetTopUpDestinationResponse{ActiveAccount: supernode_account, UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}
