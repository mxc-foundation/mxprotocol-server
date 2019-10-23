package postgres_db

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type aggPeriodInterface struct{}

var PgAggPeriod aggPeriodInterface

func (*aggPeriodInterface) CreateAggPeriodTable() error {
	_, err := PgDB.Exec(`
	
		DO $$
			BEGIN
				IF NOT EXISTS 
					(SELECT 1 FROM pg_type WHERE typname = 'aggregation_status') 
				THEN
					CREATE TYPE aggregation_status AS ENUM (
						'IN_PROCESS',
						'COMPLETED'
			);
			END IF;

			CREATE TABLE IF NOT EXISTS agg_period (
				id SERIAL PRIMARY KEY,
				fk_dl_pkt_latest_id_accounted INT REFERENCES dl_pkt (id) NOT NULL,
				duration_minutes INT NOT NULL,
				status aggregation_status NOT NULL,
				execution_start_at   TIMESTAMP NOT NULL,
				execution_end_at TIMESTAMP
			);

			END$$;

`)

	return errors.Wrap(err, "db/pg_agg_period/CreateAggPeriodTable")
}

func (*aggPeriodInterface) InsertAggPeriod(latestIdAccountedDlPkt int64, durationMinutes int64, execStartAt time.Time) (insertInd int64, err error) {
	err = PgDB.QueryRow(`
		INSERT INTO agg_period 
			(
			fk_dl_pkt_latest_id_accounted,
			duration_minutes,
			status,
			execution_start_at
			)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
`,
		latestIdAccountedDlPkt,
		durationMinutes,
		types.AGG_IN_PROCESS,
		execStartAt,
	).Scan(&insertInd)
	return insertInd, errors.Wrap(err, "db/pg_agg_period/InsertAggPeriod")
}

func (*aggPeriodInterface) UpdateExecutedAggPeriod(aggPeriodId int64, execEndAt time.Time) (err error) {
	_, err = PgDB.Exec(`
		UPDATE agg_period 
			SET 
				execution_end_at = $1 , 
				status = $2
			WHERE
				id = $3
			;
		`,
		time.Now().UTC(),
		types.AGG_COMPLETED,
		aggPeriodId,
	)
	return errors.Wrap(err, fmt.Sprintf("db/pg_agg_period/UpdateExecutedAggPeriod aggPeriodId: %d", aggPeriodId))
}

func (*aggPeriodInterface) GetLatestAccountedDlPktId() (latestAccountedDlPktId int64, err error) {

	var cntRec int64
	err = PgDB.QueryRow(`
		SELECT
			count(*)
		FROM 
			agg_period 
	;
	`).Scan(&cntRec)

	if err != nil {
		return 0, errors.Wrap(err, "db/pg_agg_period/GetLatestAccountedDlPktId: Unable to get count of records! ")
	} else if cntRec == 0 {
		return 0, nil
	} else {

		var status string
		err = PgDB.QueryRow(`
		SELECT
			fk_dl_pkt_latest_id_accounted, 
			status
		FROM 
			agg_period 
		ORDER BY id DESC
		LIMIT 1
	;
	`).Scan(&latestAccountedDlPktId, &status)

		if err != nil {
			return 0, errors.Wrap(err, "db/pg_agg_period/GetLatestAccountedDlPktId")
		}

		if status == string(types.AGG_COMPLETED) {
			return latestAccountedDlPktId, nil
		} else {
			return 0, errors.New("db/pg_agg_period/GetLatestAccountedDlPktId: last Aggregation period is not completed!")

		}
	}

}
