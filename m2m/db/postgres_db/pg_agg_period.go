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
		fk_dl_pkt_latest_id_accounted INT REFERENCES dl_pkt (id) NOT NULL,
		duration_minutes INT NOT NULL,
		execution_start_at   TIMESTAMP NOT NULL,
	    execution_end_at TIMESTAMP
	);
	
`)

	return errors.Wrap(err, "db/pg_agg_period/CreateAggPeriodTable")
}

func (*aggPeriodInterface) InsertAggPeriod(latestIdAccountedDlPkt int64, durationMinutes int64, execStartAt time.Time) (insertInd int64, err error) {
	err = PgDB.QueryRow(`
	INSERT INTO agg_period 
		(
		fk_dl_pkt_latest_id_accounted,
		duration_minutes,
		execution_start_at
		)
	VALUES ($1, $2, $3, $4)
	RETURNING id;
`,
		latestIdAccountedDlPkt,
		durationMinutes,
		execStartAt,
	).Scan(&insertInd)
	return insertInd, errors.Wrap(err, "db/pg_agg_period/InsertAggPeriod")
}
