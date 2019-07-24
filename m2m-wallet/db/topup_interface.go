package db

import (
	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

func dbCreateTopupTable() error {
	return pgDb.CreateTopupTable()
}

func dbCreateTopupRelations() error {
	return pgDb.CreateTopupFunctions()
}

func DbAddTopUpRequest(acntAdrSender string, acntAdrRcvr string, txHash string, value float64, extCurAbv string) (topupId int64, err error) {
	return pgDb.AddTopUpRequest(acntAdrSender, acntAdrRcvr, txHash, value, extCurAbv)
}

func DbGetTopupHist(walletId int64, offset int64, limit int64) ([]pstgDb.TopupHistRet, error) {
	return pgDb.GetTopupHist(walletId, offset, limit)
}
