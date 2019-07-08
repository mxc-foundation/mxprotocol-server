package db

import (
	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

func DbCreateWithdrawTable() error {
	return pgDb.CreateWithdrawTable()
}

func DbInsertWithdraw(wdr pstgDb.Withdraw) (insertIndex int, err error) {
	return pgDb.InsertWithdraw(wdr)
}

// should be use for topup
func DbGetWalletIdofActiveAcnt(acntAdr string, externalCur string) (walletId int, err error) {
	return pgDb.GetWalletIdofActiveAcnt(acntAdr, externalCur)
}

func DbCreateWithdrawSuccessfulFunction() error {
	return pgDb.CreateWithdrawSuccessfulFunction()
}

func DbUpdateWithdrawSuccessful(withdrawId int) error {
	return pgDb.UpdateWithdrawSuccessful(withdrawId)
}

func DbApplyWithdrawReq(wdr pstgDb.Withdraw, it pstgDb.InternalTx) error {
	return pgDb.ApplyWithdrawReq(wdr, it)
}

func DbApplyWithdrawReq2(wdr pstgDb.Withdraw, it pstgDb.InternalTx) error {
	return pgDb.ApplyWithdrawReq2(wdr, it)
}
