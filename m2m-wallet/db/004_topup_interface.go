package db

func dbCreateTopupTable() error {
	return dbHandler.CreateTopupTable()
}

func dbCreateTopupRelations() error {
	return dbHandler.CreateTopupFunctions()
}

func DbAddTopUpRequest(acntAdrSender string, acntAdrRcvr string, txHash string, value float64, extCurAbv string) (topupId int64, err error) {
	return dbHandler.AddTopUpRequest(acntAdrSender, acntAdrRcvr, txHash, value, extCurAbv)
}

func DbGetTopupHist(walletId int64, offset int64, limit int64) ([]TopupHistRet, error) {
	return dbHandler.GetTopupHist(walletId, offset, limit)
}
