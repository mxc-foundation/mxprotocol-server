package db

import (
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/migrations"
)

func Setup(conf config.MxpConfig) error {
	i = &dbM2M
	err := openDB(i, conf)
	if err != nil {
		return err
	}

	dbInit()

	if conf.PostgreSQL.Automigrate {
		log.Info("db/applying PostgreSQL data migrations")
		m := &migrate.AssetMigrationSource{
			Asset:    migrations.Asset,
			AssetDir: migrations.AssetDir,
			Dir:      "",
		}
		n, err := migrate.Exec(dbM2M.DB, "postgres", m, migrate.Up)
		if err != nil {
			return errors.Wrap(err, "db/applying PostgreSQL data migrations error")
		}
		log.WithField("count", n).Info("db/PostgreSQL data migrations applied")
	}

	return nil
}


func dbInit() {
	dbErrorInit()

	if err := dbCreateWalletTable(); err != nil {
		log.WithError(err).Fatal("db/dbCreateWalletTable")
	}

	if err := dbCreateInternalTxTable(); err != nil {
		log.WithError(err).Fatal("db/dbCreateInternalTxTable")
	}

	if err := dbCreateExtCurrencyTable(); err != nil {
		log.WithError(err).Fatal("db/dbCreateExtCurrencyTable")
	}

	if err := dbCreateExtAccountTable(); err != nil {
		log.WithError(err).Fatal("db/dbCreateExtAccountTable")
	}

	if err := dbCreateWithdrawFeeTable(); err != nil {
		log.WithError(err).Fatal("db/dbCreateWithdrawFeeTable")
	}

	if err := dbCreateWithdrawTable(); err != nil {
		log.WithError(err).Fatal("db/dbCreateWithdrawTable")
	}

	if err := dbCreateWithdrawRelations(); err != nil {
		log.WithError(err).Fatal("db/dbCreateWithdrawRelations")
	}

	if err := dbCreateTopupTable(); err != nil {
		log.WithError(err).Fatal("db/dbCreateTopupTable")
	}

	if err := dbCreateTopupRelations(); err != nil {
		log.WithError(err).Fatal("db/dbCreateTopupRelations")
	}

	if err := initExtCurrencyTable(); err != nil {
		log.WithError(err).Fatal("db/initExtCurrencyTable")
	}

}
