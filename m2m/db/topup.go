package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type topupDBInterface interface {
	CreateTopupTable() error
	CreateTopupFunctions() error
	AddTopUpRequest(acntAdrSender string, acntAdrRcvr string, txHash string, value float64, extCurAbv string) (topupId int64, err error)
	GetTopupHist(walletId int64, offset int64, limit int64) ([]pg.TopupHistRet, error)
	GetTopupHistRecCnt(walletId int64) (recCnt int64, err error)
}

var topup topupDBInterface

func dbCreateTopupTable() error {
	topup = &pg.PgTopup
	return topup.CreateTopupTable()
}

func dbCreateTopupRelations() error {
	return topup.CreateTopupFunctions()
}

func DbAddTopUpRequest(acntAdrSender string, acntAdrRcvr string, txHash string, value float64, extCurAbv string) (topupId int64, err error) {
	return topup.AddTopUpRequest(acntAdrSender, acntAdrRcvr, txHash, value, extCurAbv)
}

func castTopupHistRet(acntHist []pg.TopupHistRet, err1 error) (castedVal []TopupHistRetConvert, err error) {
	for _, v := range acntHist {
		castedVal = append(castedVal, TopupHistRetConvert(v))
	}
	return castedVal, err1
}

type TopupHistRetConvert pg.TopupHistRet

func DbGetTopupHist(walletId int64, offset int64, limit int64) ([]TopupHistRetConvert, error) {
	acntHist, err := topup.GetTopupHist(walletId, offset, limit)
	return castTopupHistRet(acntHist, err)
}

func DbGetTopupHistRecCnt(walletId int64) (int64, error) {
	return topup.GetTopupHistRecCnt(walletId)
}
