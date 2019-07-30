package db

import (
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
)

var CurrencyList = []ExtCurrency{}

func initExtCurrencyTable() error {
	currency := ExtCurrency{}
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
	return dbHandler.CreateExtCurrencyTable()
}

func DbInsertExtCurr(ec ExtCurrency) (insertIndex int64, err error) {
	return dbHandler.InsertExtCurr(ec)
}

func DbGetExtCurrencyIdByAbbr(extCurrencyAbbr string) (int64, error) {
	return dbHandler.GetExtCurrencyIdByAbbr(extCurrencyAbbr)
}
