package db

import (
	pg "github.com/mxc-foundation/mxprotocol-server/m2m/db/postgres_db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/types"
	"time"
)

type withdrawDBInterface interface {
	CreateWithdrawTable() error
	UpdateWithdrawSuccessful(withdrawId int64, txHash string, txApprovedTime time.Time) error
	CreateWithdrawFunctions() error
	InitWithdrawReq(walletId int64, value float64, extCurrencyAbbr string) (withdrawId int64, err error)
	UpdateWithdrawPaymentQueryId(withdrawId int64, reqIdPaymentServ int64) error
	GetWithdrawHist(walletId int64, offset int64, limit int64) ([]types.WithdrawHistRet, error)
	GetWithdrawHistRecCnt(walletId int64) (recCnt int64, err error)
}

var Withdraw = withdrawDBInterface(&pg.PgWithdraw)
