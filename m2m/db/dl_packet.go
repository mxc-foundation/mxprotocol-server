package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type dlPacketDBInterface interface {
	CreateDlPktTable() error
	InsertDlPkt(dlPkt types.DlPkt) (insertIndex int64, err error)
	GetAggDlPktDeviceWallet(startIndDlPkt, endIndDlPkt int64) (walletId []int64, count []int64, err error)
	GetAggDlPktGatewayWallet(startIndDlPkt, endIndDlPkt int64) (walletId []int64, count []int64, err error)
	GetAggDlPktFreeWallet(startIndDlPkt, endIndDlPkt int64) (walletId []int64, count []int64, err error)
	GetLastReceviedDlPktId() (latestId int64, err error)
}

var DlPacket = dlPacketDBInterface(&pg.PgDlPacket)
