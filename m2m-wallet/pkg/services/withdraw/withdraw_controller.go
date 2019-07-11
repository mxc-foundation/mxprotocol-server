package withdraw

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/auth"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var ctxWithdraw struct {
	withdrawFee map[string]float64
}

func Setup(conf config.MxpConfig) error {
	log.Info("Setup withdraw service")

	ctxWithdraw.withdrawFee = make(map[string]float64)
	for _, v := range db.CurrencyList {
		ctxWithdraw.withdrawFee[v.Abv] = 20
	}

	PaymentServiceAvailable = paymentServiceAvailable(conf)
	if false == PaymentServiceAvailable {
		log.Warning("service/withdraw")
	}

	for _, v := range db.CurrencyList {
		withdrawFee, err := db.DbGetActiveWithdrawFee(v.Abv)
		if err != nil {
			if _, err := db.DbInsertWithdrawFee(v.Abv, ctxWithdraw.withdrawFee[v.Abv]); err != nil {
				log.WithError(err).Error("service/withdraw")
				return err
			}
		} else {
			ctxWithdraw.withdrawFee[v.Abv] = withdrawFee
		}
	}

	return nil
}

type WithdrawServerAPI struct {
	serviceName string
}

func NewWithdrawServerAPI() *WithdrawServerAPI {
	return &WithdrawServerAPI{serviceName: "withdraw"}
}

func (s *WithdrawServerAPI) ModifyWithdrawFee(ctx context.Context, in *api.ModifyWithdrawFeeRequest) (*api.ModifyWithdrawFeeResponse, error) {
	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	log.WithFields(log.Fields{
		"moneyAbbr": api.Money_name[int32(in.MoneyAbbr)],
		"withdrawFee": in.WithdrawFee,
	}).Debug("grpc_api/ModifyWithdrawFee")

	if _, err := db.DbInsertWithdrawFee(api.Money_name[int32(in.MoneyAbbr)], in.WithdrawFee); err != nil {
		return &api.ModifyWithdrawFeeResponse{Status: false, UserProfile: &userProfile}, err
	}

	return &api.ModifyWithdrawFeeResponse{Status: true, UserProfile: &userProfile}, nil
}

func (s *WithdrawServerAPI) GetWithdrawFee(ctx context.Context, req *api.GetWithdrawFeeRequest) (*api.GetWithdrawFeeResponse, error) {
	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	log.WithFields(log.Fields{
		"moneyAbbr": api.Money_name[int32(req.MoneyAbbr)],
	}).Debug("grpc_api/GetWithdrawFee")

	extCurrencyAbbr := api.Money_name[int32(req.MoneyAbbr)]
	return &api.GetWithdrawFeeResponse{WithdrawFee: ctxWithdraw.withdrawFee[extCurrencyAbbr], UserProfile: &userProfile}, nil
}

func (s *WithdrawServerAPI) GetWithdrawHistory(ctx context.Context, req *api.GetWithdrawHistoryRequest) (*api.GetWithdrawHistoryResponse, error) {
	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	var count = int64(6)
	history_list := []*api.WithdrawHistory{}

	for i := 0; i < int(count); i++ {
		item := api.WithdrawHistory{
			From:      "a",
			To:        "b",
			MoneyType: "Ether",
			Amount:    12.333,
			CreatedAt: time.Now().UTC().String(),
		}

		history_list = append(history_list, &item)
	}

	return &api.GetWithdrawHistoryResponse{Count: count, WithdrawHistory: history_list, UserProfile: &userProfile}, nil
}

func (s *WithdrawServerAPI) WithdrawReq(ctx context.Context, req *api.WithdrawReqRequest) (*api.WithdrawReqResponse, error) {

	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	log.WithFields(log.Fields{
		"orgId": req.OrgId,
		"moneyAbbr": api.Money_name[int32(req.MoneyAbbr)],
		"amount": req.Amount,
	}).Debug("grpc_api/WithdrawReq")

	withdrawfee, err := db.DbGetActiveWithdrawFee(req.MoneyAbbr.String())
	if err != nil {
		return &api.WithdrawReqResponse{UserProfile: &userProfile}, status.Errorf(codes.DataLoss, "Cannot get withdrawfee from DB: %s", err)
	}

	walletID, err := db.DbGetWalletIdFromOrgId(req.OrgId)
	if err != nil {
		return &api.WithdrawReqResponse{UserProfile: &userProfile}, status.Errorf(codes.DataLoss, "Cannot get walletID from DB: %s", err)
	}

	balance, err := db.DbGetWalletBalance(walletID)
	if err != nil {
		return &api.WithdrawReqResponse{UserProfile: &userProfile}, status.Errorf(codes.DataLoss, "Cannot get wallet balance from DB: %s", err)
	}

	//check if the ext_account is enough in supernode
	if withdrawfee >= balance {
		return &api.WithdrawReqResponse{UserProfile: &userProfile}, status.Errorf(codes.Unavailable, "Not enough ext_account in super node")
	}

	withdrawID, err := db.DbInitWithdrawReq(walletID, req.Amount, req.MoneyAbbr.String())
	if err != nil {
		return &api.WithdrawReqResponse{UserProfile: &userProfile}, status.Errorf(codes.DataLoss, "Cannot get wallet withdrawID from DB: %s", err)
	}

	receiverAdd, err := db.DbGetUserExtAccountAdr(walletID, req.MoneyAbbr.String())
	if err != nil {
		return &api.WithdrawReqResponse{UserProfile: &userProfile}, status.Errorf(codes.DataLoss, "Cannot get user address from DB: %s", err)
	}
	reqIdClient, err := db.DbInitWithdrawReq(walletID, req.Amount, req.MoneyAbbr.String())
	amount := fmt.Sprintf("%f", req.Amount)
	reply, err := paymentReq(ctx, &config.Cstruct, amount, receiverAdd, reqIdClient)
	if err != nil {
		return &api.WithdrawReqResponse{UserProfile: &userProfile}, status.Errorf(codes.FailedPrecondition, "send payment request failed: %s", err)
	}

	// save reqqeryref to db
	err = db.DbUpdateWithdrawPaymentQueryId(walletID, reply.ReqQueryRef)
	if err != nil {
		return &api.WithdrawReqResponse{UserProfile: &userProfile}, status.Errorf(codes.DataLoss, "Cannot update queryID to DB: %s", err)
	}

	// make a new goroutine for checking the payment service
	go func() {
		for {
			time.Sleep(30 * time.Second)

			reply, err := CheckTxStatus(&config.Cstruct, reply.ReqQueryRef)
			if err != nil {
				log.Error("Cannot get the reply from paymentService: ", err)
				continue
			}

			if reply.Error != "" {
				log.Error("CheckTxStatusReply Error: ", reply.Error)
				continue
			}

			if reply.TxPaymentStatusEnum != 2 {
				log.Info("Still pending...")
				continue
			} else {
				timeStamp, err := time.Parse(reply.TxSentTime, "Mon Jan 2 15:04:05 -0700 MST 2006")
				if err != nil {
					log.Error("Time format error: ", err)
				}

				//Update withdrawID it into db
				err = db.DbUpdateWithdrawSuccessful(withdrawID, reply.TxHash, timeStamp)
				if err != nil {
					log.Error("Cannot update withdrawID to db: ", err)
					continue
				}
				return
			}
		}
	}()

	return &api.WithdrawReqResponse{Status: true, UserProfile: &userProfile}, nil
}
