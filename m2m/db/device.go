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

// func DbGetWalletIdFromOrgId(orgIdLora int64) (int64, error) {
// 	return pg.PgDB.GetWalletIdFromOrgId(orgIdLora)
// }
