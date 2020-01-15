package appserver

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	api "github.com/mxc-foundation/mxprotocol-server/m2m/api/m2m_ui"
	"github.com/mxc-foundation/mxprotocol-server/m2m/db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/services/wallet"
)

func (s *M2MServerAPI) GetTransactionsHistory(ctx context.Context, req *api.GetTransactionsHistoryRequest) (*api.GetTransactionsHistoryResponse, error) {
	log.WithFields(log.Fields{
		"orgId":  req.OrgId,
		"offset": req.Offset,
		"limit":  req.Limit,
	}).Debug("grpc_api/GetTransactionsHistory")

	walletId, err := wallet.GetWalletId(req.OrgId)
	if err != nil {
		log.WithError(err).Error("grpc_api/GetTransactionsHistory")
		return &api.GetTransactionsHistoryResponse{}, nil
	}

	response := api.GetTransactionsHistoryResponse{}
	start := req.Offset * req.Limit
	end := req.Offset*req.Limit + req.Limit

	topupCount, err := db.Topup.GetTopupHistRecCnt(walletId)
	withdrawCount, err := db.Withdraw.GetWithdrawHistRecCnt(walletId)
	response.Count = topupCount + withdrawCount

	if start > response.Count {
		return &response, nil
	}

	if end > response.Count {
		end = response.Count
	}

	topupArray, err := db.Topup.GetTopupHist(walletId, 0, req.Limit*(req.Offset+1))
	if err != nil {
		log.WithError(err).Error("grpc_api/GetTransactionsHistory")
		return &api.GetTransactionsHistoryResponse{}, nil
	}
	withdrawArray, err := db.Withdraw.GetWithdrawHist(walletId, 0, req.Limit*(req.Offset+1))
	if err != nil {
		log.WithError(err).Error("grpc_api/GetTransactionsHistory")
		return &api.GetTransactionsHistoryResponse{}, nil
	}

	if len(topupArray)+len(withdrawArray) == 0 {
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

	response.TransactionHistory = sumArray[start:end]

	return &response, nil
}

func (s *M2MServerAPI) GetTopUpHistory(ctx context.Context, req *api.GetTopUpHistoryRequest) (*api.GetTopUpHistoryResponse, error) {
	log.WithFields(log.Fields{
		"orgId":  req.OrgId,
		"offset": req.Offset,
		"limit":  req.Limit,
	}).Debug("grpc_api/GetTopUpHistory")

	walletId, err := wallet.GetWalletId(req.OrgId)
	if err != nil {
		log.WithError(err).Error("grpc_api/GetTopUpHistory")
		return &api.GetTopUpHistoryResponse{}, nil
	}

	response := api.GetTopUpHistoryResponse{}
	ptr, err := db.Topup.GetTopupHist(walletId, req.Offset*req.Limit, req.Limit)
	if err != nil {
		log.WithError(err).Error("grpc_api/GetTopUpHistory")
		return &api.GetTopUpHistoryResponse{}, nil
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

func (s *M2MServerAPI) GetTopUpDestination(ctx context.Context, req *api.GetTopUpDestinationRequest) (*api.GetTopUpDestinationResponse, error) {
	log.WithFields(log.Fields{
		"orgId":     req.OrgId,
		"moneyType": api.Money_name[int32(req.MoneyAbbr)],
	}).Debug("grpc_api/GetTopUpDestination")

	supernode_account, err := db.ExtAccount.GetSuperNodeExtAccountAdr(api.Money_name[int32(req.MoneyAbbr)])
	if err != nil {
		log.WithError(err).Error("grpc_api/GetTopUpDestination")
		return &api.GetTopUpDestinationResponse{ActiveAccount: ""}, err
	}

	return &api.GetTopUpDestinationResponse{ActiveAccount: supernode_account}, nil
}

func (s *M2MServerAPI) GetIncome(ctx context.Context, req *api.GetIncomeRequest) (*api.GetIncomeResponse, error) {
	amount, err := db.Wallet.GetSupernodeIncomeAmount(time.Time{}, time.Now())
	if err != nil {
		return &api.GetIncomeResponse{Amount: 0}, err
	}

	return &api.GetIncomeResponse{
		Amount: amount,
	}, nil
}
