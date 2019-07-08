package postgres_db

import (
	"github.com/apex/log"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type ExtAccount struct {
	Id                 int64     `db:"id"`
	FkWallet           int64     `db:"fk_wallet"`
	FkExtCurrency      int64     `db:"fk_ext_currency"`
	Account_adr        string    `db:"account_adr"`
	Insert_time        time.Time `db:"insert_time"`
	Status             string    `db:"status"`
	LatestCheckedBlock int       `db:"latest_checked_block"`
}

func (pgDbp DbSpec) CreateExtAccountTable() error {
	_, err := pgDbp.Db.Exec(`
		CREATE TABLE IF NOT EXISTS ext_account (
			id SERIAL PRIMARY KEY,
			fk_wallet INT REFERENCES wallet(id) NOT NULL,
			fk_ext_currency INT REFERENCES ext_currency (id) NOT NULL,
			account_adr varchar(128) NOT NULL UNIQUE,
			insert_time TIMESTAMP NOT NULL,
			status FIELD_STATUS NOT NULL,
			latest_checked_block INT DEFAULT 0
		);
		
	`)
	return errors.Wrap(err, "db: query error CreateExtAccountTable()")
}

func (pgDbp DbSpec) InsertExtAccount(ea ExtAccount) (insertIndex int, err error) {
	err = pgDbp.Db.QueryRow(`
	INSERT INTO ext_account (
			fk_wallet,
			fk_ext_currency,
			account_adr,
			insert_time,
			status,
			latest_checked_block)
		VALUES (
			$1, $2,	$3,	$4,	'ACTIVE',	$5
		)
		RETURNING id;
	`,
		ea.FkWallet,
		ea.FkExtCurrency,
		ea.Account_adr,
		ea.Insert_time,
		ea.LatestCheckedBlock).Scan(&insertIndex)

	if err != nil {
		log.WithError(err).Error("/db/InsertExtAccount")
	}
	return insertIndex, err
}

func (pgDbp DbSpec) GetSuperNodeExtAccountAdr(extCurrAbv string) (string, error) {

	var res string

	err := pgDbp.Db.QueryRow(`
		select 
			account_adr
		from
			wallet w ,ext_account ea,ext_currency ec
		WHERE
			w.id = ea.fk_wallet
			AND
			w.type = 'SUPER_ADMIN' -- @change 
			AND
			ea.fk_ext_currency = ec.id
			AND
			ea.status = 'ACTIVE'
			AND
			ec.abv = $1
		ORDER BY ea.id DESC  
		LIMIT 1 
		;
	`, extCurrAbv).Scan(&res)

	return res, errors.Wrap(err, "db: query error GetSuperNodeExtAccountAdr()")
}

func (pgDbp DbSpec) GetUserExtAccountAdr(walletId int, extCurrAbv string) (string, error) {

	var res string

	err := pgDbp.Db.QueryRow(`
		select 
			account_adr
		from
			wallet w,ext_account ea,ext_currency ec
		WHERE
			w.id = ea.fk_wallet AND
			w.type = 'USER' AND
			ea.fk_ext_currency = ec.id AND
			ea.status = 'ACTIVE' AND
			w.id = $1 AND
			ec.abv = $2
		ORDER BY ea.id DESC 
		LIMIT 1 
		;
	
	`, walletId, extCurrAbv).Scan(&res)

	return res, errors.Wrap(err, "db: query error GetUserExtAccountAdr()")
}

func (pgDbp DbSpec) GetLatestCheckedBlock(extAcntId int) (int, error) {

	var res int

	err := pgDbp.Db.QueryRow(`
		SELECT 
			latest_checked_block 
		FROM 
			 ext_account 
		WHERE
			 id = $1
	
	`, extAcntId).Scan(&res)

	return res, errors.Wrap(err, "db: query error GetLatestCheckedBlock()")
}

func (pgDbp DbSpec) UpdateLatestCheckedBlock(extAcntId int, updatedBlockNum int) error {

	_, err := pgDbp.Db.Exec(`
		UPDATE ext_account 
		SET 
		latest_checked_block = $1
		WHERE
		id = $2;
	
	`, updatedBlockNum, extAcntId)

	return errors.Wrap(err, "db: query error InsertWithdrawFee()")
}
