package device

import (
	"context"
	log "github.com/sirupsen/logrus"
	api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/appserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/api/clients/appserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
	"time"
)

func Setup() error {
	log.Info("Setup device service")
	return nil
}

func SyncDeviceProfileFromAppserver(dev types.Device, oper types.OperationType) error {
	client, err := appserver.GetPool().Get(config.Cstruct.AppServer.Server, []byte(config.Cstruct.AppServer.CACert),
		[]byte(config.Cstruct.AppServer.TLSCert), []byte(config.Cstruct.AppServer.TLSKey))
	if err != nil {
		return err
	}

	device, err := client.GetDeviceByDevEui(context.Background(), &api.GetDeviceByDevEuiRequest{DevEui: dev.DevEui})
	if err == nil && device.DevProfile == nil {
		// device no longer exist, delete from database
		err := db.Device.SetDeviceMode(dev.Id, "DELETED")
		if err != nil {
			log.WithError(err).Warn("device/SyncDeviceProfileFromAppserver")
		}
	} else if err == nil {
		// get device successfully, add/update device
		walletId, err := db.Wallet.GetWalletIdFromOrgId(device.OrgId)
		if err != nil {
			log.WithError(err).Error("device/SyncDeviceProfileFromAppserver: device is not linked to any wallet")
			err := db.Device.SetDeviceMode(dev.Id, "DELETED")
			if err != nil {
				log.WithError(err).Warn("device/SyncDeviceProfileFromAppserver")
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
	} else if err != nil {
		return err
	}

	return nil
}
