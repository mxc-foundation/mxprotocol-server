package db

import (
	"time"

	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type DeviceMode string

const (
	DV_INACTIVE              DeviceMode = "INACTIVE"
	DV_FREE_GATEWAYS_LIMITED DeviceMode = "FREE_GATEWAYS_LIMITED"
	DV_WHOLE_NETWORK         DeviceMode = "WHOLE_NETWORK"
	DV_DELETED               DeviceMode = "DELETED"
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

func DbGetDeviceMode(dvId int64) (dvMode string, err error) {
	return pg.PgDB.GetDeviceMode(dvId)
}

func DbSetDeviceMode(dvId int64, dvMode DeviceMode) (err error) {
	return pg.PgDB.SetDeviceMode(dvId, string(dvMode))
}

func DbDeleteDevice(dvId int64) (err error) {
	return DbSetDeviceMode(dvId, DV_DELETED)
}

func DbGetDeviceIdByDevEui(devEui string) (devId int64, err error) {
	return pg.PgDB.GetDeviceIdByDevEui(devEui)
}

func DbUpdateDeviceLastSeen(dvId int64, newTime time.Time) (err error) {
	return pg.PgDB.UpdateDeviceLastSeen(dvId, newTime)
}

func castDeviceList(dvList []pg.Device, err error) ([]Device, error) {
	var castedVal []Device
	for _, v := range dvList {
		castedVal = append(castedVal, Device(v))
	}
	return castedVal, err
}

func DbGetDeviceProfile(dvId int64) (Device, error) {
	prf, err := pg.PgDB.GetDeviceProfile(dvId)
	return Device(prf), err
}

func DbGetDeviceListOfWallet(walletId int64, offset int64, limit int64) (dvList []Device, err error) {
	return castDeviceList(pg.PgDB.GetDeviceListOfWallet(walletId, offset, limit))
}

func DbGetDeviceRecCnt(walletId int64) (recCnt int64, err error) {
	return pg.PgDB.GetDeviceRecCnt(walletId)
}
