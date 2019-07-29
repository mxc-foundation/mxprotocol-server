package postgres_db

import (
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Wallet struct {
	Id      int64   `db:"id"`
	FkOrgLa int64   `db:"fk_org_la"`
	TypeW   string  `db:"type"`
	Balance float64 `db:"balance"`
}

func (pgDbp DbSpec) CreateWalletTable() error {
	_, err := pgDbp.Db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS 
				(SELECT 1 FROM pg_type WHERE typname = 'wallet_type') 
			THEN
				CREATE TYPE WALLET_TYPE AS ENUM (
					'SUPER_ADMIN',
 					'USER'
		);
		END IF;
			CREATE TABLE IF NOT EXISTS wallet (
	
			id SERIAL PRIMARY KEY,
			fk_org_la INT UNIQUE NOT NULL, -- foreign_key LoRa app server DB
			type WALLET_TYPE NOT NULL,
			balance NUMERIC(28,18) NOT NULL CHECK (balance >= 0)
		);

		END$$;
		
	`)
	return errors.Wrap(err, "db/CreateWalletTable")
}

func (pgDbp DbSpec) InsertWallet(w Wallet) (insertIndex int64, err error) {
	err = pgDbp.Db.QueryRow(`
		INSERT INTO wallet (
			fk_org_la ,
			type,
			balance ) 
		VALUES 
			($1,$2,$3)
		RETURNING id ;
	`,
		w.FkOrgLa,
		w.TypeW,
		w.Balance).Scan(&insertIndex)

	// fmt.Println(val, err)
	return insertIndex, errors.Wrap(err, "db/InsertWallet")
}

func (pgDbp DbSpec) GetWalletIdFromOrgId(orgIdLora int64) (int64, error) {
	var w Wallet
	w.Id = 0
	err := pgDbp.Db.QueryRow(
		`SELECT id
		FROM wallet
		WHERE
			fk_org_la = $1;`,
		orgIdLora).Scan(&w.Id)

	return w.Id, errors.Wrap(err, "db/GetWalletIdFromOrgId")
}

func (pgDbp DbSpec) GetWalletBalance(walletId int64) (float64, error) {
	var w Wallet
	w.Id = 0
	err := pgDbp.Db.QueryRow(
		`SELECT balance
		FROM wallet
		WHERE
			id = $1;`,
		walletId).Scan(&w.Balance)

	return w.Balance, errors.Wrap(err, "db/GetWalletBalance")
}

func (pgDbp DbSpec) GetWalletIdofActiveAcnt(acntAdr string, externalCur string) (walletId int64, err error) {

	err = pgDbp.Db.QueryRow(
		`select 
			w.id as wallet_id 
			from
			wallet w,ext_account ea,ext_currency ec
		WHERE
			w.id = ea.fk_wallet AND
			w.type = 'USER' AND
			ea.fk_ext_currency = ec.id AND
			ea.status = 'ACTIVE' AND
			account_adr = $1 AND
			ec.abv = $2 
		ORDER BY ea.id DESC 
		LIMIT 1 
		;`, acntAdr, externalCur).Scan(&walletId)

	return walletId, errors.Wrap(err, "db/GetWalletIdofActiveAcnt")
}

func (pgDbp DbSpec) getWalletIdofActiveAcntSuperAdmin(acntAdr string, externalCur string) (walletId int64, err error) {

	err = pgDbp.Db.QueryRow(
		`select 
			w.id as wallet_id 
			from
			wallet w,ext_account ea,ext_currency ec
		WHERE
			w.id = ea.fk_wallet AND
			w.type = 'SUPER_ADMIN' AND
			ea.fk_ext_currency = ec.id AND
			ea.status = 'ACTIVE' AND
			account_adr = $1 AND
			ec.abv = $2 
		ORDER BY ea.id DESC 
		LIMIT 1 
		;`, acntAdr, externalCur).Scan(&walletId)

	return walletId, errors.Wrap(err, "db/getWalletIdofActiveAcntSuperAdmin")
}

func (pgDbp DbSpec) GetWalletIdSuperNode() (walletId int64, err error) {

	err = pgDbp.Db.QueryRow(
		`SELECT
			id
		FROM
			wallet 
		WHERE
			type = 'SUPER_ADMIN' 
			ORDER BY id DESC 
			LIMIT 1  -- altough only one super node exists
		;`).Scan(&walletId)

	return walletId, errors.Wrap(err, "db/GetWalletIdSuperNode")
}

func (pgDbp DbSpec) UpdateBalanceByWalletId(walletId int64, newBalance float64) error {
	_, err := pgDbp.Db.Exec(`
	UPDATE
		wallet 
	SET
		balance = $1
	WHERE
		id = $2
	;
	`, newBalance, walletId)

	return errors.Wrap(err, "db/UpdateBalanceByWalletId")
}
