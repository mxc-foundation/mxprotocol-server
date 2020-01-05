package db

import (
	pg "github.com/mxc-foundation/mxprotocol-server/m2m/db/postgres_db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/types"
)

type topupDBInterface interface {
	CreateTopupTable() error
	CreateTopupFunctions() error
	AddTopUpRequest(acntAdrSender string, acntAdrRcvr string, txHash string, value float64, extCurAbv string) (topupId int64, err error)
	GetTopupHist(walletId int64, offset int64, limit int64) ([]types.TopupHistRet, error)
	GetTopupHistRecCnt(walletId int64) (recCnt int64, err error)
}

var Topup = topupDBInterface(&pg.PgTopup)
