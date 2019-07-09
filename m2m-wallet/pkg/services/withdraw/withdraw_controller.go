package withdraw

import (
	"context"
	"errors"
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
	log.Info("setup withdraw service")

	ctxWithdraw.withdrawFee = make(map[string]float64)
	for _, v := range db.CurrencyList {
		ctxWithdraw.withdrawFee[v.Abv] = 20
	}

	if false == paymentServiceAvailable(conf) {
		err := errors.New("Setup withdraw failed: payment service not available.")
		log.WithError(err).Error("service/withdraw")
		return err
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

func (s *WithdrawServerAPI) GetWithdrawFee(ctx context.Context, req *api.GetWithdrawFeeRequest) (*api.GetWithdrawFeeResponse, error) {
	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	extCurrencyAbbr := api.Money_name[int32(req.MoneyAbbr)]
	return &api.GetWithdrawFeeResponse{WithdrawFee: ctxWithdraw.withdrawFee[extCurrencyAbbr], Error: "", UserProfile: &userProfile}, nil
}

func (s *WithdrawServerAPI) GetWithdrawHistory(ctx context.Context, req *api.GetWithdrawHistoryRequest) (*api.GetWithdrawHistoryResponse, error) {
	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	//Todo: userInfo should be the information of users eg.id,name,org,etc. Use it to get data from DB.
	fmt.Println("username = ", userProfile.User.Username)

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

	return &api.GetWithdrawHistoryResponse{Error: "", Count: count, WithdrawHistory: history_list, UserProfile: &userProfile}, nil
}

func (s *WithdrawServerAPI) WithdrawReq(ctx context.Context, req *api.WithdrawReqRequest) (*api.WithdrawReqResponse, error) {
	//todo
	userProfile, err := auth.VerifyRequestViaAuthServer(ctx, s.serviceName)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	//withdrawfee, err := db.DbGetActiveWithdrawFee(req.MoneyAbbr.String())
	//balance, err := db.DbGetWalletBalance(walletID)
	//ToDo: Sum then check if the money is enough in supernode

	amount := fmt.Sprintf("%f", req.Amount)

	//ToDo: wait (for get the withdrawID)
	//walletID, err := db.DbGetWalletIdFromOrgId(req.OrgId)
	//withdrawID, err := db.DbInitWithdrawReq(walletID, amount, req.MoneyAbbr.String())

	reply, err := paymentReq(ctx, &config.Cstruct, amount)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "send payment request failed: %s", err)
	}

	reqId_paymentsev := reply.ReqQueryRef

	//ToDo: wait (save reqqeryref to db)
	//err := db.DbUpdateWithdrawPaymentQueryId(walletID, reqId_paymentsev)

	go func() {
		reply, err := CheckTxStatus(&config.Cstruct, reqId_paymentsev)
		if err != nil {
			log.Error("Cannot get the reply from paymentService: ", err)
		}

		if reply.Error != "" {
			log.Error("CheckTxStatusReply Error: ", reply.Error)
		}

		if reply.TxPaymentStatusEnum != 2 {
			log.Info("Pending...")
		} else {
			//reply.TxHash
			//reply.TxSentTime
			//reply.TxPaymentStatusEnum

			//ToDo: update withdrawID it into db
			//db.DbUpdateWithdrawSuccessful(withdrawID)

			//for test!
			fmt.Println("Update to DB successful")

			return
		}
	}()

	return &api.WithdrawReqResponse{Status: true, Error: "", UserProfile: &userProfile}, nil
}
