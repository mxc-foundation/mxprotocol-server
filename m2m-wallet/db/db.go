package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors" // register postgresql driver
	log "github.com/sirupsen/logrus"
	pstgDb "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/config"
)

var pgDb pstgDb.DbSpec

func Setup(conf config.MxpConfig) error {

	fmt.Println("The DB url: " + conf.PostgreSQL.DSN)
	pingCheck(conf)
	dbp, err := sql.Open("postgres", conf.PostgreSQL.DSN)

	if err != nil {
		log.Fatal("No DB accessable!", err)
	} else {
		pgDb = pstgDb.DbSpec{
			Db:         dbp,
			DriverName: "postgres",
			Dburl:      conf.PostgreSQL.DSN,
		}
	}

	// create tables if not exist
	dbInit()
	// testDb()

	// init data if applys
	err = initExtCurrencyTable()
	if err != nil {
		log.WithError(err).Fatal("Create init data in ext_currency failed.")
	}

	return nil
}

func pingCheck(conf config.MxpConfig) error {
	log.Info("storage: connecting to PostgreSQL database")
	d, err := sql.Open("postgres", conf.PostgreSQL.DSN)
	if err != nil {
		return errors.Wrap(err, "storage: PostgreSQL connection error")
	}
	for {
		if err := d.Ping(); err != nil {
			log.WithError(err).Warning("storage: ping PostgreSQL database error, will retry in 2s")
			time.Sleep(2 * time.Second) // to be modified
		} else {
			break
		}
	}
	return nil
}

func dbInit() {
	if err := DbCreateWalletTable(); err != nil {
		log.Fatal("Unable to create table wallet!", err)
	}

	if err := DbCreateInternalTxTable(); err != nil {
		log.Fatal("Unable to create table internal_tx!", err)
	}

	if err := DbCreateExtCurrencyTable(); err != nil {
		log.Fatal("Unable to create table ext_currency!", err)
	}

	if err := DbCreateExtAccountTable(); err != nil {
		log.Fatal("Unable to create table ext_account!", err)
	}

	if err := DbCreateWithdrawTable(); err != nil {
		log.Fatal("Unable to create table withdraw!", err)
	}

	if err := DbCreateWithdrawFunctions(); err != nil {
		log.Fatal("Unable to create table withdraw!", err)
	}

	if err := DbCreateTopupTable(); err != nil {
		log.Fatal("Unable to create table top_up!", err)
	}

	if err := DbCreateTopupFunctions(); err != nil {
		log.Fatal("Unable to create table top_up!", err)
	}

	if err := DbCreateWithdrawFeeTable(); err != nil {
		log.Fatal("Unable to create table withdraw_fee!", err)
	}

}

// db holds the PostgreSQL connection pool.
var db *DBLogger

const (
	redisDialWriteTimeout = time.Second
	redisDialReadTimeout  = time.Minute
	onBorrowPingInterval  = time.Minute
)

// DBLogger is a DB wrapper which logs the executed sql queries and their
// duration.
type DBLogger struct {
	*sqlx.DB
}

// Beginx returns a transaction with logging.
func (db *DBLogger) Beginx() (*TxLogger, error) {
	tx, err := db.DB.Beginx()
	return &TxLogger{tx}, err
}

// Query logs the queries executed by the Query method.
func (db *DBLogger) Query(query string, args ...interface{}) (*sql.Rows, error) {
	start := time.Now()
	rows, err := db.DB.Query(query, args...)
	logQuery(query, time.Since(start), args...)
	return rows, err
}

// Queryx logs the queries executed by the Queryx method.
func (db *DBLogger) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	start := time.Now()
	rows, err := db.DB.Queryx(query, args...)
	logQuery(query, time.Since(start), args...)
	return rows, err
}

// QueryRowx logs the queries executed by the QueryRowx method.
func (db *DBLogger) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	start := time.Now()
	row := db.DB.QueryRowx(query, args...)
	logQuery(query, time.Since(start), args...)
	return row
}

// Exec logs the queries executed by the Exec method.
func (db *DBLogger) Exec(query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	res, err := db.DB.Exec(query, args...)
	logQuery(query, time.Since(start), args...)
	return res, err
}

// TxLogger logs the executed sql queries and their duration.
type TxLogger struct {
	*sqlx.Tx
}

// Query logs the queries executed by the Query method.
func (q *TxLogger) Query(query string, args ...interface{}) (*sql.Rows, error) {
	start := time.Now()
	rows, err := q.Tx.Query(query, args...)
	logQuery(query, time.Since(start), args...)
	return rows, err
}

// Queryx logs the queries executed by the Queryx method.
func (q *TxLogger) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	start := time.Now()
	rows, err := q.Tx.Queryx(query, args...)
	logQuery(query, time.Since(start), args...)
	return rows, err
}

// QueryRowx logs the queries executed by the QueryRowx method.
func (q *TxLogger) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	start := time.Now()
	row := q.Tx.QueryRowx(query, args...)
	logQuery(query, time.Since(start), args...)
	return row
}

// Exec logs the queries executed by the Exec method.
func (q *TxLogger) Exec(query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	res, err := q.Tx.Exec(query, args...)
	logQuery(query, time.Since(start), args...)
	return res, err
}

func logQuery(query string, duration time.Duration, args ...interface{}) {
	log.WithFields(log.Fields{
		"query":    query,
		"args":     args,
		"duration": duration,
	}).Debug("sql query executed")
}

// DB returns the PostgreSQL database object.
func DB() *DBLogger {
	return db
}

// Transaction wraps the given function in a transaction. In case the given
// functions returns an error, the transaction will be rolled back.
func Transaction(f func(tx sqlx.Ext) error) error {
	tx, err := db.Beginx()
	if err != nil {
		return errors.Wrap(err, "storage: begin transaction error")
	}

	err = f(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return errors.Wrap(rbErr, "storage: transaction rollback error")
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "storage: transaction commit error")
	}
	return nil
}
