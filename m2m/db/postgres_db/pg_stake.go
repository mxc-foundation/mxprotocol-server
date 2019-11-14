package postgres_db

import (
	"time"

	"github.com/pkg/errors"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

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
 					'UNSTAKED'
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

func (*stakeInterface) InsertStake(walletId int64, amount float64) (insertIndex int64, err error) {
	err = PgDB.QueryRow(`
		INSERT INTO stake (
			fk_wallet ,
			amount ,
			status ,	
			start_stake_time
			) 
		VALUES 
			($1,$2,$3,$4)
		RETURNING id ;
	`,
		walletId,
		amount,
		types.STAKING_ACTIVE,
		time.Now().UTC(),
	).Scan(&insertIndex)
	return insertIndex, errors.Wrap(err, "db/pg_stake/InsertStake")
}

func (*stakeInterface) Unstake(stakeId int64) error {
	// TODO
	// in a single operation:
	// UPDATE status, unstake_time
	// update balance (user, stake)
	// make internal_tx

	return nil
}

func (*stakeInterface) GetStakeWalletId(stakeId int64) (walletId int64, err error) {
	err = PgDB.QueryRow(`
		SELECT
			fk_wallet
		FROM 
			stake
		WHERE
			id = $1
	;
	`,
		stakeId,
	).Scan(&walletId)
	return walletId, errors.Wrap(err, "db/pg_stake/GetStakeWalletId")
}

func (*stakeInterface) GetActiveStake(walletId int64) (stakeProfile types.Stake, err error) {

	err = PgDB.QueryRow(
		`SELECT
			id, fk_wallet, amount, status, start_stake_time
		FROM
			stake 
		WHERE
			fk_wallet = $1 
		AND
			status = 'ACTIVE'
		ORDER BY id DESC 
		LIMIT 1  
		;`, walletId).Scan(
		&stakeProfile.Id,
		&stakeProfile.FkWallet,
		&stakeProfile.Amount,
		&stakeProfile.Status,
		&stakeProfile.StartStakeTime)
	return stakeProfile, errors.Wrap(err, "db/pg_stake/GetActiveStake")

}

func (*stakeInterface) GetActiveStakes() (stakeProfiles []types.Stake, err error) {

	rows, err := PgDB.Query(
		`SELECT
			id, fk_wallet, amount, status, start_stake_time
		FROM
			stake 
		WHERE
			status = 'ACTIVE'
	;`)

	defer rows.Close()

	stakePrf := types.Stake{}

	for rows.Next() {
		rows.Scan(
			&stakePrf.Id,
			&stakePrf.FkWallet,
			&stakePrf.Amount,
			&stakePrf.Status,
			&stakePrf.StartStakeTime,
		)

		stakeProfiles = append(stakeProfiles, stakePrf)
	}
	return stakeProfiles, errors.Wrap(err, "db/pg_stake/GetActiveStakes")

}
