package appserver

import (
	"context"
	log "github.com/sirupsen/logrus"
	api "github.com/mxc-foundation/mxprotocol-server/m2m/api/m2m_ui"
	"github.com/mxc-foundation/mxprotocol-server/m2m/db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/config"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/services/wallet"
)

func (s *M2MServerAPI) GetWalletBalance(ctx context.Context, req *api.GetWalletBalanceRequest) (*api.GetWalletBalanceResponse, error) {
	log.WithFields(log.Fields{
		"orgId": req.OrgId,
	}).Debug("grpc_api/GetWalletBalance")

	balance, err := wallet.GetBalance(req.OrgId)
	if err != nil {
		log.WithError(err).Error("grpc_api/GetWalletBalance")
		return &api.GetWalletBalanceResponse{}, nil
	}

	return &api.GetWalletBalanceResponse{Balance: balance}, nil
}

func (s *M2MServerAPI) GetVmxcTxHistory(ctx context.Context, req *api.GetVmxcTxHistoryRequest) (*api.GetVmxcTxHistoryResponse, error) {
	log.WithFields(log.Fields{
		"orgId":  req.OrgId,
		"offset": req.Offset,
		"limit":  req.Limit,
	})

	return &api.GetVmxcTxHistoryResponse{}, nil
}

func (s *M2MServerAPI) GetWalletUsageHist(ctx context.Context, req *api.GetWalletUsageHistRequest) (*api.GetWalletUsageHistResponse, error) {
	walletId, err := db.Wallet.GetWalletIdFromOrgId(req.OrgId)
	if err != nil {
		log.WithError(err).Error("grpc_api/GetWalletUsageHist")
		return &api.GetWalletUsageHistResponse{}, nil
	}

	offset := req.Offset * req.Limit

	wuList, err := db.AggWalletUsage.GetWalletUsageHist(walletId, offset, req.Limit)
	if err != nil {
		log.WithError(err).Error("grpc_api/GetWalletUsageHist")
		return &api.GetWalletUsageHistResponse{}, nil
	}

	count, err := db.AggWalletUsage.GetWalletUsageHistCnt(walletId)
	if err != nil {
		log.WithError(err).Error("grpc_api/GetWalletUsageHist")
		return &api.GetWalletUsageHistResponse{}, nil
	}

	resp := &api.GetWalletUsageHistResponse{}
	resp.Count = count

	for _, v := range wuList {
		wuHist := &api.GetWalletUsageHist{}
		wuHist.StartAt = v.StartAt.String()
		wuHist.DurationMinutes = v.DurationMinutes
		wuHist.DlCntDv = v.DlCntDv
		wuHist.DlCntDvFree = v.DlCntDvFree
		wuHist.DlCntGw = v.DlCntGw
		wuHist.DlCntGwFree = v.DlCntGwFree
		wuHist.UlCntDv = v.UlCntDv
		wuHist.UlCntDvFree = v.UlCntDvFree
		wuHist.UlCntGw = v.UlCntGw
		wuHist.UlCntGwFree = v.UlCntGwFree
		wuHist.Income = v.Income
		wuHist.Spend = v.Spend
		wuHist.UpdatedBalance = v.UpdatedBalance

		resp.WalletUsageHis = append(resp.WalletUsageHis, wuHist)
	}
	return resp, nil
}

func (s *M2MServerAPI) GetDlPrice(ctx context.Context, req *api.GetDownLinkPriceRequest) (*api.GetDownLinkPriceResponse, error) {
	dlPrice := config.Cstruct.SuperNode.DlPrice
	return &api.GetDownLinkPriceResponse{
		DownLinkPrice: dlPrice,
	}, nil
}
