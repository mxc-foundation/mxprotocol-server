package db

import (
	"time"

	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
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

func DbUpdateWithdrawPaymentQueryId(withdrawId int64, reqIdPaymentServ int64) error {
	return pgDb.UpdateWithdrawPaymentQueryId(withdrawId, reqIdPaymentServ)
}

func DbGetWithdrawHist(walletId int64, offset int64, limit int64) ([]pstgDb.WithdrawHistRet, error) {
	return pgDb.GetWithdrawHist(walletId, offset, limit)
}
