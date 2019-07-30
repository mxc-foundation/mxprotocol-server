package db

import (
	"time"
)

func dbCreateWithdrawTable() error {
	return dbHandler.CreateWithdrawTable()
}

func dbCreateWithdrawRelations() error {
	return dbHandler.CreateWithdrawFunctions()
}

func DbGetWalletIdByActiveAcnt(acntAdr string, externalCur string) (walletId int64, err error) {
	return dbHandler.GetWalletIdofActiveAcnt(acntAdr, externalCur)
}

func DbUpdateWithdrawSuccessful(withdrawId int64, txHash string, txApprovedTime time.Time) error {
	return dbHandler.UpdateWithdrawSuccessful(withdrawId, txHash, txApprovedTime)
}

func DbInitWithdrawReq(walletId int64, amount float64, extCurAbv string) (withdrawId int64, err error) {
	return dbHandler.InitWithdrawReq(walletId, amount, extCurAbv)
}

func DbUpdateWithdrawPaymentQueryId(withdrawId int64, reqIdPaymentServ int64) error {
	return dbHandler.UpdateWithdrawPaymentQueryId(withdrawId, reqIdPaymentServ)
}

func DbGetWithdrawHist(walletId int64, offset int64, limit int64) ([]WithdrawHistRet, error) {
	return dbHandler.GetWithdrawHist(walletId, offset, limit)
}
