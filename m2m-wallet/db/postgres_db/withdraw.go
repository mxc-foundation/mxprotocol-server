package postgres_db

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type TxStatus string // db:tx_status
const (
	NOT_SENT_TO_PS TxStatus = "NOT_SENT_TO_PS"
	PENDING        TxStatus = "PENDING"
	SUCCESSFUL     TxStatus = "SUCCESSFUL"
)

type Withdraw struct {
	Id                       int64     `db:"id"`
	FkExtAcntSender          int64     `db:"fk_ext_account_sender"`
	FkExtAcntRcvr            int64     `db:"fk_ext_account_receiver"`
	FkExtCurr                int64     `db:"fk_ext_currency"`
	Value                    float64   `db:"value"`
	FkWithdrawFee            int64     `db:"fk_withdraw_fee"`
	TxSentTime               time.Time `db:"tx_sent_time"`
	TxStatus                 string    `db:"tx_status"`
	TxAprvdTime              time.Time `db:"tx_approved_time"`
	FkQueryIdePaymentService int64     `db:"fk_query_id_payment_service"`
	TxHash                   string    `db:"tx_hash"`
}

func (pgDbp DbSpec) CreateWithdrawTable() error {
	_, err := pgDbp.Db.Exec(`

	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tx_status') THEN
	CREATE TYPE TX_STATUS AS ENUM (
		'NOT_SENT_TO_PS',
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
		fk_query_id_payment_service INT ,
		tx_hash varchar (128) 
		);

	END$$;

		
	`)
	pgDbp.CreateWithdrawSuccessfulFunction()
	return errors.Wrap(err, "db: PostgreSQL connection error")
}

func (pgDbp DbSpec) CreateWithdrawSuccessfulFunction() error {
	_, err := pgDbp.Db.Exec(`

	
	`)
	return errors.Wrap(err, "db: PostgreSQL connection error")
}

func (pgDbp DbSpec) InsertWithdraw(wdr Withdraw) (insertIndex int64, err error) {
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

func (pgDbp DbSpec) UpdateWithdrawSuccessful(withdrawId int64, txHash string, txApprovedTime time.Time) error {
	_, err := pgDbp.Db.Exec(`
		select withdraw_success($1,$2,$3);
		
	`, withdrawId, txHash, txApprovedTime)
	return errors.Wrap(err, "db: PostgreSQL connection error UpdateWithdrawSuccessful()")
}

func (pgDbp DbSpec) CreateWithdrawFunctions() error {
	_, err := pgDbp.Db.Exec(`



	CREATE OR REPLACE FUNCTION withdraw_success (withdrawId INT, txHash varchar(128), txAprvdTime TIMESTAMP) RETURNS void
		LANGUAGE plpgsql
			AS $$
			BEGIN
			UPDATE withdraw
				SET	tx_stat = 'SUCCESSFUL',
				tx_approved_time = txAprvdTime,
				tx_hash = txHash
			WHERE
				id = withdrawId ;
			END;
		$$;




	CREATE OR REPLACE FUNCTION withdraw_req_init (
		v_fk_ext_account_sender INT,
		v_fk_ext_account_receiver INT,
		v_fk_ext_currency INT,
		v_value NUMERIC(28,18),
		v_fk_withdraw_fee INT,
		v_tx_sent_time TIMESTAMP,
		v_tx_stat tx_status,
		v_fk_wallet_sernder INT,
		v_fk_wallet_receiver INT,
		v_payment_cat PAYMENT_CATEGORY,
		v_value_fee_included NUMERIC(28,18)
		) RETURNS INT
	LANGUAGE plpgsql
	AS $$
	
	declare wdr_id INT;

	BEGIN
		INSERT INTO withdraw (
			fk_ext_account_sender,
			fk_ext_account_receiver,
			fk_ext_currency,
			value,
			fk_withdraw_fee,
			tx_sent_time,
			tx_stat)
		VALUES (
			v_fk_ext_account_sender ,
			v_fk_ext_account_receiver,
			v_fk_ext_currency,
			v_value,
			v_fk_withdraw_fee,
			v_tx_sent_time,
			v_tx_stat
		)RETURNING id INTO wdr_id;


		INSERT INTO internal_tx (
			fk_wallet_sernder,
			fk_wallet_receiver,
			payment_cat,
			tx_internal_ref,
			value,
			time_tx )
		VALUES (
			v_fk_wallet_sernder,
			v_fk_wallet_receiver,
			v_payment_cat,
			wdr_id,
			v_value_fee_included,
			v_tx_sent_time)
		;


		UPDATE
			wallet 
		SET
			balance = balance - v_value_fee_included
		WHERE
			id = v_fk_wallet_sernder
		;

	RETURN wdr_id;

	END;
	$$;

	`)

	return err
}

func (pgDbp DbSpec) InitWithdrawReqApply(wdr Withdraw, it InternalTx) (withdrawId int64, err error) {

	err = pgDbp.Db.QueryRow(`
		select withdraw_req_init($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11);
		
	`, wdr.FkExtAcntSender,
		wdr.FkExtAcntRcvr,
		wdr.FkExtCurr,
		wdr.Value,
		wdr.FkWithdrawFee,
		wdr.TxSentTime,
		wdr.TxStatus,
		it.FkWalletSender,
		it.FkWalletRcvr,
		it.PaymentCat,
		it.Value).Scan(&withdrawId)

	return withdrawId, errors.Wrap(err, "db: PostgreSQL connection error InitWithdrawReqApply()")

}

func (pgDbp DbSpec) InitWithdrawReq(walletId int64, value float64, extCurrencyAbbr string) (withdrawId int64, err error) {

	wdr := Withdraw{
		Value:      value,
		TxSentTime: time.Now().UTC(),
		TxStatus:   string(NOT_SENT_TO_PS),
	}

	wdr.FkExtAcntSender, err = pgDbp.GetUserExtAccountId(walletId, extCurrencyAbbr)
	if err != nil {
		return withdrawId, errors.Wrap(err, "db: InitWithdrawReq query error GetUserExtAccountId()")
	}

	wdr.FkExtAcntRcvr, err = pgDbp.GetSuperNodeExtAccountId(extCurrencyAbbr)
	if err != nil {
		return withdrawId, errors.Wrap(err, "db: InitWithdrawReq query error GetSuperNodeExtAccountId()")
	}

	wdr.FkExtCurr, err = pgDbp.GetExtCurrencyIdByAbbr(extCurrencyAbbr)
	if err != nil {
		return withdrawId, errors.Wrap(err, "db: InitWithdrawReq query error GetExtCurrencyIdByAbbr()")
	}

	wdr.FkWithdrawFee, err = pgDbp.GetActiveWithdrawFeeId(extCurrencyAbbr)
	if err != nil {
		return withdrawId, errors.Wrap(err, "db: InitWithdrawReq query error GetActiveWithdrawFeeId()")
	}

	var withdrawFeeAmnt float64
	withdrawFeeAmnt, err = pgDbp.GetActiveWithdrawFee(extCurrencyAbbr)
	if err != nil {
		return withdrawId, errors.Wrap(err, "db: InitWithdrawReq query error GetActiveWithdrawFee()")
	}

	it := InternalTx{
		FkWalletSender: walletId,
		PaymentCat:     string(WITHDRAW),
		Value:          value + withdrawFeeAmnt,
	}

	it.FkWalletRcvr, err = pgDbp.GetWalletIdSuperNode()
	if err != nil {
		return withdrawId, errors.Wrap(err, "db: InitWithdrawReq query error GetWalletIdSuperNode()")
	}

	return pgDbp.InitWithdrawReqApply(wdr, it)

}

func (pgDbp DbSpec) UpdateWithdrawPaymentQueryId(walletId int64, reqIdPaymentServ int64) error {
	_, err := pgDbp.Db.Exec(`
		UPDATE withdraw 
		SET
			tx_stat = 'PENDING',
			fk_query_id_payment_service = $1
		WHERE
			id = $2
		;
	
	`, reqIdPaymentServ,
		walletId)

	return errors.Wrap(err, "db: query error UpdateWithdrawPaymentQueryId()")
}
