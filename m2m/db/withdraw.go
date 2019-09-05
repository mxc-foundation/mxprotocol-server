package db

import (
	"time"

	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type WithdrawHistRet pg.WithdrawHistRet

func dbCreateWithdrawTable() error {
	return pg.CreateWithdrawTable()
}

func dbCreateWithdrawRelations() error {
	return pg.CreateWithdrawFunctions()
}

func DbGetWalletIdByActiveAcnt(acntAdr string, externalCur string) (walletId int64, err error) {
	return pg.GetWalletIdofActiveAcnt(acntAdr, externalCur)
}

func DbUpdateWithdrawSuccessful(withdrawId int64, txHash string, txApprovedTime time.Time) error {
	return pg.UpdateWithdrawSuccessful(withdrawId, txHash, txApprovedTime)
}

func DbInitWithdrawReq(walletId int64, amount float64, extCurAbv string) (withdrawId int64, err error) {
	return pg.InitWithdrawReq(walletId, amount, extCurAbv)
}

func DbUpdateWithdrawPaymentQueryId(withdrawId int64, reqIdPaymentServ int64) error {
	return pg.UpdateWithdrawPaymentQueryId(withdrawId, reqIdPaymentServ)
}

func castWithdrawHistRet(acntHist []pg.WithdrawHistRet, err1 error) (castedVal []WithdrawHistRet, err error) {
	for _, v := range acntHist {
		castedVal = append(castedVal, WithdrawHistRet(v))
	}
	return castedVal, err1
}

func DbGetWithdrawHist(walletId int64, offset int64, limit int64) ([]WithdrawHistRet, error) {
	return castWithdrawHistRet(pg.GetWithdrawHist(walletId, offset, limit))
}

func DbGetWithdrawHistRecCnt(walletId int64) (int64, error) {
	return pg.GetWithdrawHistRecCnt(walletId)
}
