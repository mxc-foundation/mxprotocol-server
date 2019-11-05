package postgres_db

import (
	"github.com/pkg/errors"
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
			income_to_stake_percentage FLOAT NOT NULL,
			exec_start_time TIMESTAMP,
			exec_end_time TIMESTAMP,
			status stake_revenue_period_status
		);	

		END$$;
		
	`)
	return errors.Wrap(err, "db/pg_stake_revenue_period/CreateStakeRevenuePeriodTable")

}
