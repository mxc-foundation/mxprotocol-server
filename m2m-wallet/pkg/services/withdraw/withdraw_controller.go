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

func init() {
	for _, v := range db.CurrencyList {
		ctxWithdraw.withdrawFee[v.Abv] = 20
	}
}

func Setup(conf config.MxpConfig) error {
	log.Info("setup withdraw service")

	if false == paymentServiceAvailable(conf) {
		err := errors.New("Setup withdraw failed: payment service not available.")
		log.WithError(err).Error("service/withdraw")
		return err
	}

	for _, v := range db.CurrencyList {
		withdrawFee, err := db.DbGetActiveWithdrawFee(v.Abv)
		if err != nil {
			db.DbInsertWithdrawFee(v.Abv, ctxWithdraw.withdrawFee[v.Abv])
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

	//Todo: userInfo should be the information of users eg.id,name,org,etc. Use it to get data from DB.
	fmt.Println("username = ", userProfile.User.Username)

	amount := fmt.Sprintf("%f", req.Amount)
	reply, err := paymentReq(ctx, &config.Cstruct, amount)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "send payment request failed: %s", err)
	}

	//Todo: save reqqueryref info into db
	fmt.Println(reply.ReqQueryRef)

	return &api.WithdrawReqResponse{Status: true, Error: "", UserProfile: &userProfile}, nil
}
