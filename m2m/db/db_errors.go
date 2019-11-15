package db

import (
	"database/sql"

	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

func dbErrorInit() {
	types.DbError.NoRowQueryRes = sql.ErrNoRows
}
