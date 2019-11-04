package postgres_db

import "github.com/pkg/errors"

type stakeInterface struct{}

var PgStake stakeInterface

func (*stakeInterface) CreateStakeTable() error {
	_, err := PgDB.Exec(

		`DO $$
		BEGIN
			IF NOT EXISTS 
				(SELECT 1 FROM pg_type WHERE typname = 'stake_status') 
			THEN
				CREATE TYPE stake_status AS ENUM (
					'ACTIVE',
 					'ARC'
		);
		END IF;
		CREATE TABLE IF NOT EXISTS stake (
			id SERIAL PRIMARY KEY,
			fk_wallet INT REFERENCES wallet (id) NOT NULL,
			amount NUMERIC(28,18) DEFAULT 0 NOT NULL  CHECK (amount >= 0),
			status  stake_status NOT NULL,	
			start_stake_time  TIMESTAMP NOT NULL,
			unstake_time  TIMESTAMP
		);	

		END$$;
		
	`)
	return errors.Wrap(err, "db/pg_stake/CreateStakeTable")

}
