package db

import (
	"time"

	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type deviceDBInterface interface {
	CreateDeviceTable() error
	InsertDevice(dv types.Device) (insertIndex int64, err error)
	GetDeviceMode(dvId int64) (dvMode types.DeviceMode, err error)
	SetDeviceMode(dvId int64, dvMode types.DeviceMode) (err error)
	GetDeviceIdByDevEui(devEui string) (devId int64, err error)
	UpdateDeviceLastSeen(dvId int64, newTime time.Time) (err error)
	GetDeviceProfile(dvId int64) (dv types.Device, err error)
	GetDeviceListOfWallet(walletId int64, offset int64, limit int64) (dvList []types.Device, err error)
	GetDeviceRecCnt(walletId int64) (recCnt int64, err error)
}
var device deviceDBInterface

func DbCreateDeviceTable() error {
	device = &pg.PgDevice
	return device.CreateDeviceTable()
}

func DbInsertDevice(dv types.Device) (insertIndex int64, err error) {
	return device.InsertDevice(dv)
}

func DbGetDeviceMode(dvId int64) (dvMode types.DeviceMode, err error) {
	return device.GetDeviceMode(dvId)
}

func DbSetDeviceMode(dvId int64, dvMode types.DeviceMode) (err error) {
	return device.SetDeviceMode(dvId, dvMode)
}

func DbDeleteDevice(dvId int64) (err error) {
	return DbSetDeviceMode(dvId, types.DV_DELETED)
}

func DbGetDeviceIdByDevEui(devEui string) (devId int64, err error) {
	return device.GetDeviceIdByDevEui(devEui)
}

func DbUpdateDeviceLastSeen(dvId int64, newTime time.Time) (err error) {
	return device.UpdateDeviceLastSeen(dvId, newTime)
}

func DbGetDeviceProfile(dvId int64) (types.Device, error) {
	return device.GetDeviceProfile(dvId)
}

func DbGetDeviceListOfWallet(walletId int64, offset int64, limit int64) (dvList []types.Device, err error) {
	return device.GetDeviceListOfWallet(walletId, offset, limit)
}

func DbGetDeviceRecCnt(walletId int64) (recCnt int64, err error) {
	return device.GetDeviceRecCnt(walletId)
}
