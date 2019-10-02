package m2m_networkserver

import (
	"context"
	"fmt"
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

	var gw []*networkserver.GwMac
	gw1 := networkserver.GwMac{GwMac: "gw1"}
	gw = append(gw, &gw1)
	gw2 := networkserver.GwMac{GwMac: "40d63cfffe020f84"}
	gw = append(gw, &gw2)
	dvmd := networkserver.DeviceMode_DV_FREE_GATEWAYS_LIMITED

	if req.DvEui == "70b3d5fffe1cb16a" {
		gw2 = networkserver.GwMac{GwMac: "40d63cfffe030f5c"}
		dvmd = networkserver.DeviceMode_DV_WHOLE_NETWORK
	}

	return &networkserver.DvUsageModeResponse{
		DvMode:        dvmd,
		FreeGwMac:     gw,
		EnoughBalance: true,
	}, nil

	/*dvWalletId, err := db.Device.GetDevWalletIdByEui(req.DvEui)
	if err != nil {
		log.WithError(err).Error("db/cannot get dvWalletId")
		return &networkserver.DvUsageModeResponse{}, err
	}*/

	dvWalletId, dvMode, err := db.Device.GetDevWalletIdAndModeByEui(req.DvEui)
	if err != nil {
		log.WithError(err).Error("db/cannot get dvWalletId and dvMode")
		return &networkserver.DvUsageModeResponse{}, err
	}

	gwList, err := db.Gateway.GetGatewayListOfWallet(dvWalletId, 0, 100)
	if err != nil {
		log.WithError(err).Error("db/cannot get gwListOfWallet")
		return &networkserver.DvUsageModeResponse{}, err
	}

	/*dvMode, err := db.Device.GetDeviceModeByEui(req.DvEui)
	if err != nil {
		log.WithError(err).Error("db/cannot get dvMode")
		return &networkserver.DvUsageModeResponse{}, err
	}
	*/

	switch string(dvMode) {
	case networkserver.DeviceMode_name[int32(networkserver.DeviceMode_DV_INACTIVE)]:
		return &networkserver.DvUsageModeResponse{DvMode: networkserver.DeviceMode_DV_INACTIVE}, nil

	case networkserver.DeviceMode_name[int32(networkserver.DeviceMode_DV_FREE_GATEWAYS_LIMITED)]:
		for _, v := range gwList {
			fgws := networkserver.GwMac{}
			fgws.GwMac = v.Mac
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

		for _, v := range gwList {
			fgws := networkserver.GwMac{}
			fgws.GwMac = v.Mac
			resp.FreeGwMac = append(resp.FreeGwMac, &fgws)
		}

		if balance > dlPrice {
			resp.DvMode = networkserver.DeviceMode_DV_WHOLE_NETWORK
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

/*func (*M2MNetworkServerAPI) FreeGateway (ctx context.Context, req* networkserver.FreeGatewayRequest) (rep *networkserver.FreeGatewayResponse, err error) {
	log.WithFields(log.Fields{
	}).Debug("grpc_api/FreeGateway")

	freeGws, err := db.Gateway.GetFreeGwList()
	if err != nil {
		return &networkserver.FreeGatewayResponse{}, err
	}

	for _, id := range freeGws {
		freeGwId := networkserver.FreeGwId{}
		freeGwId.GwId = id
		rep.GwId = append(rep.GwId, &freeGwId)
	}

	return &networkserver.FreeGatewayResponse{GwId:rep.GwId}, nil
}*/

func (*M2MNetworkServerAPI) DlPktSent(ctx context.Context, req *networkserver.DlPktSentRequest) (*networkserver.DlPktSentResponse, error) {

	fmt.Println("-- dl packet sent req: ", req)

	var dlPkt = types.DlPkt{}
	dlPkt.Id = req.DlPkt.DlIdNs
	dlPkt.Category = types.DlCategory(req.DlPkt.Category)
	dlPkt.Nonce = req.DlPkt.Nonce

	if createAt, err := time.Parse(timeLayout, req.DlPkt.CreateAt); err != nil {
		log.WithError(err).Error("time format error")
		return &networkserver.DlPktSentResponse{}, err
	} else {
		dlPkt.SentAt = createAt
	}

	dlPkt.Size = req.DlPkt.Size
	//dlPkt.Token = req.DlPkt.TokenDlFrm1

	_, err := db.DlPacket.InsertDlPkt(dlPkt)
	if err != nil {
		log.WithError(err).Error("db/cannot update dlPkt")
		return &networkserver.DlPktSentResponse{}, err
	}

	dvId, err := db.Device.GetDeviceIdByDevEui(req.DlPkt.DevEui)
	if err != nil {
		log.WithError(err).Error("db/cannot get devID")
		return &networkserver.DlPktSentResponse{}, err
	}

	err = db.Device.UpdateDeviceLastSeen(dvId, time.Now())
	if err != nil {
		log.WithError(err).Error("db/cannot update devicelastseen")
		return &networkserver.DlPktSentResponse{}, err
	}

	gwId, err := db.Gateway.GetGatewayIdByMac(req.DlPkt.GwMac)
	if err != nil {
		log.WithError(err).Error("db/cannot get gwID")
		return &networkserver.DlPktSentResponse{}, err
	}

	dvwalletId, err := db.Device.GetWalletIdOfDevice(dvId)
	if err != nil {
		log.WithError(err).Error("db/cannot get dvWalletID")
		return &networkserver.DlPktSentResponse{}, err
	}

	gwwalletId, err := db.Gateway.GetWalletIdOfGateway(gwId)
	if err != nil {
		log.WithError(err).Error("db/cannot get gwWalletID")
		return &networkserver.DlPktSentResponse{}, err
	}

	if dvwalletId == gwwalletId {
		return &networkserver.DlPktSentResponse{}, nil
	} else {
		err = db.Wallet.TmpBalanceUpdatePktTx(dvId, gwId, dlPrice)
		if err != nil {
			log.WithError(err).Error("db/cannot update balance")
			return &networkserver.DlPktSentResponse{}, err
		}
	}

	return &networkserver.DlPktSentResponse{}, nil
}

func (*M2MNetworkServerAPI) UlPktSent(ctx context.Context, req *networkserver.UlPktSentRequest) (*networkserver.UlPktSentResponse, error) {
	ulPktCount++
	println("Pkt no: ", ulPktCount, "!")
	return nil, nil
}
