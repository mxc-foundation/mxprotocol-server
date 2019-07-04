package postgres_db

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type WalletType string // db:wallet_type

const (
	USER        WalletType = "USER"
	SUPER_ADMIN WalletType = "SUPER_ADMIN"
)

type Wallet struct {
	Id      int     `db:"id"`
	FkOrgLa int     `db:"fk_org_la"`
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
			balance NUMERIC(28,18) NOT NULL   
		);

		END$$;
		
	`)
	return errors.Wrap(err, "db: query error CreateWalletTable()")
}

func (pgDbp DbSpec) InsertWallet(w Wallet) error {
	_, err := pgDbp.Db.Exec(`
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
		w.Balance)

	// fmt.Println(val, err)
	return errors.Wrap(err, "db: query error InsertWallet()")
}

func (pgDbp DbSpec) GetWalletIdFromOrgId(orgIdLora int) (int, error) {
	var w Wallet
	w.Id = 0
	query := pgDbp.Db.QueryRow(
		`SELECT id
		FROM wallet
		WHERE
			fk_org_la = $1;`,
		orgIdLora)

	err := query.Scan(
		&w.Id,
	)
	if err != nil {
		// fmt.Println(err)
		log.WithError(err).Warning("db: query error GetWalletIdFromOrgId()")

	}
	return w.Id, err
}

func (pgDbp DbSpec) GetWallet(wp *Wallet, walletId int) error {

	query := pgDbp.Db.QueryRow(
		`SELECT *
		FROM wallet
		WHERE
			id = $1;`,
		walletId)

	err := query.Scan(
		&wp.Id,
		&wp.FkOrgLa,
		&wp.TypeW,
		&wp.Balance,
	)

	if err != nil {
		fmt.Println("error getWallet: ", err)
		fmt.Println("query res: ", query)
		log.WithError(err).Warning("db:  query error GetWallet()") //@@ should be changed

	}

	return err
}

func (pgDbp DbSpec) GetWalletBalance(walletId int) (float64, error) {
	var w Wallet
	w.Id = 0
	query := pgDbp.Db.QueryRow(
		`SELECT balance
		FROM wallet
		WHERE
			id = $1;`,
		walletId)

	err := query.Scan(
		&w.Balance,
	)
	if err != nil {
		fmt.Println("GetWalletBalance error. wallet id: "+string(walletId)+"error: ", err)
		log.WithError(err).Warning("GetWalletBalance error. wallet id: "+string(walletId)+"error: ", err)
	}
	return w.Balance, err
}
