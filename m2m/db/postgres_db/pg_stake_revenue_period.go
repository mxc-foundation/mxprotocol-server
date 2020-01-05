package postgres_db

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/mxc-foundation/mxprotocol-server/m2m/types"
)

type stakeRevenuePeriodInterface struct{}

var PgStakeRevenuePeriod stakeRevenuePeriodInterface

func (*stakeRevenuePeriodInterface) CreateStakeRevenuePeriodTable() error {
	_, err := PgDB.Exec(

		`DO $$
		BEGIN
			IF NOT EXISTS 
				(SELECT 1 FROM pg_type WHERE typname = 'stake_revenue_period_status') 
			THEN
				CREATE TYPE stake_revenue_period_status AS ENUM (
					'IN_PROCESS',
 					'COMPLETED'
		);
		END IF;
		CREATE TABLE IF NOT EXISTS stake_revenue_period (
			id SERIAL PRIMARY KEY,
			staking_period_start TIMESTAMP NOT NULL,
			staking_period_end TIMESTAMP NOT NULL,
			supernode_income NUMERIC(28,18) NOT NULL,
			income_to_stake_portion FLOAT NOT NULL  CHECK (income_to_stake_portion >= 0 AND income_to_stake_portion <= 1),
			exec_start_time TIMESTAMP,
			exec_end_time TIMESTAMP,
			status stake_revenue_period_status
		);	

		END$$;
		
	`)
	return errors.Wrap(err, "db/pg_stake_revenue_period/CreateStakeRevenuePeriodTable")

}

func (*stakeRevenuePeriodInterface) InsertStakeRevenuePeriod(stakingPeriodStart time.Time, stakingPeriodEnd time.Time, superNodeIncome float64, incomeToStakePortion float64) (insertIndex int64, err error) {
	err = PgDB.QueryRow(`
		INSERT INTO stake_revenue_period (
			staking_period_start ,
			staking_period_end ,
			supernode_income ,	
			income_to_stake_portion,
			exec_start_time,
			status
			) 
		VALUES 
			($1,$2,$3,$4,$5,$6)
		RETURNING id ;
	`,
		stakingPeriodStart,
		stakingPeriodEnd,
		superNodeIncome,
		incomeToStakePortion,
		time.Now().UTC(),
		types.STAKE_REVENUE_IN_PROCESS,
	).Scan(&insertIndex)
	return insertIndex, errors.Wrap(err, "db/pg_stake_revenue_period/InsertStakeRevenuePeriod")
}

func (*stakeRevenuePeriodInterface) UpdateCompletedStakeRevenuePeriod(stakeRevenuePeriodId int64) error {
	_, err := PgDB.Exec(` 
	UPDATE 
		stake_revenue_period
	SET
		status = $1 ,
		exec_end_time = $2
	WHERE
		id = $3
	;`, types.STAKE_REVENUE_COMPLETED,
		time.Now().UTC(),
		stakeRevenuePeriodId)
	return errors.Wrap(err, fmt.Sprintf("db/pg_stake_revenue_period/UpdateCompletedStakeRevenuePeriod   stakeRevenuePeriodId: %d", stakeRevenuePeriodId))
}
