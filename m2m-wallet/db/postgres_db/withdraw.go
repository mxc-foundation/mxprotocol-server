package postgres_db

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type TxStatus string // db:tx_status
const (
	PENDING    TxStatus = "PENDING"
	SUCCESSFUL TxStatus = "SUCCESSFUL"
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

	END$$;

		
	`)
	pgDbp.CreateWithdrawSuccessfulFunction()
	return errors.Wrap(err, "db: PostgreSQL connection error")
}

func (pgDbp DbSpec) CreateWithdrawSuccessfulFunction() error {
	_, err := pgDbp.Db.Exec(`

	CREATE OR REPLACE FUNCTION withdraw_success (a INT) RETURNS void
		LANGUAGE plpgsql
			AS $$
			BEGIN
			UPDATE withdraw
				SET	tx_stat = 'SUCCESSFUL',
				tx_approved_time = NOW()
			WHERE
				id = a ;
			END;
		$$;
	`)
	return errors.Wrap(err, "db: PostgreSQL connection error")
}

func (pgDbp DbSpec) InsertWithdraw(wdr Withdraw) (insertIndex int, err error) {
	err = pgDbp.Db.QueryRow(`
		INSERT INTO withdraw (
			fk_ext_account_sender,
			fk_ext_account_receiver,
			fk_ext_currency,
			value,
			fk_withdraw_fee,
			tx_sent_time,
			tx_stat,
			tx_approved_time,
			fk_query_id_payment_service,
			tx_hash )
		VALUES (
			$1,	$2,	$3,	$4,	$5,	$6,	$7,	$8,	$9,	$10
			)
			RETURNING id;
	`,
		wdr.FkExtAcntSender,
		wdr.FkExtAcntRcvr,
		wdr.FkExtCurr,
		wdr.Value,
		wdr.FkWithdrawFee,
		wdr.TxSentTime,
		wdr.TxStatus,
		wdr.TxAprvdTime,
		wdr.FkQueryIdePaymentService,
		wdr.TxHash,
	).Scan(&insertIndex)

	return insertIndex, errors.Wrap(err, "db: query error InsertWithdrawFee()")
}

func (pgDbp DbSpec) UpdateWithdrawSuccessful(withdrawId int) error {
	_, err := pgDbp.Db.Exec(`
		
		select withdraw_success($1);
		
	`, withdrawId)
	return errors.Wrap(err, "db: PostgreSQL connection error")
}

func (pgDbp DbSpec) ApplyWithdrawReq(wdr Withdraw, it InternalTx) (err error) {
	err = pgDbp.Db.QueryRow(`
		DO $$
		BEGIN
		INSERT INTO withdraw (
			fk_ext_account_sender,
			fk_ext_account_receiver,
			fk_ext_currency,
			value,
			fk_withdraw_fee,
			tx_sent_time,
			tx_stat,
			tx_approved_time,
			fk_query_id_payment_service,
			tx_hash )
		VALUES (
			$1,	$2,	$3,	$4,	$5,	$6,	$7,	$8,	$9,	$10
			)
			;


		INSERT INTO internal_tx (
			fk_wallet_sernder,
			fk_wallet_receiver,
			payment_cat,
			tx_internal_ref,
			value,
			time_tx )
			VALUES (
			$11,
			$12,
			$13,
			$14,
			$15,
			$16)
			;


			UPDATE wallet 
			SET
				balance = balance - $17 
			WHERE
				id = $18
				; 


			UPDATE wallet 
			SET
				balance = balance + $17 
			WHERE
				id = $19 
				;

	END;
	$$;
	`,
		wdr.FkExtAcntSender,
		wdr.FkExtAcntRcvr,
		wdr.FkExtCurr,
		wdr.Value,
		wdr.FkWithdrawFee,
		wdr.TxSentTime,
		wdr.TxStatus,
		wdr.TxAprvdTime,
		wdr.FkQueryIdePaymentService,
		wdr.TxHash,
		it.FkWalletSender,
		it.FkWalletRcvr,
		it.PaymentCat,
		it.TxInternalRef,
		it.Value,
		it.TimeTx,
		it.Value,
		it.FkWalletSender,
		it.FkWalletRcvr).Scan()

	return errors.Wrap(err, "db: query error ApplyWithdrawReq()")
}
