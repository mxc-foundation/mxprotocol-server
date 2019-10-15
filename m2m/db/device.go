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
	DeleteDevice(dvId int64) (err error)
	GetDeviceIdByDevEui(devEui string) (devId int64, err error)
	UpdateDeviceLastSeen(dvId int64, newTime time.Time) (err error)
	GetDeviceProfile(dvId int64) (dv types.Device, err error)
	GetDeviceListOfWallet(walletId int64, offset int64, limit int64) (dvList []types.Device, err error)
	GetAllDeviceDevEui() (devEuiLIst []string, err error)
	GetDeviceRecCnt(walletId int64) (recCnt int64, err error)
	GetWalletIdOfDevice(dvId int64) (dvWalletId int64, err error)
	GetDevWalletIdByEui(devEui string) (walletId int64, err error)
	GetDevWalletIdAndModeByEui(devEui string) (dvWalletId int64, dvMode types.DeviceMode, err error)
}

var Device = deviceDBInterface(&pg.PgDevice)
