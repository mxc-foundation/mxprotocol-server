package m2m_networkserver

import (
	"fmt"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/networkserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"testing"
	log "github.com/sirupsen/logrus"
)

/*func tt(){
	t1 := time.Now() // get current time
	TestDvUsageMode("98749874640c0bf2")

	//fmt.Println("Resp Here: ", resp)
	//elapsed := time.Since(t1)
	//fmt.Println("Spent: ", elapsed)
}*/

var dvEui = "98749874640c0bf2"

func TestDvUsageMode(t *testing.T) {
	log.WithFields(log.Fields{
		"dvId": t,
	}).Debug("grpc_api/DvUsageMode")

	dvWalletId, dvMode, err := db.Device.GetDevWalletIdAndModeByEui(dvEui)
	if err != nil {
		log.WithError(err).Error("db/cannot get dvWalletId and dvMode")
		fmt.Println(err)
	}

	switch string(dvMode) {
	case networkserver.DeviceMode_name[int32(networkserver.DeviceMode_DV_INACTIVE)]:
		fmt.Println("DeviceMode_DV_INACTIVE")

	case networkserver.DeviceMode_name[int32(networkserver.DeviceMode_DV_FREE_GATEWAYS_LIMITED)]:
		_, gwMac, err := db.Gateway.GetFreeGwList(dvWalletId)
		if err != nil {
			log.WithError(err).Error("db/cannot get gwListOfWallet")
			fmt.Println(err)
		}

		for _, v := range gwMac {
			fgws := networkserver.GwMac{}
			fgws.GwMac = v
			//resp = append(resp, &fgws)
			fmt.Println(fgws.GwMac)
		}

		resp := "DeviceMode_DV_FREE_GATEWAYS_LIMITED"

		fmt.Println(resp)

	case networkserver.DeviceMode_name[int32(networkserver.DeviceMode_DV_WHOLE_NETWORK)]:
		walletId, err := db.Device.GetDevWalletIdByEui(dvEui)
		if err != nil {
			log.WithError(err).Error("db/cannot get walletId")
			fmt.Println(err)
		}

		balance, err := db.Wallet.GetWalletBalance(walletId)
		if err != nil {
			log.WithError(err).Error("db/cannot get balance")
			fmt.Println(err)
		}

		_, gwMac, err := db.Gateway.GetFreeGwList(dvWalletId)
		if err != nil {
			log.WithError(err).Error("db/cannot get gwListOfWallet")
			fmt.Println(err)
		}

		for _, v := range gwMac {
			fgws := networkserver.GwMac{}
			fgws.GwMac = v
			//resp.FreeGwMac = append(resp.FreeGwMac, &fgws)
			fmt.Println(fgws.GwMac)
		}

		resp := "networkserver.DeviceMode_DV_WHOLE_NETWORK"
		if balance < dlPrice {
			//resp.EnoughBalance = false
			fmt.Println("EnoughBalance")
		} else {
			//resp.EnoughBalance = true
			fmt.Println("Not Enough Balance")
		}
		fmt.Println(resp)

	case networkserver.DeviceMode_name[int32(networkserver.DeviceMode_DV_DELETED)]:
		fmt.Println("DeviceMode_DV_DELETED")
	}

	fmt.Println("Not in the case!")
}
