package postgres_db

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type WithdrawFee struct {
	Id         int64     `db:"id"`
	FkExtCurr  int64     `db:"fk_ext_currency"`
	Fee        float64   `db:"fee"`
	InsertTime time.Time `db:"insert_time"`
	Status     string    `db:"status"`
}

func (pgDbp DbSpec) CreateWithdrawFeeTable() error {
	_, err := pgDbp.Db.Exec(`

		CREATE TABLE IF NOT EXISTS 
		withdraw_fee (
			id SERIAL PRIMARY KEY,
			fk_ext_currency INT REFERENCES ext_currency (id) NOT NULL,
			fee NUMERIC(28,18) NOT NULL,
			insert_time TIMESTAMP NOT NULL,
			status FIELD_STATUS NOT NULL
		);
		
	`)
	return errors.Wrap(err, "db: query error CreateWalletTable()")
}

func (pgDbp DbSpec) InsertWithdrawFee(wf WithdrawFee) (insertIndex int64, err error) {
	err = pgDbp.Db.QueryRow(`
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

	return insertIndex, errors.Wrap(err, "db: query error InsertWithdrawFee()")
}

func (pgDbp DbSpec) GetActiveWithdrawFee(extCurrAbv string) (withdrawFee float64, err error) {
	err = pgDbp.Db.QueryRow(`
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

	if err != nil {
		err = errors.Wrap(err, "db/GetActiveWithdrawFee")
	}

	return withdrawFee, err
}

func (pgDbp DbSpec) GetActiveWithdrawFeeId(extCurrAbv string) (withdrawFee int64, err error) {
	err = pgDbp.Db.QueryRow(`
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

	return withdrawFee, errors.Wrap(err, "db: query error GetActiveWithdrawFeeId()")
}
