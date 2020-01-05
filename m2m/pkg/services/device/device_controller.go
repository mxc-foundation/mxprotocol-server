package device

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	api "github.com/mxc-foundation/mxprotocol-server/m2m/api/appserver"
	"github.com/mxc-foundation/mxprotocol-server/m2m/db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/api/clients/appserver"
	"github.com/mxc-foundation/mxprotocol-server/m2m/pkg/config"
	"github.com/mxc-foundation/mxprotocol-server/m2m/types"
	"time"
)

var timer *time.Timer

func Setup() error {
	log.Info("Syncronize devices from appserver")
	timer = time.AfterFunc(1*time.Second, syncDevicesFromAppserverByBatch)

	// give it time to sync before whole service starts
	time.Sleep(5 * time.Second)
	return nil
}

func syncDevicesFromAppserverByBatch() {
	// get device list from local database
	localDeviceList, err := db.Device.GetAllDevices()
	if err != nil {
		log.WithError(err).Error("service/device/syncDevicesFromAppserverByBatch")
		// reset timer
		timer.Reset(10 * time.Second)
		return
	}
	//log.Debug("syncDevicesFromAppserverByBatch_local: count=", len(localDeviceList))

	// get device list from appserver
	client, err := appserver.GetPool().Get(config.Cstruct.AppServer.Server, []byte(config.Cstruct.AppServer.CACert),
		[]byte(config.Cstruct.AppServer.TLSCert), []byte(config.Cstruct.AppServer.TLSKey))
	if err != nil {
		log.WithError(err).Error("service/device/syncDevicesFromAppserverByBatch")
		// reset timer
		timer.Reset(10 * time.Second)
		return
	}

	devEuiList, err := client.GetDeviceDevEuiList(context.Background(), &empty.Empty{})
	if err != nil {
		log.WithError(err).Error("service/device/syncDevicesFromAppserverByBatch")
		// reset timer
		timer.Reset(10 * time.Second)
		return
	}

	//log.Debug("syncDevicesFromAppserverByBatch_appserver: count=", len(devEuiList.DevEui), " list=", devEuiList.DevEui)

	// if len(localDeviceList) == 0, len(devEuiList.DevEui) == 0, just return
	if len(localDeviceList) == 0 && len(devEuiList.DevEui) == 0 {
		return
	}

	// if len(localDeviceList) == 0, len(devEuiList.DevEui) != 0, just insert new device
	if len(localDeviceList) == 0 && len(devEuiList.DevEui) != 0 {
		for _, v := range devEuiList.DevEui {
			device, err := getDeviceFromAppserver(v)
			if err != nil {
				log.WithError(err).Error("service/device/syncDevicesFromAppserverByBatch")
				// reset timer
				timer.Reset(10 * time.Second)
				return
			}

			_, err = db.Device.InsertDevice(device)
			if err != nil {
				log.WithError(err).Error("service/device/syncDevicesFromAppserverByBatch")
				timer.Reset(10 * time.Second)
				return
			}
		}

		return
	}

	// if len(localDeviceList) != 0, len(devEuiList.DevEui) == 0, just delete all devices
	if len(localDeviceList) != 0 && len(devEuiList.DevEui) == 0 {
		for _, v := range localDeviceList {
			if err := db.Device.SetDeviceMode(v.Id, types.DV_DELETED); err != nil {
				log.WithError(err).Error("service/device/syncDevicesFromAppserverByBatch")
				timer.Reset(10 * time.Second)
				return
			}
		}

		return
	}

	// if len(localDeviceList) != 0, len(devEuiList) != 0, compare and synchronize
	if len(localDeviceList) != 0 && len(devEuiList.DevEui) != 0 {
		type syncDevice struct {
			device             types.Device
			existInAppserver   bool
			existInLocalServer bool
		}
		syncDeviceList := make(map[string]syncDevice)

		for _, localDevIter := range localDeviceList {
			dev := syncDevice{device: localDevIter, existInAppserver: false, existInLocalServer: true}
			syncDeviceList[dev.device.DevEui] = dev
		}

		for _, appDevIter := range devEuiList.DevEui {
			if val, ok := syncDeviceList[appDevIter]; ok {
				val.existInAppserver = true
			} else {
				newDev := syncDevice{device: types.Device{DevEui: appDevIter}, existInAppserver: true, existInLocalServer: false}
				syncDeviceList[newDev.device.DevEui] = newDev
			}
		}

		// process syncDeviceList
		for k, v := range syncDeviceList {
			// synchronize devices
			// v.existInLocalServer == true && v.existInAppserver == false, delete device locally
			// v.existInLocalServer == false && v.existInAppserver == true, insert new device
			// v.existInLocalServer == true && v.existInAppserver == true, do nothing, continue loop
			// v.existInLocalServer == false && v.existInAppserver == false, this option does not exist

			if v.existInLocalServer == true && v.existInAppserver == false {
				// delete device locally
				if err := db.Device.SetDeviceMode(v.device.Id, types.DV_DELETED); err != nil {
					log.WithError(err).Error("service/device/syncDevicesFromAppserverByBatch")
					timer.Reset(10 * time.Second)
					return
				}
			}

			if v.existInLocalServer == false && v.existInAppserver == true {
				// insert new device
				device, err := getDeviceFromAppserver(k)
				if err != nil {
					log.WithError(err).Error("service/device/syncDevicesFromAppserverByBatch")
					timer.Reset(10 * time.Second)
					return
				}

				_, err = db.Device.InsertDevice(device)
				if err != nil {
					log.WithError(err).Error("service/device/syncDevicesFromAppserverByBatch")
					timer.Reset(10 * time.Second)
					return
				}

			}

			if v.existInLocalServer == true && v.existInAppserver == true {
				// do nothing
				continue
			}

		}

		return
	}
}

func getDeviceFromAppserver(devEui string) (types.Device, error) {
	device := types.Device{}
	appserverClient, err := appserver.GetPool().Get(config.Cstruct.AppServer.Server, []byte(config.Cstruct.AppServer.CACert),
		[]byte(config.Cstruct.AppServer.TLSCert), []byte(config.Cstruct.AppServer.TLSKey))
	if err != nil {
		return device, err
	}

	resp, err := appserverClient.GetDeviceByDevEui(context.Background(), &api.GetDeviceByDevEuiRequest{DevEui: devEui})
	if err != nil {
		return device, err
	}

	walletId, err := db.Wallet.GetWalletIdFromOrgId(resp.OrgId)
	if err != nil {
		return device, err
	}

	createdTimeUpdated, _ := ptypes.Timestamp(resp.DevProfile.CreatedAt)
	device.DevEui = devEui
	device.Mode = types.DV_WHOLE_NETWORK
	device.Name = resp.DevProfile.Name
	device.FkWallet = walletId
	device.CreatedAt = createdTimeUpdated
	device.ApplicationId = resp.DevProfile.ApplicationId

	return device, nil
}

/*
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
}*/
