package db

import (
	"github.com/apex/log"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

var CurrencyList = []pstgDb.ExtCurrency{}

func initExtCurrencyTable() error {
	currency := pstgDb.ExtCurrency{}
	for _, v := range api.Money_name {
		currency.Id = 0
		currency.Name = v
		currency.Abv = v
		CurrencyList = append(CurrencyList, currency)
	}

	for _, element := range CurrencyList {
		DbInsertExtCurr(element)
	}
	return nil
}

func DbCreateExtCurrencyTable() error {
	return pgDb.CreateExtCurrencyTable()
}

func DbInsertExtCurr(ec pstgDb.ExtCurrency) (insertIndex int64, err error) {
	return pgDb.InsertExtCurr(ec)
}

func DbGetExtCurrencyIdByAbbr(extCurrencyAbbr string) (int64, error) {
	id, err := pgDb.GetExtCurrencyIdByAbbr(extCurrencyAbbr)
	if err != nil {
		log.WithError(err).Error("DbGetExtCurrencyIdByAbbr")
	}

	return id, err
}
