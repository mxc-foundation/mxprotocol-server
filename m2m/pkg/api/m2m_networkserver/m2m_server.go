package m2m_networkserver

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/networkserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type M2MNetworkServerAPI struct{}

// M2MNetworkServerServerAPI returns a new M2MServerAPI.
func NewM2MNetworkServerAPI() *M2MNetworkServerAPI {
	dlPrice = config.Cstruct.SuperNode.DlPrice
	return &M2MNetworkServerAPI{}
}

var ulPktCount int64
var timeLayout = "2006-01-02T15:04:05.000000Z"
var dlPrice float64

func (*M2MNetworkServerAPI) DvUsageMode(ctx context.Context, req *networkserver.DvUsageModeRequest) (resp *networkserver.DvUsageModeResponse, err error) {
	log.WithFields(log.Fields{
		"dvId": req.DvEui,
	}).Debug("grpc_api/DvUsageMode")

	dvWalletId, dvMode, err := db.Device.GetDevWalletIdAndModeByEui(req.DvEui)
	if err != nil {
		log.WithError(err).Error("db/cannot get dvWalletId and dvMode")
		return &networkserver.DvUsageModeResponse{}, err
	}

	switch string(dvMode) {
	case networkserver.DeviceMode_name[int32(networkserver.DeviceMode_DV_INACTIVE)]:
		return &networkserver.DvUsageModeResponse{DvMode: networkserver.DeviceMode_DV_INACTIVE}, nil

	case networkserver.DeviceMode_name[int32(networkserver.DeviceMode_DV_FREE_GATEWAYS_LIMITED)]:
		_, gwMac, err := db.Gateway.GetFreeGwList(dvWalletId)
		if err != nil {
			log.WithError(err).Error("db/cannot get gwListOfWallet")
			return &networkserver.DvUsageModeResponse{}, err
		}

		for _, v := range gwMac {
			fgws := networkserver.GwMac{}
			fgws.GwMac = v
			resp.FreeGwMac = append(resp.FreeGwMac, &fgws)
		}

		resp.DvMode = networkserver.DeviceMode_DV_FREE_GATEWAYS_LIMITED

		return resp, nil

	case networkserver.DeviceMode_name[int32(networkserver.DeviceMode_DV_WHOLE_NETWORK)]:
		walletId, err := db.Device.GetDevWalletIdByEui(req.DvEui)
		if err != nil {
			log.WithError(err).Error("db/cannot get walletId")
			return &networkserver.DvUsageModeResponse{}, err
		}

		balance, err := db.Wallet.GetWalletBalance(walletId)
		if err != nil {
			log.WithError(err).Error("db/cannot get balance")
			return &networkserver.DvUsageModeResponse{}, err
		}

		_, gwMac, err := db.Gateway.GetFreeGwList(dvWalletId)
		if err != nil {
			log.WithError(err).Error("db/cannot get gwListOfWallet")
			return &networkserver.DvUsageModeResponse{}, err
		}

		for _, v := range gwMac {
			fgws := networkserver.GwMac{}
			fgws.GwMac = v
			resp.FreeGwMac = append(resp.FreeGwMac, &fgws)
		}

		resp.DvMode = networkserver.DeviceMode_DV_WHOLE_NETWORK
		if balance < dlPrice {
			resp.EnoughBalance = false
		} else {
			resp.EnoughBalance = true
		}
		return resp, nil

	case networkserver.DeviceMode_name[int32(networkserver.DeviceMode_DV_DELETED)]:
		return &networkserver.DvUsageModeResponse{DvMode: networkserver.DeviceMode_DV_DELETED}, nil
	}

	return &networkserver.DvUsageModeResponse{}, nil
}

func (*M2MNetworkServerAPI) GwUsageMode(ctx context.Context, req *networkserver.GwUsageModeRequest) (*networkserver.GwUsageModeResponse, error) {
	log.WithFields(log.Fields{
		"gwId": req.GwMac,
	}).Debug("grpc_api/GwUsageMode")

	gwMode, err := db.Gateway.GetGatewayIdByMac(req.GwMac)
	if err != nil {
		log.WithError(err).Error("db/cannot get gwMode")
		return &networkserver.GwUsageModeResponse{}, err
	}

	return &networkserver.GwUsageModeResponse{GwMode: string(gwMode)}, nil
}

func (*M2MNetworkServerAPI) DlPktSent(ctx context.Context, req *networkserver.DlPktSentRequest) (*networkserver.DlPktSentResponse, error) {
	//create a new thread for db
	go func()(error) {
		log.WithFields(log.Fields{
			"DlPktId": req.DlPkt.DlIdNs,
		}).Debug("grpc_api/DlPktSent")

		var dlPkt = types.DlPkt{}
		dlPkt.DlIdNs = req.DlPkt.DlIdNs
		dlPkt.Category = types.DlCategory(req.DlPkt.Category)
		dlPkt.Nonce = req.DlPkt.Nonce

		if createAt, err := time.Parse(timeLayout, req.DlPkt.CreateAt); err != nil {
			log.WithError(err).Error("time format error")
			return err
		} else {
			dlPkt.CreatedAt = createAt
		}

		dlPkt.Size = req.DlPkt.Size
		dlPkt.TokenDlFrm1 = req.DlPkt.TokenDlFrm1
		dlPkt.TokenDlFrm2 = req.DlPkt.TokenDlFrm2

		dvId, err := db.Device.GetDeviceIdByDevEui(req.DlPkt.DevEui)
		if err != nil {
			log.WithError(err).Error("db/cannot get devID")
			return err
		}

		dlPkt.FkDevice = dvId

		gwId, err := db.Gateway.GetGatewayIdByMac(req.DlPkt.GwMac)
		if err != nil {
			log.WithError(err).Error("db/cannot get gwID")
			return err
		}

		dlPkt.FkGateway = gwId

		_, err = db.DlPacket.InsertDlPkt(dlPkt)
		if err != nil {
			log.WithError(err).Error("db/cannot update dlPkt")
			return err
		}

		err = db.Device.UpdateDeviceLastSeen(dvId, time.Now())
		if err != nil {
			log.WithError(err).Error("db/cannot update devicelastseen")
			return err
		}

		dvwalletId, err := db.Device.GetWalletIdOfDevice(dvId)
		if err != nil {
			log.WithError(err).Error("db/cannot get dvWalletID")
			return err
		}

		gwwalletId, err := db.Gateway.GetWalletIdOfGateway(gwId)
		if err != nil {
			log.WithError(err).Error("db/cannot get gwWalletID")
			return err
		}

		if dvwalletId == gwwalletId {
			return nil
		} else {
			err = db.Wallet.TmpBalanceUpdatePktTx(dvId, gwId, dlPrice)
			if err != nil {
				log.WithError(err).Error("db/cannot update balance")
				return err
			}
		}

		return nil
	}()
	return &networkserver.DlPktSentResponse{}, nil
}

func (*M2MNetworkServerAPI) UlPktSent(ctx context.Context, req *networkserver.UlPktSentRequest) (*networkserver.UlPktSentResponse, error) {
	ulPktCount++
	println("Pkt no: ", ulPktCount, "!")
	return nil, nil
}
