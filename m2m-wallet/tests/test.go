package tests

import (
	"database/sql"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/config"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/migrations"
	"os"
)

func GetConfig() config.MxpConfig {
	log.SetLevel(log.ErrorLevel)
	var c config.MxpConfig

	if v := os.Getenv("TEST_POSTGRES_DSN"); v != "" {
		c.PostgreSQL.DSN = v
	}

	if v := os.Getenv("TEST_LORA_APP_SERVER"); v != "" {
		c.General.AuthServer = v
	}

	return c
}

func ResetDB(db *sql.DB) {
	m := &migrate.AssetMigrationSource{
		Asset:    migrations.Asset,
		AssetDir: migrations.AssetDir,
		Dir:      "",
	}
	if _, err := migrate.Exec(db, "postgres", m, migrate.Down); err != nil {
		log.Fatal(err)
	}
	if _, err := migrate.Exec(db, "postgres", m, migrate.Up); err != nil {
		log.Fatal(err)
	}
}
