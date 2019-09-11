package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type internalTxDBInterface interface {
	CreateInternalTxTable() error
	InsertInternalTx(it types.InternalTx) (insertIndex int64, err error)
}

var InternalTx = internalTxDBInterface(&pg.PgInternalTx)