package postgres_db

import (
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type extCurrencyInterface struct{}

var PgExtCurrency extCurrencyInterface

func (*extCurrencyInterface) CreateExtCurrencyTable() error {
	_, err := PgDB.Exec(`
		CREATE TABLE IF NOT EXISTS 
		ext_currency (
			id SERIAL PRIMARY KEY,
			name VARCHAR(64),
			abv VARCHAR(16) UNIQUE NOT NULL
		);
	`)
	return errors.Wrap(err, "db/CreateExtCurrencyTable")
}

func (*extCurrencyInterface) InsertExtCurr(ec types.ExtCurrency) (insertIndex int64, err error) {
	log.WithFields(log.Fields{
		"name": ec.Name,
		"abbr": ec.Abv,
	}).Info("/db/ext_currency_interface: insert ext_currency")
	err = PgDB.QueryRow(`
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
	return insertIndex, errors.Wrap(err, "db/InsertExtCurr")
}

func (*extCurrencyInterface) GetExtCurrencyIdByAbbr(extCurrencyAbbr string) (int64, error) {
	var extCurrencyId int64
	err := PgDB.QueryRow(`
		select id 
		from 
			ext_currency 
		where 
			abv=$1;`, extCurrencyAbbr).Scan(&extCurrencyId)

	return extCurrencyId, errors.Wrap(err, "db/GetExtCurrencyIdByAbbr")
}
