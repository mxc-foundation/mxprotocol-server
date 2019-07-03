package postgres_db

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	USER        = "USER"
	SUPER_ADMIN = "SUPER_ADMIN"
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
		END$$;

		CREATE TABLE IF NOT EXISTS wallet (
			id SERIAL PRIMARY KEY,
			fk_org_la INT UNIQUE NOT NULL, -- foreign_key LoRa app server DB
			type WALLET_TYPE NOT NULL,
			balance NUMERIC(28,18) NOT NULL   
		);
	`)
	return errors.Wrap(err, "storage: PostgreSQL connection error")
}

func (pgDbp DbSpec) InsertWallet(w Wallet) error {
	val, err := pgDbp.Db.Exec(`
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

	fmt.Println(val, err)
	return errors.Wrap(err, "storage: PostgreSQL connection error")
}

func (pgDbp DbSpec) GetWalletId(orgIdLora int) int {

	var w Wallet
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
		fmt.Println(err)
		log.WithError(err).Warning("storage: ping PostgreSQL database error, will retry in 2s")
		// log.WithError(err).Warning("storage: ping PostgreSQL database error, will retry in 2s")

	}

	// fmt.Println("errSelect:", errIns)

	if w.Id > 0 {
		return w.Id
	} else {
		return 0
	}

}

func (pgDbp DbSpec) GetWallet(wp *Wallet, orgIdLora int) error {

	query := pgDbp.Db.QueryRow(
		`SELECT *
		FROM wallet
		WHERE
			fk_org_la = $1;`,
		orgIdLora)

	err := query.Scan(
		wp.Id,
		wp.FkOrgLa,
		wp.TypeW,
		wp.Balance,
	)

	if err != nil {
		fmt.Println("error getWallet: ", err)
		fmt.Println("query res: ", query)
		log.WithError(err).Warning("storage: ping PostgreSQL database error, will retry in 2s")

	}

	return err

}
