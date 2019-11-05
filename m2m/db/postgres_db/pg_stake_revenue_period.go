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
				(SELECT 1 FROM pg_type WHERE typname = 'STAKE_REVENUE_PERIOD_STATUS') 
			THEN
				CREATE TYPE STAKE_REVENUE_PERIOD_STATUS AS ENUM (
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
			status STAKE_REVENUE_PERIOD_STATUS
		);	

		END$$;
		
	`)
	return errors.Wrap(err, "db/pg_stake_revenue_period/CreateStakeRevenuePeriodTable")

}
