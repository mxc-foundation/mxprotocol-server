package postgres_db

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type WithdrawFee struct {
	Id         int       `db:"id"`
	FkExtCurr  int       `db:"fk_ext_currency"`
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

func (pgDbp DbSpec) InsertWithdrawFee(wf *WithdrawFee) error {
	err := pgDbp.Db.QueryRow(`
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
	).Scan(&wf.Id)

	return errors.Wrap(err, "db: query error InsertWithdrawFee()")
}
