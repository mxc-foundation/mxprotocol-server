package postgres_db

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Withdraw struct {
	Id                       int       `db:"id"`
	FkExtAcntSender          int       `db:"fk_ext_account_sender"`
	FkExtAcntRcvr            int       `db:"fk_ext_account_receiver"`
	FkExtCurr                int       `db:"fk_ext_currency"`
	Value                    float64   `db:"value"`
	FkWithdrawFee            int       `db:"fk_withdraw_fee"`
	TxSentTime               time.Time `db:"tx_sent_time"`
	TxStatus                 string    `db:"tx_status"`
	TxAprvdTime              time.Time `db:"tx_approved_time"`
	FkQueryIdePaymentService int       `db:"fk_query_id_payment_service"`
	TxHash                   string    `db:"tx_hash"`
}

func (pgDbp DbSpec) CreateWithdrawTable() error {
	_, err := pgDbp.Db.Exec(`

	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tx_status') THEN
	CREATE TYPE TX_STATUS AS ENUM (
	'PENDING',
	'SUCCESSFUL'
	);
	END IF;
	END$$;


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
	`)
	return errors.Wrap(err, "db: PostgreSQL connection error")
}
