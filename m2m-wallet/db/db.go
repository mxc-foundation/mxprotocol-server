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
	// dbp, err := sql.Open("postgres", "postgres://m2m_db@postgres:5432/m2m_database?sslmode=disable")

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
	err = dbInit()
	if err != nil {
		log.Fatal("Unable to init DB!", err)
	}
	//testDb()

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

func dbInit() error {
	// db, err := sql.Open("postgres", "postgres://m2m_db@postgres:5432/m2m_database?sslmode=disable")
	// fmt.Println(db, err)

	_, err := pgDb.Db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tx_status') THEN
		CREATE TYPE TX_STATUS AS ENUM (
		'PENDING',
		'SUCCESSFUL'
		);
		END IF;
		END$$;

		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'field_status') THEN
		CREATE TYPE FIELD_STATUS AS ENUM (
		'ACTIVE', 'ARC');
		END IF;
		END$$;


		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'ext_url_type') THEN
		CREATE TYPE EXT_URL_TYPE AS ENUM (
		'ACCOUNT_BALANCE_CHECK',
		'TX_STATUS_CHECK'
		);
		END IF;
		END$$;


		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_category') THEN
		CREATE TYPE PAYMENT_CATEGORY AS ENUM (
		'UPLINK',
		'DOWNLINK',
		'PURCHASE_SUBSCRIPTION',
		'gateway_income',
		'TOP_UP',
		'WITHDRAW'
		);
		END IF;
		END$$;


		-- CREATE tables:



		CREATE TABLE IF NOT EXISTS internal_tx (
		id SERIAL PRIMARY KEY,
		fk_wallet_sernder INT REFERENCES wallet(id) NOT NULL,
		fk_wallet_receiver INT REFERENCES wallet(id) NOT NULL,
		payment_cat PAYMENT_CATEGORY,
		tx_internal_ref INT NOT NULL,
		value NUMERIC(28,18),
		CONSTRAINT payment_cat_tx_internal_ref UNIQUE (payment_cat, tx_internal_ref)
		);


		CREATE TABLE IF NOT EXISTS ext_currency (
		id SERIAL PRIMARY KEY,
		name VARCHAR(64),
		abv VARCHAR(16) UNIQUE NOT NULL
		);

		-- can be handeled independant of the DB (simply we can drop this table):
		CREATE TABLE IF NOT EXISTS api_url_ext_curr (
		id SERIAL PRIMARY KEY,
		fk_external_currency INT REFERENCES ext_currency (id) NOT NULL,
		url VARCHAR(256) NOT NULL,
		status FIELD_STATUS NOT NULL,
		type EXT_URL_TYPE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS ext_account (
			id SERIAL PRIMARY KEY,
			fk_wallet INT REFERENCES wallet(id) NOT NULL,
			fk_ext_currency INT REFERENCES ext_currency (id) NOT NULL,
			account_adr varchar(128) NOT NULL UNIQUE,
			insert_time TIMESTAMP NOT NULL,
			status FIELD_STATUS NOT NULL,
			latest_checked_block INT DEFAULT 0
		);

		CREATE TABLE IF NOT EXISTS withdraw_fee (
		id SERIAL PRIMARY KEY,
		fk_ext_currency INT REFERENCES ext_currency (id) NOT NULL,
		fee NUMERIC(28,18) NOT NULL,
		insert_time TIMESTAMP NOT NULL,
		status FIELD_STATUS NOT NULL
		);

		CREATE TABLE IF NOT EXISTS withdraw (
		id SERIAL PRIMARY KEY,
		fk_ext_account_sender INT REFERENCES  ext_account(id) NOT NULL,
		fk_ext_account_receiver INT REFERENCES  ext_account(id) NOT NULL,
		fk_ext_currency INT REFERENCES  ext_currency(id) NOT NULL,
		value NUMERIC(28,18) NOT NULL,
		fk_withdraw_fee INT REFERENCES  withdraw(id) NOT NULL,
		tx_sent_time TIMESTAMP NOT NULL,
		tx_stat tx_status NOT NULL,
		tx_approved_time TIMESTAMP,
		fk_query_id_payment_service INT NOT NULL,
		tx_hash varchar (128)
		);


		CREATE TABLE IF NOT EXISTS topup (
		id SERIAL PRIMARY KEY,
		fk_ext_account_sender INT REFERENCES  ext_account(id) NOT NULL,
		fk_ext_account_receiver INT REFERENCES  ext_account(id) NOT NULL,
		fk_ext_currency INT REFERENCES ext_currency(id) NOT NULL,
		value NUMERIC(28,18) NOT NULL,
		tx_approved_time TIMESTAMP,
		tx_hash varchar (128) NOT NULL
		);


	`,
	)
	return err
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
