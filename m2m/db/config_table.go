package db

import (
	pg "github.com/mxc-foundation/mxprotocol-server/m2m/db/postgres_db"
)

type configTableDBInterface interface {
	CreateConfigTable() error
	UpdateConfig(key string, value string) (err error)
	UpdateConfigs(data map[string]interface{}) (err error)
	InsertConfig(key string, value string) (err error)
	InsertConfigs(data map[string]interface{}, ignoreDuplicateKey bool) (err error)
	GetConfig(key string) (val string, err error)
	GetConfigs(keys []string) (configs []pg.Config, err error)
}

var ConfigTable = configTableDBInterface(&pg.PgConfigTable)
