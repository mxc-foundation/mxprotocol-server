package db

import (
	pg "github.com/mxc-foundation/mxprotocol-server/m2m/db/postgres_db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/types"
)

type aggGatewayUsageDBInterface interface {
	CreateAggGwUsgTable() error
	InsertAggGwUsg(agu types.AggGwUsg) (insertIndex int64, err error)
}

var AggGatewayUsage = aggGatewayUsageDBInterface(&pg.PgAggGatewayUsage)
