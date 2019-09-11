package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type aggGatewayUsageDBInterface interface {
	CreateAggGwUsgTable() error
	InsertAggGwUsg(agu types.AggGwUsg) (insertIndex int64, err error)
}
var AggGatewayUsage = aggGatewayUsageDBInterface(&pg.PgAggGatewayUsage)

