package db

import (
	"time"

	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type withdrawDBInterface interface {
	CreateWithdrawTable() error
	UpdateWithdrawSuccessful(withdrawId int64, txHash string, txApprovedTime time.Time) error
	CreateWithdrawFunctions() error
	InitWithdrawReq(walletId int64, value float64, extCurrencyAbbr string) (withdrawId int64, err error)
	UpdateWithdrawPaymentQueryId(withdrawId int64, reqIdPaymentServ int64) error
	GetWithdrawHist(walletId int64, offset int64, limit int64) ([]pg.WithdrawHistRet, error)
	GetWithdrawHistRecCnt(walletId int64) (recCnt int64, err error)
}

var withdraw withdrawDBInterface

type WithdrawHistRet pg.WithdrawHistRet

func dbCreateWithdrawTable() error {
	withdraw = &pg.PgWithdraw
	return withdraw.CreateWithdrawTable()
}

func dbCreateWithdrawRelations() error {
	return withdraw.CreateWithdrawFunctions()
}

func DbUpdateWithdrawSuccessful(withdrawId int64, txHash string, txApprovedTime time.Time) error {
	return withdraw.UpdateWithdrawSuccessful(withdrawId, txHash, txApprovedTime)
}

func DbInitWithdrawReq(walletId int64, amount float64, extCurAbv string) (withdrawId int64, err error) {
	return withdraw.InitWithdrawReq(walletId, amount, extCurAbv)
}

func DbUpdateWithdrawPaymentQueryId(withdrawId int64, reqIdPaymentServ int64) error {
	return withdraw.UpdateWithdrawPaymentQueryId(withdrawId, reqIdPaymentServ)
}

func castWithdrawHistRet(acntHist []pg.WithdrawHistRet, err1 error) (castedVal []WithdrawHistRet, err error) {
	for _, v := range acntHist {
		castedVal = append(castedVal, WithdrawHistRet(v))
	}
	return castedVal, err1
}

func DbGetWithdrawHist(walletId int64, offset int64, limit int64) ([]WithdrawHistRet, error) {
	return castWithdrawHistRet(withdraw.GetWithdrawHist(walletId, offset, limit))
}

func DbGetWithdrawHistRecCnt(walletId int64) (int64, error) {
	return withdraw.GetWithdrawHistRecCnt(walletId)
}
