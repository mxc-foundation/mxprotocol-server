package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
)

type configTableDBInterface interface {
	CreateConfigTable() error
	Insert(config *pg.Config, ignoreDuplicateKey bool) (err error)
	Update(config *pg.Config) (err error)
	UpdateOne(key pg.ConfigKey, value string) (err error)
	Get() (config *pg.Config, err error)
	GetOne(key pg.ConfigKey) (value string, err error)
}

var ConfigTable = configTableDBInterface(&pg.PgConfigTable)
