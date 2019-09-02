package db

import (
	"time"

	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type DeviceMode string

const (
	INACTIVE              DeviceMode = "INACTIVE"
	FREE_GATEWAYS_LIMITED DeviceMode = "FREE_GATEWAYS_LIMITED"
	WHOLE_NETWORK         DeviceMode = "WHOLE_NETWORK"
	DELETED               DeviceMode = "DELETED"
)

type Device pg.Device

func DbCreateDeviceTable() error {
	return pg.PgDB.CreateDeviceTable()
}

func DbInsertDevice(devEui string, fkWallet int64, mode DeviceMode, appId int64, name string) (insertIndex int64, err error) {
	dv := pg.Device{
		DevEui:        devEui,
		FkWallet:      fkWallet,
		Mode:          string(mode),
		CreatedAt:     time.Now().UTC(),
		ApplicationId: appId,
		Name:          name,
	}
	return pg.PgDB.InsertDevice(dv)
}

func DbGetDeviceListByWallet(walletId int64) (dvList []Device, err error) {
	return nil, nil
}

func DbGetDeviceProfile(dvId int64) (dv Device, err error) {
	return Device{}, nil
}

func DbGetDeviceMode(dvId int64) (dvMode DeviceMode, err error) {
	return INACTIVE, nil
}

func DbSetDeviceMode(dvId int64, dvMode DeviceMode) (err error) {
	return nil
}

func DbDeletDevice(dvId int64) (err error) {
	return DbSetDeviceMode(dvId, DELETED)
}

func DbGetDeviceIdByDevEui(devEui string) (devId int64, err error) {
	return 1, nil
}

func DbUpdateDeviceLastSeen(newTime time.Time, err error) {

}
