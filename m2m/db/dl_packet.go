package db

import (
	"time"

	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type DlCategory string

const (
	JOIN_ANS    DlCategory = "JOIN_ANS"
	MAC_COMMAND DlCategory = "MAC_COMMAND"
	PAYLOAD     DlCategory = "PAYLOAD"
	UNKNOWN     DlCategory = "UNKNOWN"
)

type DlPkt pg.DlPkt

func DbCreateDlPktTable() error {
	return pg.PgDB.CreateDlPktTable()
}

func DbInsertDlPkt(dvId int64, gwId int64, nonce int64, sentTime time.Time, size float64, category DlCategory) (insertIndex int64, err error) {
	dlp := pg.DlPkt{
		FkDevice:  dvId,
		FkGateway: gwId,
		Nonce:     nonce,
		SentAt:    sentTime,
		Size:      size,
		Category:  string(category),
	}
	return pg.PgDB.InsertDlPkt(dlp)
}
