package db

import (
	"time"

	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type gatewayDBInterface interface {
	CreateGatewayTable() error
	InsertGateway(gw types.Gateway) (insertIndex int64, err error)
	GetGatewayMode(gwId int64) (gwMode types.GatewayMode, err error)
	SetGatewayMode(gwId int64, gwMode types.GatewayMode) (err error)
	GetGatewayIdByMac(mac string) (gwId int64, err error)
	UpdateGatewayLastSeen(gwId int64, newTime time.Time) (err error)
	GetGatewayProfile(gwId int64) (gw types.Gateway, err error)
	GetGatewayListOfWallet(orgId int64, offset int64, limit int64) (gwList []types.Gateway, err error)
	GetGatewayRecCnt(walletId int64) (recCnt int64, err error)
}

var Gateway = gatewayDBInterface(&pg.PgGateway)