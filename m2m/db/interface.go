package db

import (
	"database/sql"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/config"
	"time"
)

type DBInterface interface {
	AddDB(d *sql.DB)
}
var i DBInterface

type TxHandler struct {
	*sql.Tx
}

var dbM2M PostgresHandler

func openDB(inter DBInterface, conf config.MxpConfig) (err error) {
	var db *sql.DB
	// postgres
	if _, ok := inter.(*PostgresHandler); ok {
		db, err = openDBWithPing("postgres", conf.PostgreSQL.DSN)
		if err != nil {
			log.WithError(err).Error("db/openDB")
			return err
		}

		i.AddDB(db)
		return nil

	}

	//if _, ok := inter.(*OtherDBHandler); ok {
	//		db, err = openDBWithPing(driverName, dsn)
	//		i.AddDB(db)
	//		return nil
	//	}

	return errors.New("db/openDB: unknown db driver")
}

func openDBWithPing(driverName string, dsn string) (*sql.DB, error) {
	log.Debug("db/openDBWithPing")

	d, err := sql.Open(driverName, dsn)
	if err != nil {
		log.WithError(err).Error("db/openDBWithPing")
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