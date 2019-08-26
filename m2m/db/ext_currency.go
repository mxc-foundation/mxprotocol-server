package db

import (
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
)

var CurrencyList = []pg.ExtCurrency{}

func initExtCurrencyTable() error {
	currency := pg.ExtCurrency{}
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
		if _, err := dbInsertExtCurr(element); err != nil {
			return err
		}
	}
	return nil
}

func dbCreateExtCurrencyTable() error {
	return pg.PgDB.CreateExtCurrencyTable()
}

func dbInsertExtCurr(ec pg.ExtCurrency) (insertIndex int64, err error) {
	return pg.PgDB.InsertExtCurr(ec)
}

func DbGetExtCurrencyIdByAbbr(extCurrencyAbbr string) (int64, error) {
	return pg.PgDB.GetExtCurrencyIdByAbbr(extCurrencyAbbr)
}
