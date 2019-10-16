package device

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/appserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/api/clients/appserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
	"time"
)

var timer *time.Timer

func Setup() error {
	log.Info("Syncronize devices from appserver")
	timer = time.AfterFunc(1 * time.Second, syncDevicesFromAppserverByBatch)
	return nil
}

func syncDevicesFromAppserverByBatch() {
	// get device list from local database
	localDeviceList, err := db.Device.GetAllDevices()
	if err != nil {
		// reset timer
		timer.Reset(10 * time.Second)
		return
	}
	log.Debug("syncDevicesFromAppserverByBatch_local: count=", len(localDeviceList))

	// get device list from appserver
	client, err := appserver.GetPool().Get(config.Cstruct.AppServer.Server, []byte(config.Cstruct.AppServer.CACert),
		[]byte(config.Cstruct.AppServer.TLSCert), []byte(config.Cstruct.AppServer.TLSKey))
	if err != nil {
		// reset timer
		timer.Reset(10 * time.Second)
		return
	}

	devEuiList, err := client.GetDeviceDevEuiList(context.Background(), &empty.Empty{})
	if err != nil {
		// reset timer
		timer.Reset(10 * time.Second)
		return
	}

	log.Debug("syncDevicesFromAppserverByBatch_appserver: count=", len(devEuiList.DevEui), " list=", devEuiList.DevEui)

	// do synchronization
	for _, localDev := range localDeviceDevEuiList {
		for _, appDev := range devEuiList.DevEui {
			id, err := db.Device.GetDeviceIdByDevEui(localDev)
		}
	}

	return
}

func SyncDeviceProfileByDevEuiFromAppserver(devId int64, devEui string) error {
	client, err := appserver.GetPool().Get(config.Cstruct.AppServer.Server, []byte(config.Cstruct.AppServer.CACert),
		[]byte(config.Cstruct.AppServer.TLSCert), []byte(config.Cstruct.AppServer.TLSKey))
	if err != nil {
		return err
	}

	device, err := client.GetDeviceByDevEui(context.Background(), &api.GetDeviceByDevEuiRequest{DevEui: devEui})
	if err == nil && device.DevProfile == nil {
		// device no longer exist, delete from database
		err := db.Device.SetDeviceMode(devId, types.DV_DELETED)
		if err != nil {
			log.WithError(err).Warn("device/SyncDeviceProfileByDevEuiFromAppserver: devId", devId)
		}
	} else if err == nil {
		// get device successfully, add/update device
		walletId, err := db.Wallet.GetWalletIdFromOrgId(device.OrgId)
		if err != nil {
			log.WithError(err).Error("device/SyncDeviceProfileByDevEuiFromAppserver: device is not linked to any wallet")
			err := db.Device.SetDeviceMode(devId, types.DV_DELETED)
			if err != nil {
				log.WithError(err).Warn("device/SyncDeviceProfileByDevEuiFromAppserver: devId", devId)
			}
			// in this case, it is not necessary to retry again
			return nil
		}

		_, err = db.Device.InsertDevice(types.Device{
			DevEui:        device.DevProfile.DevEui,
			FkWallet:      walletId,
			Mode:          types.DV_WHOLE_NETWORK,
			CreatedAt:     time.Now(),
			ApplicationId: device.DevProfile.ApplicationId,
			Name:          device.DevProfile.Name,
		})
		if err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}
