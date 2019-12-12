package appserver

import (
	"context"
	log "github.com/sirupsen/logrus"
	api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/appserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/auth"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *M2MServerAPI) GetWalletBalance(ctx context.Context, req *api.GetWalletBalanceRequest) (*api.GetWalletBalanceResponse, error) {
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

		balance, err := wallet.GetBalance(req.OrgId)
		if err != nil {
			log.WithError(err).Error("grpc_api/GetWalletBalance")
			return &api.GetWalletBalanceResponse{UserProfile: &userProfile}, nil
		}

		return &api.GetWalletBalanceResponse{Balance: balance, UserProfile: &userProfile}, nil

	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *M2MServerAPI) GetVmxcTxHistory(ctx context.Context, req *api.GetVmxcTxHistoryRequest) (*api.GetVmxcTxHistoryResponse, error) {
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

func (s *M2MServerAPI) GetWalletUsageHist(ctx context.Context, req *api.GetWalletUsageHistRequest) (*api.GetWalletUsageHistResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.GetWalletUsageHistResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:
		walletId, err := db.Wallet.GetWalletIdFromOrgId(req.OrgId)
		if err != nil {
			log.WithError(err).Error("grpc_api/GetWalletUsageHist")
			return &api.GetWalletUsageHistResponse{UserProfile: &userProfile}, nil
		}

		offset := req.Offset * req.Limit

		wuList, err := db.AggWalletUsage.GetWalletUsageHist(walletId, offset, req.Limit)
		if err != nil {
			log.WithError(err).Error("grpc_api/GetWalletUsageHist")
			return &api.GetWalletUsageHistResponse{UserProfile: &userProfile}, nil
		}

		count, err := db.AggWalletUsage.GetWalletUsageHistCnt(walletId)
		if err != nil {
			log.WithError(err).Error("grpc_api/GetWalletUsageHist")
			return &api.GetWalletUsageHistResponse{UserProfile: &userProfile}, nil
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

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *M2MServerAPI) GetDlPrice(ctx context.Context, req *api.GetDownLinkPriceRequest) (*api.GetDownLinkPriceResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		return &api.GetDownLinkPriceResponse{UserProfile: &userProfile},
			status.Errorf(codes.NotFound, "This organization has been deleted from this user's profile.")

	case auth.OK:
		dlPrice := config.Cstruct.SuperNode.DlPrice
		return &api.GetDownLinkPriceResponse{
			DownLinkPrice: dlPrice,
			UserProfile:   &userProfile,
		}, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}