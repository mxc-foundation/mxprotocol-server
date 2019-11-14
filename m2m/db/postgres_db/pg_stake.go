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

func (*stakeInterface) CreateStakeFunctions() error {
	_, err := PgDB.Exec(`

		CREATE OR REPLACE FUNCTION unstake_exec (
			v_fk_stake INT,
			v_time TIMESTAMP,
			v_fk_wallet_stake_storage INT,
			v_fk_wallet_user INT,
			v_payment_cat PAYMENT_CATEGORY,
			v_stake_status stake_status
		) RETURNS  VOID
		LANGUAGE plpgsql
		AS $$

		declare updated_wlt_balance NUMERIC(28,18);
		declare s_stake_amount NUMERIC(28,18);

		BEGIN


		SELECT 
			amount 
		INTO
			s_stake_amount 
		FROM
			stake
		WHERE 
			id = v_fk_stake 
		;

		UPDATE
			wallet 
		SET
			balance = balance + s_stake_amount,
			tmp_balance = tmp_balance + s_stake_amount
			
		WHERE
			id = v_fk_wallet_user
		RETURNING 
			balance INTO updated_wlt_balance
		;

		UPDATE
			wallet 
		SET
			balance = balance - s_stake_amount,
			tmp_balance = tmp_balance - s_stake_amount
			
		WHERE
			id = v_fk_wallet_stake_storage
		;


		UPDATE 
			stake 
		SET 
			status= v_stake_status, 
			unstake_time = v_time 
		WHERE 
			id = v_fk_stake
		;

		INSERT INTO internal_tx (
			fk_wallet_sender,
			fk_wallet_receiver,
			payment_cat,
			tx_internal_ref,
			value,
			time_tx )
		VALUES (
			v_fk_wallet_stake_storage,
			v_fk_wallet_user,
			v_payment_cat,
			v_fk_stake,
			s_stake_amount,
			v_time)
			;

		END;
		$$;
		`)
	return errors.Wrap(err, "db/pg_stake/CreateStakeFunctions")
}

func (*stakeInterface) InsertStake(walletId int64, amount float64) (insertIndex int64, err error) {
	err = PgDB.QueryRow(`
		INSERT INTO stake (
			fk_wallet ,
			amount ,
			status ,	
			start_stake_time , 
			unstake_time
			) 
		VALUES 
			($1,$2,$3,$4,$5)
		RETURNING id ;
	`,
		walletId,
		amount,
		types.STAKING_ACTIVE,
		time.Now().UTC(),
		time.Time{},
	).Scan(&insertIndex)
	return insertIndex, errors.Wrap(err, "db/pg_stake/InsertStake")
}

func (*stakeInterface) Unstake(stakeId int64) error {
	// first:  check if the status is not unstaked

	// return errors.Wrap(err, fmt.Sprintf("db/pg_stake/Unstake  stakeId: %d ", stakeId))
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

func (*stakeInterface) GetStakeProfile(stakeId int64) (stkPrf types.Stake, err error) {
	err = PgDB.QueryRow(`
		SELECT
			id,
			fk_wallet,
			amount,
			status,
			start_stake_time,
			unstake_time 
		FROM 
			stake
		WHERE
			id = $1
	;
	`,
		stakeId,
	).Scan(
		&stkPrf.Id,
		&stkPrf.FkWallet,
		&stkPrf.Amount,
		&stkPrf.Status,
		&stkPrf.StartStakeTime,
		&stkPrf.UnstakeTime)
	return stkPrf, errors.Wrap(err, "db/pg_stake/GetStakeProfile")

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
