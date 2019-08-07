package db

import (
	"database/sql"
	migrate "github.com/rubenv/sql-migrate"
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/migrations"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/config"
)

func Setup(conf config.MxpConfig) error {
	dbp, err := openPostgresDBWithPing(conf)
	var db DBHandler
	i = &pg.PgDB
	if err != nil {
		return err
	} else {
		db = addDB(i, dbp)
		if db.DB == nil {
			return errors.New("db/Setup: wrong db driver")
		}
	}

	dbInit()

	if conf.PostgreSQL.Automigrate {
		log.Info("db/applying PostgreSQL data migrations")
		m := &migrate.AssetMigrationSource{
			Asset:    migrations.Asset,
			AssetDir: migrations.AssetDir,
			Dir:      "",
		}
		n, err := migrate.Exec(db.DB, "postgres", m, migrate.Up)
		if err != nil {
			return errors.Wrap(err, "db/applying PostgreSQL data migrations error")
		}
		log.WithField("count", n).Info("db/PostgreSQL data migrations applied")
	}

	return nil
}

func openPostgresDBWithPing(conf config.MxpConfig) (*sql.DB, error) {
	log.Debug("db/connect_db")

	d, err := sql.Open("postgres", conf.PostgreSQL.DSN)
	if err != nil {
		log.WithError(err).Error("db/connect_db")
		return nil, err
	}
	for i := 0; i <= 3; i++ {
		if err := d.Ping(); err != nil {
			log.WithError(err).Warning("db/ping_db")
			time.Sleep(2 * time.Second) // to be modified
		} else {
			return d, nil
		}
	}

	err = errors.New("db/ping_db: failed")
	log.Error(err)
	return nil, err
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

	if err := dbCreateWithdrawFeeTable(); err != nil {
		log.WithError(err).Fatal("db/dbCreateWithdrawFeeTable")
	}

	if err := initExtCurrencyTable(); err != nil {
		log.WithError(err).Fatal("db/initExtCurrencyTable")
	}

}
