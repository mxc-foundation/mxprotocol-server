package db

import (
	"time"
)

func dbCreateWithdrawTable() error {
	return db.CreateWithdrawTable()
}

func dbCreateWithdrawRelations() error {
	return db.CreateWithdrawFunctions()
}

func DbGetWalletIdByActiveAcnt(acntAdr string, externalCur string) (walletId int64, err error) {
	return db.GetWalletIdofActiveAcnt(acntAdr, externalCur)
}

func DbUpdateWithdrawSuccessful(withdrawId int64, txHash string, txApprovedTime time.Time) error {
	return db.UpdateWithdrawSuccessful(withdrawId, txHash, txApprovedTime)
}

func DbInitWithdrawReq(walletId int64, amount float64, extCurAbv string) (withdrawId int64, err error) {
	return db.InitWithdrawReq(walletId, amount, extCurAbv)
}

func DbUpdateWithdrawPaymentQueryId(withdrawId int64, reqIdPaymentServ int64) error {
	return db.UpdateWithdrawPaymentQueryId(withdrawId, reqIdPaymentServ)
}

func DbGetWithdrawHist(walletId int64, offset int64, limit int64) ([]WithdrawHistRet, error) {
	return db.GetWithdrawHist(walletId, offset, limit)
}
