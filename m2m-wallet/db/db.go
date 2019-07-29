package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/config"
)

var pgDb pstgDb.DbSpec

func Setup(conf config.MxpConfig) error {
	dbp, err := openDBWithPing(conf)

	if err != nil {
		return err
	} else {
		pgDb = pstgDb.DbSpec{
			Db:         dbp,
			DriverName: "postgres",
			Dburl:      conf.PostgreSQL.DSN,
		}
	}

	dbInit()

	return nil
}

func openDBWithPing(conf config.MxpConfig) (*sql.DB, error) {
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
