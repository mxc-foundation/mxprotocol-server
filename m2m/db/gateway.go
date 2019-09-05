package db

import (
	"time"

	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	types "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

func DbCreateGatewayTable() error {
	return pg.CreateGatewayTable()
}

func DbInsertGateway(gw types.Gateway) (insertIndex int64, err error) {
	return pg.InsertGateway(gw)
}

func DbGetGatewayMode(gwId int64) (gwMode types.GatewayMode, err error) {
	return pg.GetGatewayMode(gwId)
}

func DbSetGatewayMode(gwId int64, gwMode types.GatewayMode) (err error) {
	return pg.SetGatewayMode(gwId, gwMode)
}

func DbDeleteGateway(gwId int64) (err error) {
	return DbSetGatewayMode(gwId, types.GW_DELETED)
}

func DbGetGatewayIdByMac(mac string) (gwId int64, err error) {
	return pg.GetGatewayIdByMac(mac)
}

func DbUpdateGatewayLastSeen(gwId int64, newTime time.Time) (err error) {
	return pg.UpdateGatewayLastSeen(gwId, newTime)
}

func DbGetGatewayProfile(gwId int64) (gw types.Gateway, err error) {
	return pg.GetGatewayProfile(gwId)
}

func DbGetGatewayListOfWallet(walletId int64, offset int64, limit int64) (gwList []types.Gateway, err error) {
	return pg.GetGatewayListOfWallet(walletId, offset, limit)
}

func DbGetGatewayRecCnt(walletId int64) (recCnt int64, err error) {
	return pg.GetGatewayRecCnt(walletId)
}
