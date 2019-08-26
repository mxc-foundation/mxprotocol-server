package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

func dbCreateTopupTable() error {
	return pg.PgDB.CreateTopupTable()
}

func dbCreateTopupRelations() error {
	return pg.PgDB.CreateTopupFunctions()
}

func DbAddTopUpRequest(acntAdrSender string, acntAdrRcvr string, txHash string, value float64, extCurAbv string) (topupId int64, err error) {
	return pg.PgDB.AddTopUpRequest(acntAdrSender, acntAdrRcvr, txHash, value, extCurAbv)
}

func castTopupHistRet(acntHist []pg.TopupHistRet, err1 error) (castedVal []TopupHistRet, err error) {
	for _, v := range acntHist {
		castedVal = append(castedVal, TopupHistRet(v))
	}
	return castedVal, err1
}

type TopupHistRet pg.TopupHistRet

func DbGetTopupHist(walletId int64, offset int64, limit int64) ([]TopupHistRet, error) {
	return castTopupHistRet(pg.PgDB.GetTopupHist(walletId, offset, limit))
}

func DbGetTopupHistRecCnt(walletId int64) (int64, error) {
	return pg.PgDB.GetTopupHistRecCnt(walletId)
}
