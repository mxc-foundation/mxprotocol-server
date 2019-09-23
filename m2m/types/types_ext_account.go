package types

import "time"

type ExtAccountHistRet struct {
	AccountAdr     string
	InsertTime     time.Time
	Status         string
	ExtCurrencyAbv string
}

type ExtCurrency struct {
	Id   int64
	Name string
	Abv  string
}
