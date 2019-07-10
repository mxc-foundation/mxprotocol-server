package db

func dbCreateTopupTable() error {
	return pgDb.CreateTopupTable()
}

func dbCreateTopupRelations() error {
	return pgDb.CreateTopupFunctions()
}

func DbAddTopUpRequest(acntAdrSender string, acntAdrRcvr string, txHash string, value float64, extCurAbv string) (topupID int64, err error) {
	return pgDb.AddTopUpRequest(acntAdrSender, acntAdrRcvr, txHash, value, extCurAbv)
}
