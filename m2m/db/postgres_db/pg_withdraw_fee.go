package postgres_db

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type withdrawFeeInterface struct{}

var PgWithdrawFee withdrawFeeInterface

type withdrawFee struct {
	Id         int64     `db:"id"`
	FkExtCurr  int64     `db:"fk_ext_currency"`
	Fee        float64   `db:"fee"`
	InsertTime time.Time `db:"insert_time"`
	Status     string    `db:"status"`
}

func (*withdrawFeeInterface) CreateWithdrawFeeTable() error {
	_, err := PgDB.Exec(`

		CREATE TABLE IF NOT EXISTS 
		withdraw_fee (
			id SERIAL PRIMARY KEY,
			fk_ext_currency INT REFERENCES ext_currency (id) NOT NULL,
			fee NUMERIC(28,18) NOT NULL,
			insert_time TIMESTAMP NOT NULL,
			status FIELD_STATUS NOT NULL
		);
		
	`)
	return errors.Wrap(err, "db/CreateWithdrawFeeTable")
}

func (*withdrawFeeInterface) InsertWithdrawFee(extCurrencyAbbr string, wdFee float64) (insertIndex int64, err error) {
	id, err := PgExtCurrency.GetExtCurrencyIdByAbbr(extCurrencyAbbr)
	if err != nil {
		return insertIndex, errors.Wrap(err, "db/InsertWithdrawFee")
	}

	wf := withdrawFee{
		FkExtCurr:  id,
		Fee:        wdFee,
		InsertTime: time.Now().UTC(),
		Status:     "ACTIVE",
	}
	err = PgDB.QueryRow(`
		INSERT INTO withdraw_fee (
			fk_ext_currency,
			fee,
			insert_time,
			status)
		VALUES (
			$1,
			$2,
			$3,
			$4
			)
		RETURNING id;
	`,
		wf.FkExtCurr,
		wf.Fee,
		wf.InsertTime,
		wf.Status,
	).Scan(&insertIndex)

	if err == nil {
		wf.Id = insertIndex
		err2 := changeStatus2ArcOldRowWithdrawFee(wf)
		if err2 != nil {
			return insertIndex, errors.Wrap(err, "db/InsertWithdrawFee")
		}
	}

	return insertIndex, errors.Wrap(err, "db/InsertWithdrawFee")
}

func changeStatus2ArcOldRowWithdrawFee(wf withdrawFee) (err error) {
	_, err = PgDB.Exec(`
	UPDATE 
		withdraw_fee 
	SET 
		status = 'ARC'
	WHERE
		fk_ext_currency = $1
		AND
		id <> $2   
	;
	`,
		wf.FkExtCurr,
		wf.Id)

	return errors.Wrap(err, "db/changeStatus2ArcOldRowWithdrawFee")
}

func (*withdrawFeeInterface) GetActiveWithdrawFee(extCurrAbv string) (withdrawFee float64, err error) {
	err = PgDB.QueryRow(`
		SELECT 
			wf.fee
		FROM
			withdraw_fee wf, ext_currency ec
		WHERE
			wf.fk_ext_currency = ec.id	AND
			ec.abv = $1 	AND
			status = 'ACTIVE'
		ORDER BY ec.id DESC 
		LIMIT 1 
		;
	`, extCurrAbv).Scan(&withdrawFee)

	return withdrawFee, errors.Wrap(err, "db/GetActiveWithdrawFee")
}

func (*withdrawFeeInterface) GetActiveWithdrawFeeId(extCurrAbv string) (withdrawFee int64, err error) {
	err = PgDB.QueryRow(`
		SELECT 
			wf.id
		FROM
			withdraw_fee wf, ext_currency ec
		WHERE
			wf.fk_ext_currency = ec.id	AND
			ec.abv = $1 	AND
			status = 'ACTIVE'
		ORDER BY ec.id DESC 
		LIMIT 1 
		;
	`,
		extCurrAbv).Scan(&withdrawFee)

	return withdrawFee, errors.Wrap(err, "db/GetActiveWithdrawFeeId")
}
