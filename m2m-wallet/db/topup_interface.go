package db

import (
	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

func DbCreateTopupTable() error {
	return pgDb.CreateTopupTable()
}

func DbInsertTopup(tu pstgDb.Topup) (insertIndex int64, err error) {
	return pgDb.InsertTopup(tu)
}

func DbCreateTopupFunctions() error {
	return pgDb.CreateTopupFunctions()
}

func DbApplyTopup(tu pstgDb.Topup, it pstgDb.InternalTx) (topupID int64, err error) {
	return pgDb.ApplyTopup(tu, it)

}

func DbAddTopUpRequest(acntAdrSender string, acntAdrRcvr string, txHash string, value float64, extCurAbv string) (topupID int64, err error) {
	return pgDb.AddTopUpRequest(acntAdrSender, acntAdrRcvr, txHash, value, extCurAbv)
}
