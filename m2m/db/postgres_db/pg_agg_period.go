package postgres_db

import (
	"time"

	"github.com/pkg/errors"
)

type aggPeriodInterface struct{}

var PgAggPeriod aggPeriodInterface

func (*aggPeriodInterface) CreateAggPeriodTable() error {
	_, err := PgDB.Exec(`
	
	CREATE TABLE IF NOT EXISTS agg_period (
		id SERIAL PRIMARY KEY,
		start_at  TIMESTAMP NOT NULL,
		duration_minutes INT NOT NULL,
		execution_time  TIMESTAMP 
	);
	
`)

	return errors.Wrap(err, "db/pg_agg_period/CreateAggPeriodTable")
}

func (*aggPeriodInterface) InsertAggPeriod(startAt time.Time, durationMinutes int64, executionTime time.Time) (insertInd int64, err error) {
	err = PgDB.QueryRow(`
	INSERT INTO agg_period 
		(start_at,
		duration_minutes,
		execution_time)
	VALUES ($1, $2,$3)
	RETURNING id;
`,
		startAt,
		durationMinutes,
		executionTime,
	).Scan(&insertInd)
	return insertInd, errors.Wrap(err, "db/pg_agg_period/InsertAggPeriod")
}
