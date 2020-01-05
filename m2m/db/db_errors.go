package db

import (
	"database/sql"

	"github.com/mxc-foundation/mxprotocol-server/m2m/types"
)

func dbErrorInit() {
	types.DbError.NoRowQueryRes = sql.ErrNoRows
}
