package db

func dbCreateTopupTable() error {
	return db.CreateTopupTable()
}

func dbCreateTopupRelations() error {
	return db.CreateTopupFunctions()
}

func DbAddTopUpRequest(acntAdrSender string, acntAdrRcvr string, txHash string, value float64, extCurAbv string) (topupId int64, err error) {
	return db.AddTopUpRequest(acntAdrSender, acntAdrRcvr, txHash, value, extCurAbv)
}

func DbGetTopupHist(walletId int64, offset int64, limit int64) ([]TopupHistRet, error) {
	return db.GetTopupHist(walletId, offset, limit)
}
