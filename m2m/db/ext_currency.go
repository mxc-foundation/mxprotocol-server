package db

import (
	pg "github.com/mxc-foundation/mxprotocol-server/m2m/db/postgres_db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/types"
)

type extCurrencyDBInterface interface {
	CreateExtCurrencyTable() error
	InsertExtCurr(ec types.ExtCurrency) (insertIndex int64, err error)
	GetExtCurrencyIdByAbbr(extCurrencyAbbr string) (int64, error)
}

var ExtCurrency = extCurrencyDBInterface(&pg.PgExtCurrency)
