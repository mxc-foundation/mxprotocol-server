package db

import (
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

var CurrencyList = []pstgDb.ExtCurrency{}

func initExtCurrencyTable() error {
	currency := pstgDb.ExtCurrency{}
	for _, v := range api.Money_name {
		if _, err := DbGetExtCurrencyIdByAbbr(v); err == nil {
			continue
		}

		currency.Id = 0
		currency.Name = v
		currency.Abv = v
		CurrencyList = append(CurrencyList, currency)
	}

	for _, element := range CurrencyList {
		if _, err := DbInsertExtCurr(element); err != nil {
			return err
		}
	}
	return nil
}

func dbCreateExtCurrencyTable() error {
	return pgDb.CreateExtCurrencyTable()
}

func DbInsertExtCurr(ec pstgDb.ExtCurrency) (insertIndex int64, err error) {
	return pgDb.InsertExtCurr(ec)
}

func DbGetExtCurrencyIdByAbbr(extCurrencyAbbr string) (int64, error) {
	return pgDb.GetExtCurrencyIdByAbbr(extCurrencyAbbr)
}
