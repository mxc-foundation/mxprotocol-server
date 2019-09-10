package db

import (
	"time"

	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	types "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type gatewayDBInterface interface {
	CreateGatewayTable() error
	InsertGateway(gw types.Gateway) (insertIndex int64, err error)
	GetGatewayMode(gwId int64) (gwMode types.GatewayMode, err error)
	SetGatewayMode(gwId int64, gwMode types.GatewayMode) (err error)
	GetGatewayIdByMac(mac string) (gwId int64, err error)
	UpdateGatewayLastSeen(gwId int64, newTime time.Time) (err error)
	GetGatewayProfile(gwId int64) (gw types.Gateway, err error)
	GetGatewayListOfWallet(walletId int64, offset int64, limit int64) (gwList []types.Gateway, err error)
	GetGatewayRecCnt(walletId int64) (recCnt int64, err error)
}
var gateway gatewayDBInterface

func DbCreateGatewayTable() error {
	gateway = &pg.PgGateway
	return gateway.CreateGatewayTable()
}

func DbInsertGateway(gw types.Gateway) (insertIndex int64, err error) {
	return gateway.InsertGateway(gw)
}

func DbGetGatewayMode(gwId int64) (gwMode types.GatewayMode, err error) {
	return gateway.GetGatewayMode(gwId)
}

func DbSetGatewayMode(gwId int64, gwMode types.GatewayMode) (err error) {
	return gateway.SetGatewayMode(gwId, gwMode)
}

func DbDeleteGateway(gwId int64) (err error) {
	return DbSetGatewayMode(gwId, types.GW_DELETED)
}

func DbGetGatewayIdByMac(mac string) (gwId int64, err error) {
	return gateway.GetGatewayIdByMac(mac)
}

func DbUpdateGatewayLastSeen(gwId int64, newTime time.Time) (err error) {
	return gateway.UpdateGatewayLastSeen(gwId, newTime)
}

func DbGetGatewayProfile(gwId int64) (gw types.Gateway, err error) {
	return gateway.GetGatewayProfile(gwId)
}

func DbGetGatewayListOfWallet(walletId int64, offset int64, limit int64) (gwList []types.Gateway, err error) {
	return gateway.GetGatewayListOfWallet(walletId, offset, limit)
}

func DbGetGatewayRecCnt(walletId int64) (recCnt int64, err error) {
	return gateway.GetGatewayRecCnt(walletId)
}
