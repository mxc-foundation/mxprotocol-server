package db

import (
	"time"

	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	types "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

func DbCreateDeviceTable() error {
	return pg.PgDB.CreateDeviceTable()
}

func DbInsertDevice(dv types.Device) (insertIndex int64, err error) {
	return pg.PgDB.InsertDevice(dv)
}

func DbGetDeviceMode(dvId int64) (dvMode types.DeviceMode, err error) {
	return pg.PgDB.GetDeviceMode(dvId)
}

func DbSetDeviceMode(dvId int64, dvMode types.DeviceMode) (err error) {
	return pg.PgDB.SetDeviceMode(dvId, dvMode)
}

func DbDeleteDevice(dvId int64) (err error) {
	return DbSetDeviceMode(dvId, types.DV_DELETED)
}

func DbGetDeviceIdByDevEui(devEui string) (devId int64, err error) {
	return pg.PgDB.GetDeviceIdByDevEui(devEui)
}

func DbUpdateDeviceLastSeen(dvId int64, newTime time.Time) (err error) {
	return pg.PgDB.UpdateDeviceLastSeen(dvId, newTime)
}

func DbGetDeviceProfile(dvId int64) (types.Device, error) {
	return pg.PgDB.GetDeviceProfile(dvId)
}

func DbGetDeviceListOfWallet(walletId int64, offset int64, limit int64) (dvList []types.Device, err error) {
	return pg.PgDB.GetDeviceListOfWallet(walletId, offset, limit)
}

func DbGetDeviceRecCnt(walletId int64) (recCnt int64, err error) {
	return pg.PgDB.GetDeviceRecCnt(walletId)
}
