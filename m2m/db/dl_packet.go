package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type dlPacketDBInterface interface {
	CreateDlPktTable() error
	InsertDlPkt(dlPkt types.DlPkt) (insertIndex int64, err error)
}

var dlPacket dlPacketDBInterface

func DbCreateDlPktTable() error {
	dlPacket = &pg.PgDlPacket
	return dlPacket.CreateDlPktTable()
}

func DbInsertDlPkt(dlp types.DlPkt) (insertIndex int64, err error) {
	return dlPacket.InsertDlPkt(dlp)
}
