package postgres_db

import (
	"time"

	"github.com/ethereum/go-ethereum/log"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Topup struct {
	Id              int64     `db:"id"`
	FkExtAcntSender int64     `db:"fk_ext_account_sender"`
	FkExtAcntRcvr   int64     `db:"fk_ext_account_receiver"`
	FkExtCurr       int64     `db:"fk_ext_currency"`
	Value           float64   `db:"value"`
	TxAprvdTime     time.Time `db:"tx_approved_time"`
	TxHash          string    `db:"tx_hash"`
}

type TopupHistRet struct {
	AcntSender  string
	AcntRcvr    string
	ExtCurrency string
	Value       float64
	TxAprvdTime time.Time
	TxHash      string
}

func (pgDbp DbSpec) CreateTopupTable() error {
	_, err := pgDbp.Db.Exec(`
	CREATE TABLE IF NOT EXISTS topup (
		id SERIAL PRIMARY KEY,
		fk_ext_account_sender INT REFERENCES  ext_account(id) NOT NULL,
		fk_ext_account_receiver INT REFERENCES  ext_account(id) NOT NULL,
		fk_ext_currency INT REFERENCES ext_currency(id) NOT NULL,
		value NUMERIC(28,18) NOT NULL,
		tx_approved_time TIMESTAMP,
		tx_hash varchar (128) NOT NULL UNIQUE
		);
	`)
	return errors.Wrap(err, "db/CreateTopupTable")
}

func (pgDbp DbSpec) insertTopup(tu Topup) (insertIndex int64, err error) {
	err = pgDbp.Db.QueryRow(`
		INSERT INTO topup (
			fk_ext_account_sender,
			fk_ext_account_receiver,
			fk_ext_currency,
			value ,
			tx_approved_time,
			tx_hash )
		VALUES (
			$1,	$2,	$3,	$4, $5, $6
		)
		RETURNING id;
		
	`,
		tu.FkExtAcntSender,
		tu.FkExtAcntRcvr,
		tu.FkExtCurr,
		tu.Value,
		tu.TxAprvdTime,
		tu.TxHash).Scan(&insertIndex)

	return insertIndex, errors.Wrap(err, "db/InsertTopup")
}

func (pgDbp DbSpec) CreateTopupFunctions() error {
	_, err := pgDbp.Db.Exec(`

	CREATE OR REPLACE FUNCTION topup_req_apply (
			v_fk_ext_account_sender INT,
			v_fk_ext_account_receiver INT,
			v_fk_ext_currency INT,
			v_value NUMERIC(28,18),
			v_tx_approved_time TIMESTAMP,
			v_tx_hash VARCHAR(128),
			v_fk_wallet_sender INT,
			v_fk_wallet_receiver INT,
			v_payment_cat PAYMENT_CATEGORY
		) RETURNS INT
	LANGUAGE plpgsql
	AS $$
		
		declare topup_id INT;
		
		BEGIN
		
		INSERT INTO topup (
			fk_ext_account_sender,
			fk_ext_account_receiver,
			fk_ext_currency,
			value,
			tx_approved_time,
			tx_hash )
			VALUES (
			v_fk_ext_account_sender ,
			v_fk_ext_account_receiver,
			v_fk_ext_currency,
			v_value,
			v_tx_approved_time,
			v_tx_hash
		)RETURNING id INTO topup_id;
		
		
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
			topup_id,
			v_value,
			v_tx_approved_time)
			;
		
		
		UPDATE
			wallet 
		SET
			balance = balance + v_value
		WHERE
			id = v_fk_wallet_receiver
		;
		
		RETURN topup_id;
		
		END;
		$$;
		

	`)

	return errors.Wrap(err, "db/CreateTopupFunctions")
}

func (pgDbp DbSpec) applyTopup(tu Topup, it InternalTx) (topupId int64, err error) {
	err = pgDbp.Db.QueryRow(`
		select topup_req_apply($1,$2,$3,$4,$5,$6,$7,$8,$9);
		
	`, tu.FkExtAcntSender,
		tu.FkExtAcntRcvr,
		tu.FkExtCurr,
		tu.Value,
		tu.TxAprvdTime,
		tu.TxHash,
		it.FkWalletSender,
		it.FkWalletRcvr,
		it.PaymentCat).Scan(&topupId)

	return topupId, errors.Wrap(err, "db/ApplyTopup")

}

func (pgDbp DbSpec) AddTopUpRequest(acntAdrSender string, acntAdrRcvr string, txHash string, value float64, extCurAbv string) (topupId int64, err error) {

	tu := Topup{
		Value:       value,
		TxAprvdTime: time.Now().UTC(),
		TxHash:      txHash,
	}

	tu.FkExtAcntSender, err = pgDbp.GetExtAccountIdByAdr(acntAdrSender)
	if err != nil {
		return topupId, errors.Wrap(err, "db/AddTopUpRequest")
	}

	tu.FkExtAcntRcvr, err = pgDbp.GetExtAccountIdByAdr(acntAdrRcvr)
	if err != nil {
		return topupId, errors.Wrap(err, "db/AddTopUpRequest")
	}

	tu.FkExtCurr, err = pgDbp.GetExtCurrencyIdByAbbr(extCurAbv)
	if err != nil {
		return topupId, errors.Wrap(err, "db/AddTopUpRequest")
	}

	it := InternalTx{
		PaymentCat: string(TOP_UP),
	}

	it.FkWalletRcvr, err = pgDbp.GetWalletIdofActiveAcnt(acntAdrSender, extCurAbv)
	if err != nil {
		return topupId, errors.Wrap(err, "db/AddTopUpRequest")
	}

	it.FkWalletSender, err = pgDbp.getWalletIdofActiveAcntSuperAdmin(acntAdrRcvr, extCurAbv)
	if err != nil {
		return topupId, errors.Wrap(err, "db/AddTopUpRequest")
	}

	return pgDbp.applyTopup(tu, it)

}

func (pgDbp DbSpec) GetTopupHist(walletId int64, offset int64, limit int64) ([]TopupHistRet, error) {

	rows, err := pgDbp.Db.Query(
		`select
			ea.account_adr AS sender_adr, 
			ea2.account_adr AS receiver_adr, 
			ec.abv AS currency_abv,
			tu.value,
			tu.tx_approved_time,
			tu.tx_hash
		from
			topup tu,
			ext_account ea,
			ext_account ea2,
			wallet w, 
			ext_currency ec
		WHERE
			tu.fk_ext_account_sender = ea.id AND
			tu.fk_ext_account_receiver = ea2.id AND
			tu.fk_ext_currency = ec.id AND
			ea.fk_wallet = w.id AND
			ea.fk_ext_currency = ec.id AND
			ea2.fk_ext_currency = ec.id AND
			w.id = $1 
		ORDER BY tu.tx_approved_time DESC
		LIMIT $2  
		OFFSET $3 
		;`, walletId, limit, offset)

	defer rows.Close()

	res := make([]TopupHistRet, 0)
	var topupVal TopupHistRet
	var timeRead string

	for rows.Next() {
		rows.Scan(
			&topupVal.AcntSender,
			&topupVal.AcntRcvr,
			&topupVal.ExtCurrency,
			&topupVal.Value,
			&timeRead,
			&topupVal.TxHash,
		)
		if conTime, errTime := time.Parse(timeLayout, timeRead); errTime == nil {
			topupVal.TxAprvdTime = conTime
		} else {
			log.Debug("db/GetTopupHist Unable to convert time: ", err)
		}
		res = append(res, topupVal)
	}
	return res, errors.Wrap(err, "db/GetTopupHist")
}
