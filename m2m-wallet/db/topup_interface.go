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

func DbApplyTopup(tu pstgDb.Topup, it pstgDb.InternalTx) error {
	// return pgDb.ApplyTopupReq(tu, it)  // to add
	return nil
}

func DbAddTopUpRequest(acntAdrSender string, acntAdrRcvr string, txHash string, value string) (topupID int64, err error) {
	return 0, nil
}
