package postgres_db

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type ExtAccount struct {
	Id            int       `db:"id"`
	FkWallet      int       `db:"fk_wallet"`
	FkExtCurrency int       `db:"fk_ext_currency"`
	Account_adr   string    `db:"account_adr"`
	Insert_time   time.Time `db:"insert_time"`
	Status        string    `db:"status"`
}

func (pgDbp DbSpec) CreateExtAccountTable() error {
	_, err := pgDbp.Db.Exec(`
		CREATE TABLE IF NOT EXISTS ext_account (
			id SERIAL PRIMARY KEY,
			fk_wallet INT REFERENCES wallet(id) NOT NULL,
			fk_ext_currency INT REFERENCES ext_currency (id) NOT NULL,
			account_adr varchar(128) NOT NULL,
			insert_time TIMESTAMP NOT NULL,
			status FIELD_STATUS NOT NULL
		);
		
	`)
	return errors.Wrap(err, "db: query error CreateExtAccountTable()")
}

func (pgDbp DbSpec) InsertExtAccount(ea *ExtAccount) error {
	err := pgDbp.Db.QueryRow(`
	INSERT INTO ext_account (
		fk_wallet,
		fk_ext_currency,
		account_adr,
		insert_time,
		status)
		VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
		)
		RETURNING id;
	`,
		ea.FkWallet,
		ea.FkExtCurrency,
		ea.Account_adr,
		ea.Insert_time,
		ea.Status,
	).Scan(&ea.Id)

	return errors.Wrap(err, "db: query error InsertWithdrawFee()")
}
