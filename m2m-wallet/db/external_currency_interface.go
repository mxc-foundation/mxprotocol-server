package db

import (
	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

func DbCreateExtCurrencyTable() error {
	return pgDb.CreateExtCurrencyTable()
}

func DbInsertExtCurr(ec pstgDb.ExtCurrency) error {
	return pgDb.InsertExtCurr(ec)
}
