package db

import (
	pg "github.com/mxc-foundation/mxprotocol-server/m2m/db/postgres_db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/types"
)

type dlPacketDBInterface interface {
	CreateDlPktTable() error
	InsertDlPkt(dlPkt types.DlPkt) (insertIndex int64, err error)
	GetAggDlPktDeviceWallet(startIndDlPkt, endIndDlPkt int64) (walletId []int64, count []int64, err error)
	GetAggDlPktGatewayWallet(startIndDlPkt, endIndDlPkt int64) (walletId []int64, count []int64, err error)
	GetAggDlPktFreeWallet(startIndDlPkt, endIndDlPkt int64) (walletId []int64, count []int64, err error)
	GetLastReceivedDlPktId() (latestId int64, err error)
}

var DlPacket = dlPacketDBInterface(&pg.PgDlPacket)
