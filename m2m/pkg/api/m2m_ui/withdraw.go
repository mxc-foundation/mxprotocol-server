package ui

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	api "github.com/mxc-foundation/mxprotocol-server/m2m/api/m2m_ui"
	"github.com/mxc-foundation/mxprotocol-server/m2m/db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/auth"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/config"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/services/wallet"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/services/withdraw"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

type WithdrawServerAPI struct {
	serviceName string
}

func NewWithdrawServerAPI() *WithdrawServerAPI {
	return &WithdrawServerAPI{serviceName: "withdraw"}
}

func (s *WithdrawServerAPI) ModifyWithdrawFee(ctx context.Context, in *api.ModifyWithdrawFeeRequest) (*api.ModifyWithdrawFeeResponse, error) {
	if in.OrgId != 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Only superadmin is authorized to modify withdraw fee.")
	}

	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, in.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.ModifyWithdrawFeeResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:

		log.WithFields(log.Fields{
			"moneyAbbr":   api.Money_name[int32(in.MoneyAbbr)],
			"withdrawFee": in.WithdrawFee,
		}).Debug("grpc_api/ModifyWithdrawFee")

		if _, err := db.WithdrawFee.InsertWithdrawFee(api.Money_name[int32(in.MoneyAbbr)], in.WithdrawFee); err != nil {
			log.WithError(err).Error("grpc_api/ModifyWithdrawFee")
			return &api.ModifyWithdrawFeeResponse{Status: false, UserProfile: &userProfile}, nil
		}

		withdraw.WithdrawFee[api.Money_name[int32(in.MoneyAbbr)]] = in.WithdrawFee
		return &api.ModifyWithdrawFeeResponse{Status: true, UserProfile: &userProfile}, nil

	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *WithdrawServerAPI) GetWithdrawFee(ctx context.Context, req *api.GetWithdrawFeeRequest) (*api.GetWithdrawFeeResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.GetWithdrawFeeResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:

		log.WithFields(log.Fields{
			"moneyAbbr": api.Money_name[int32(req.MoneyAbbr)],
		}).Debug("grpc_api/GetWithdrawFee")

		withdrawFee, err := db.WithdrawFee.GetActiveWithdrawFee(api.Money_name[int32(req.MoneyAbbr)])
		if err != nil {
			return &api.GetWithdrawFeeResponse{UserProfile: &userProfile}, nil
		}

		return &api.GetWithdrawFeeResponse{WithdrawFee: withdrawFee, UserProfile: &userProfile}, nil

	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *WithdrawServerAPI) GetWithdrawHistory(ctx context.Context, req *api.GetWithdrawHistoryRequest) (*api.GetWithdrawHistoryResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.GetWithdrawHistoryResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:

		log.WithFields(log.Fields{
			"orgId":  req.OrgId,
			"offset": req.Offset,
			"limit":  req.Limit,
		}).Debug("grpc_api/GetWithdrawHistory")

		walletId, err := wallet.GetWalletId(req.OrgId)
		if err != nil {
			log.WithError(err).Error("grpc_api/GetWithdrawHistory")
			return &api.GetWithdrawHistoryResponse{UserProfile: &userProfile}, nil
		}

		response := api.GetWithdrawHistoryResponse{UserProfile: &userProfile}
		ptr, err := db.Withdraw.GetWithdrawHist(walletId, req.Offset*req.Limit, req.Limit)
		if err != nil {
			log.WithError(err).Error("grpc_api/GetWithdrawHistory")
			return &api.GetWithdrawHistoryResponse{UserProfile: &userProfile}, nil
		}

		for _, v := range ptr {
			if v.ExtCurrency != api.Money_name[int32(req.MoneyAbbr)] {
				continue
			}

			history := api.WithdrawHistory{}
			history.From = v.AcntSender
			history.To = v.AcntRcvr
			history.MoneyType = v.ExtCurrency
			history.Amount = v.Value
			history.WithdrawFee = v.WithdrawFee
			history.TxSentTime = v.TxSentTime.String()
			history.TxStatus = v.TxStatus
			history.TxApprovedTime = v.TxAprvdTime.String()
			history.TxHash = v.TxHash

			response.WithdrawHistory = append(response.WithdrawHistory, &history)
		}
		response.Count, err = db.Withdraw.GetWithdrawHistRecCnt(walletId)

		return &response, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *WithdrawServerAPI) WithdrawReq(ctx context.Context, req *api.WithdrawReqRequest) (*api.WithdrawReqResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.WithdrawReqResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:
		log.WithFields(log.Fields{
			"orgId":     req.OrgId,
			"moneyAbbr": api.Money_name[int32(req.MoneyAbbr)],
			"amount":    req.Amount,
		}).Debug("grpc_api/WithdrawReq")

		withdrawfee, err := db.WithdrawFee.GetActiveWithdrawFee(req.MoneyAbbr.String())
		if err != nil {
			return &api.WithdrawReqResponse{UserProfile: &userProfile}, status.Errorf(codes.DataLoss, "Cannot get withdrawfee from DB: %s", err)
		}

		walletID, err := db.Wallet.GetWalletIdFromOrgId(req.OrgId)
		if err != nil {
			return &api.WithdrawReqResponse{UserProfile: &userProfile}, status.Errorf(codes.DataLoss, "Cannot get walletID from DB: %s", err)
		}

		balance, err := db.Wallet.GetWalletBalance(walletID)
		if err != nil {
			return &api.WithdrawReqResponse{UserProfile: &userProfile}, status.Errorf(codes.DataLoss, "Cannot get wallet balance from DB: %s", err)
		}

		//check if the ext_account balance is enough in wallet
		if withdrawfee+req.Amount > balance {
			return &api.WithdrawReqResponse{UserProfile: &userProfile}, status.Errorf(codes.Unavailable, "Not enough balance in user wallet")
		}

		receiverAdd, err := db.ExtAccount.GetUserExtAccountAdr(walletID, req.MoneyAbbr.String())
		if err != nil {
			return &api.WithdrawReqResponse{UserProfile: &userProfile}, status.Errorf(codes.DataLoss, "Cannot get user address from DB: %s", err)
		}

		// also updated the balance and history
		withdrawID, err := db.Withdraw.InitWithdrawReq(walletID, req.Amount, req.MoneyAbbr.String())
		if err != nil {
			return &api.WithdrawReqResponse{UserProfile: &userProfile}, status.Errorf(codes.DataLoss, "Cannot get wallet withdrawID from DB: %s", err)
		}

		// make a new goroutine for checking the payment service
		go func() {
			paymentRoutine(&config.Cstruct, receiverAdd, walletID, withdrawID, req)
		}()
		return &api.WithdrawReqResponse{Status: true, UserProfile: &userProfile}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}

func paymentRoutine(conf *config.MxpConfig, receiverAdd string, walletID, withdrawID int64, req *api.WithdrawReqRequest) {
	amount := fmt.Sprintf("%f", req.Amount)
	paymentReply, err := withdraw.PaymentReq(&config.Cstruct, amount, receiverAdd, withdrawID)
	if err != nil {
		log.Error("paymentRoutine/send payment request failed: ", err)
		for {
			time.Sleep(time.Duration(config.Cstruct.Withdraw.ResendToPS) * time.Second)
			paymentReply, err = withdraw.PaymentReq(&config.Cstruct, amount, receiverAdd, withdrawID)
			if err != nil {
				continue
			}
			break
		}
	}

	// save fk_query_id_payment_service to db
	err = db.Withdraw.UpdateWithdrawPaymentQueryId(walletID, paymentReply.ReqQueryRef)
	if err != nil {
		log.Error(codes.DataLoss, "Cannot update queryID to DB: %s", err)
	}

	for {
		time.Sleep(time.Duration(conf.Withdraw.RecheckStat) * time.Second)
		statusReply, err := withdraw.CheckTxStatus(&config.Cstruct, paymentReply.ReqQueryRef)
		if err != nil {
			log.Error("Cannot get the reply from paymentService: ", err)
			continue
		}

		if statusReply.Error != "" {
			log.Error("CheckTxStatusReply Error: ", statusReply.Error)
			continue
		}

		status := fmt.Sprintf("%s", statusReply.TxPaymentStatusEnum)
		if status != withdraw.SUCCESSFUL {
			log.Info("Still pending..., the queryId is:", paymentReply.ReqQueryRef)
			continue
		} else {
			layout := "2006-01-02 15:04:05"
			idx := strings.Index(statusReply.TxSentTime, " +")
			tt := statusReply.TxSentTime[:idx]

			timeStamp, err := time.Parse(layout, tt)
			if err != nil {
				log.Error("Time format error: ", err)
			}

			//Update withdrawID it into db
			err = db.Withdraw.UpdateWithdrawSuccessful(withdrawID, statusReply.TxHash, timeStamp)
			if err != nil {
				log.Error("Cannot update withdrawID to db: ", err)
				continue
			}
			break
		}
	}
}
