package postgres_db

import (
	"time"

	"github.com/ethereum/go-ethereum/log"
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

type WithdrawHistRet struct {
	AcntSender  string
	AcntRcvr    string
	ExtCurrency string
	Value       float64
	WithdrawFee float64
	TxSentTime  time.Time `db:"tx_sent_time"`
	TxStatus    string    `db:"tx_status"`
	TxAprvdTime time.Time
	TxHash      string
}

func  CreateWithdrawTable() error {
	_, err := PgDB.Exec(`

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
		fk_withdraw_fee INT REFERENCES  withdraw_fee(id) NOT NULL,
		tx_sent_time TIMESTAMP NOT NULL,
		tx_stat tx_status NOT NULL,
		tx_approved_time TIMESTAMP,
		fk_query_id_payment_service INT ,
		tx_hash varchar (128) 
		);

	END$$;

		
	`)

	return errors.Wrap(err, "db/CreateWithdrawTable")
}

func  insertWithdraw(wdr Withdraw) (insertIndex int64, err error) {
	err = PgDB.QueryRow(`
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

	return insertIndex, errors.Wrap(err, "db/InsertWithdraw")
}

func  UpdateWithdrawSuccessful(withdrawId int64, txHash string, txApprovedTime time.Time) error {
	_, err := PgDB.Exec(`
		SELECT withdraw_success($1,$2,$3);
		
	`, withdrawId, txHash, txApprovedTime)
	return errors.Wrap(err, "db/UpdateWithdrawSuccessful")
}

func  CreateWithdrawFunctions() error {
	_, err := PgDB.Exec(`



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
		v_fk_wallet_sender INT,
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
			fk_wallet_sender,
			fk_wallet_receiver,
			payment_cat,
			tx_internal_ref,
			value,
			time_tx )
		VALUES (
			v_fk_wallet_sender,
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
			id = v_fk_wallet_sender
		;

	RETURN wdr_id;

	END;
	$$;

	`)

	return errors.Wrap(err, "db/CreateWithdrawFunctions")
}

func  initWithdrawReqApply(wdr Withdraw, it InternalTx) (withdrawId int64, err error) {

	err = PgDB.QueryRow(`
		SELECT withdraw_req_init($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11);
		
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

	return withdrawId, errors.Wrap(err, "db/InitWithdrawReqApply")

}

func  InitWithdrawReq(walletId int64, value float64, extCurrencyAbbr string) (withdrawId int64, err error) {

	wdr := Withdraw{
		Value:      value,
		TxSentTime: time.Now().UTC(),
		TxStatus:   string(NOT_SENT_TO_PS),
	}

	wdr.FkExtAcntRcvr, err = GetUserExtAccountId(walletId, extCurrencyAbbr)
	if err != nil {
		return withdrawId, errors.Wrap(err, "db/InitWithdrawReq")
	}

	wdr.FkExtAcntSender, err = GetSuperNodeExtAccountId(extCurrencyAbbr)
	if err != nil {
		return withdrawId, errors.Wrap(err, "db/InitWithdrawReq")
	}

	wdr.FkExtCurr, err = GetExtCurrencyIdByAbbr(extCurrencyAbbr)
	if err != nil {
		return withdrawId, errors.Wrap(err, "db/InitWithdrawReq")
	}

	wdr.FkWithdrawFee, err = GetActiveWithdrawFeeId(extCurrencyAbbr)
	if err != nil {
		return withdrawId, errors.Wrap(err, "db/InitWithdrawReq")
	}

	var withdrawFeeAmnt float64
	withdrawFeeAmnt, err = GetActiveWithdrawFee(extCurrencyAbbr)
	if err != nil {
		return withdrawId, errors.Wrap(err, "db/InitWithdrawReq")
	}

	it := InternalTx{
		FkWalletSender: walletId,
		PaymentCat:     string(WITHDRAW),
		Value:          value + withdrawFeeAmnt,
	}

	it.FkWalletRcvr, err = GetWalletIdSuperNode()
	if err != nil {
		return withdrawId, errors.Wrap(err, "db/InitWithdrawReq")
	}

	return initWithdrawReqApply(wdr, it)

}

func  UpdateWithdrawPaymentQueryId(withdrawId int64, reqIdPaymentServ int64) error {
	_, err := PgDB.Exec(`
		UPDATE withdraw 
		SET
			tx_stat = 'PENDING',
			fk_query_id_payment_service = $1
		WHERE
			id = $2
		;
	
	`, reqIdPaymentServ,
		withdrawId)

	return errors.Wrap(err, "db/UpdateWithdrawPaymentQueryId")
}

func  GetWithdrawHist(walletId int64, offset int64, limit int64) ([]WithdrawHistRet, error) {

	rows, err := PgDB.Query(
		`SELECT
			ea.account_adr AS sender_adr, 
			ea2.account_adr AS receiver_adr, 
			ec.abv AS currency_abv,
			wdr.value,
			wf.fee,
			wdr.tx_sent_time,
			wdr.tx_stat,
			wdr.tx_approved_time,
			wdr.tx_hash
		FROM
			withdraw wdr,
			ext_account ea,
			ext_account ea2,
			wallet w, 
			ext_currency ec,
			withdraw_fee wf
		WHERE
			wdr.fk_ext_account_sender = ea.id AND
			wdr.fk_ext_account_receiver = ea2.id AND
			wdr.fk_ext_currency = ec.id AND
			ea2.fk_wallet = w.id AND
			ea.fk_ext_currency = ec.id AND
			ea2.fk_ext_currency = ec.id AND
			wdr.fk_withdraw_fee = wf.id AND
			w.id = $1
		ORDER BY wdr.tx_sent_time DESC
		LIMIT $2 
		OFFSET $3
		;`, walletId, limit, offset)

	defer rows.Close()

	res := make([]WithdrawHistRet, 0)
	var withVal WithdrawHistRet
	var aprvdTime, sentTime string

	for rows.Next() {
		rows.Scan(
			&withVal.AcntSender,
			&withVal.AcntRcvr,
			&withVal.ExtCurrency,
			&withVal.Value,
			&withVal.WithdrawFee,
			&sentTime,
			&withVal.TxStatus,
			&aprvdTime,
			&withVal.TxHash,
		)
		if conTime, errTime := time.Parse(timeLayout, sentTime); errTime == nil {
			withVal.TxSentTime = conTime
		} else {
			log.Debug("db/GetWithdrawHist Unable to convert time: ", err)
		}
		if conTime, errTime := time.Parse(timeLayout, aprvdTime); errTime == nil {
			withVal.TxAprvdTime = conTime
		} else {
			log.Debug("db/GetWithdrawHist Unable to convert time: ", err)
		}
		res = append(res, withVal)
	}
	return res, errors.Wrap(err, "db/GetWithdrawHist")
}

func  GetWithdrawHistRecCnt(walletId int64) (recCnt int64, err error) {

	err = PgDB.QueryRow(`
	SELECT
			COUNT(*)
		FROM
			withdraw wdr,
			ext_account ea2,
			wallet w			
		WHERE
			wdr.fk_ext_account_receiver = ea2.id AND
			ea2.fk_wallet = w.id AND
			w.id = $1		
	`, walletId).Scan(&recCnt)

	return recCnt, errors.Wrap(err, "db/GetWithdrawHistRecCnt")
}
