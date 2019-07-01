package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors" // register postgresql driver
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/config"
)

func Setup(conf config.MxpConfig) error {

	fmt.Println("setupDb called CC   !!")
	kmyFunc()
	dbInit()

	return nil
}

func dbInit() {
	db, err := sqlx.Open("postgres", "m2m_db:@localhost/m2m_database?sslmode=disable")
	fmt.Println(db, err)

	a, err2 := db.Exec(`
	
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'wallet_type') THEN
CREATE TYPE WALLET_TYPE AS ENUM (
'SUPER_ADMIN',
 'USER'
);
END IF;
END$$;

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

CREATE TABLE IF NOT EXISTS wallet (
id SERIAL PRIMARY KEY,
fk_org_la INT UNIQUE NOT NULL, -- foreign_key LoRa app server DB
type WALLET_TYPE NOT NULL,
balance NUMERIC(28,18) NOT NULL   -- varchar is considered for all of the values rather than float
);


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
account_adr varchar(128) NOT NULL,
insert_time TIMESTAMP NOT NULL,
status FIELD_STATUS NOT NULL
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
	fmt.Println("a:", a)
	fmt.Println("err2:", err2)
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
