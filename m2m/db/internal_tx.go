package db

import (
	pg "github.com/mxc-foundation/mxprotocol-server/m2m/db/postgres_db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/types"
)

type internalTxDBInterface interface {
	CreateInternalTxTable() error
	InsertInternalTx(it types.InternalTx) (insertIndex int64, err error)
}

var InternalTx = internalTxDBInterface(&pg.PgInternalTx)
