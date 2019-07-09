package db

import (
	"time"

	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

func DbCreateWithdrawTable() error {
	return pgDb.CreateWithdrawTable()
}

func DbInsertWithdraw(wdr pstgDb.Withdraw) (insertIndex int64, err error) {
	return pgDb.InsertWithdraw(wdr)
}

// should be use for topup
func DbGetWalletIdofActiveAcnt(acntAdr string, externalCur string) (walletId int64, err error) {
	return pgDb.GetWalletIdofActiveAcnt(acntAdr, externalCur)
}

func DbCreateWithdrawFunctions() error {
	return pgDb.CreateWithdrawFunctions()
}

func DbUpdateWithdrawSuccessful(withdrawId int64, txHash string, txApprovedTime time.Time) error {
	return pgDb.UpdateWithdrawSuccessful(withdrawId, txHash, txApprovedTime)
}

func DbInitWithdrawReqApply(wdr pstgDb.Withdraw, it pstgDb.InternalTx) (withdrawId int64, err error) {
	return pgDb.InitWithdrawReqApply(wdr, it)
}

func DbInitWithdrawReq(walletId int64, amount float64, extCurAbv string) (withdrawId int64, err error) {
	return pgDb.InitWithdrawReq(walletId, amount, extCurAbv)
}

func DbUpdateWithdrawPaymentQueryId(walletId int64, reqIdPaymentServ int64) error {
	return pgDb.UpdateWithdrawPaymentQueryId(walletId, reqIdPaymentServ)
}

//db.DbUpdateWithdrawPaymentQueryId(walletID, reqId_paymentsev)
