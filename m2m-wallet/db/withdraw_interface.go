package db

import (
	"time"
)

func dbCreateWithdrawTable() error {
	return pgDb.CreateWithdrawTable()
}

func dbCreateWithdrawRelations() error {
	return pgDb.CreateWithdrawFunctions()
}

func DbGetWalletIdByActiveAcnt(acntAdr string, externalCur string) (walletId int64, err error) {
	return pgDb.GetWalletIdofActiveAcnt(acntAdr, externalCur)
}

func DbUpdateWithdrawSuccessful(withdrawId int64, txHash string, txApprovedTime time.Time) error {
	return pgDb.UpdateWithdrawSuccessful(withdrawId, txHash, txApprovedTime)
}

func DbInitWithdrawReq(walletId int64, amount float64, extCurAbv string) (withdrawId int64, err error) {
	return pgDb.InitWithdrawReq(walletId, amount, extCurAbv)
}

func DbUpdateWithdrawPaymentQueryId(walletId int64, reqIdPaymentServ int64) error {
	return pgDb.UpdateWithdrawPaymentQueryId(walletId, reqIdPaymentServ)
}
