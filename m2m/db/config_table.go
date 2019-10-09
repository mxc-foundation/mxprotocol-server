package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type configTableDBInterface interface {
	CreateConfigTable() error
	UpdateConfig(key string, value string) (err error)
	InsertConfig(key string, value string) (err error)
	GetConfig(key string) (val string, err error)
}

var ConfigTable = configTableDBInterface(&pg.PgConfigTable)
