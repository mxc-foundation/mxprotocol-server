package db

import (
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api"
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type extCurrencyDBInterface interface {
	CreateExtCurrencyTable() error
	InsertExtCurr(ec pg.ExtCurrency) (insertIndex int64, err error)
	GetExtCurrencyIdByAbbr(extCurrencyAbbr string) (int64, error)
}

var extCurrency extCurrencyDBInterface

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
	extCurrency = &pg.PgExtCurrency
	return extCurrency.CreateExtCurrencyTable()
}

func dbInsertExtCurr(ec pg.ExtCurrency) (insertIndex int64, err error) {
	return extCurrency.InsertExtCurr(ec)
}

func DbGetExtCurrencyIdByAbbr(extCurrencyAbbr string) (int64, error) {
	return extCurrency.GetExtCurrencyIdByAbbr(extCurrencyAbbr)
}
