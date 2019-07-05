package postgres_db

import (
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type ExtCurrency struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Abv  string `db:"abv"`
}

func (pgDbp DbSpec) CreateExtCurrencyTable() error {
	_, err := pgDbp.Db.Exec(`
		CREATE TABLE IF NOT EXISTS 
		ext_currency (
			id SERIAL PRIMARY KEY,
			name VARCHAR(64),
			abv VARCHAR(16) UNIQUE NOT NULL
		);
	`)
	return errors.Wrap(err, "storage: query error CreateWalletTable()")
}

func (pgDbp DbSpec) InsertExtCurr(ec ExtCurrency) (insertIndex int, err error) {
	err = pgDbp.Db.QueryRow(`


	INSERT INTO ext_currency (
		name ,
		abv)
		VALUES (
		$1,
		$2
		)
		RETURNING id;
	`,
		ec.Name,
		ec.Abv).Scan(&insertIndex)

	// fmt.Println(val, err)
	return insertIndex, errors.Wrap(err, "storage: query error InsertExtCurr()")
}
