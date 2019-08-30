package postgres_db

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Device struct {
	id            int64     `db:"id"`
	DevEui        string    `db:"dev_eui"` // fk in App server
	FkWallet      int64     `db:"fk_wallet"`
	Mode          string    `db:"mode"`
	CreatedAt     time.Time `db:"created_at"`
	LastSeenAt    time.Time `db:"last_seen_at"`
	ApplicationId int64     `db:"application_id"`
	Name          string    `db:"name"`
}

func (pgDbp *PGHandler) CreateDeviceTable() error {
	_, err := pgDbp.DB.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS 
				(SELECT 1 FROM pg_type WHERE typname = 'device_mode') 
			THEN
				CREATE TYPE device_mode AS ENUM (
					'INACTIVE',
					 'FREE_GATEWAYS_LIMITED',
					 'WHOLE_NETWORK',
					 'DELETED'
		);
		END IF;
			CREATE TABLE IF NOT EXISTS device (
			id SERIAL PRIMARY KEY,
			dev_eui VARCHAR(64) NOT NULL,
			fk_wallet INT REFERENCES wallet (id) NOT NULL,
			mode DEVICE_MODE    NOT NULL,
			created_at     TIMESTAMP,
			last_seen_at    TIMESTAMP,
			application_id INT    ,
			name  varchar(128)    
		);

		END$$;
		
	`)
	return errors.Wrap(err, "db/CreateDevice")
}

func (pgDbp *PGHandler) InsertDevice(dv Device) (insertIndex int64, err error) {
	err = pgDbp.DB.QueryRow(`
		INSERT INTO device (
			dev_eui ,
			fk_wallet,
			mode,
			created_at,
			last_seen_at,
			application_id,
			name
			) 
		VALUES 
			($1,$2,$3,$4,$5,$6,$7)
		RETURNING id ;
	`,
		dv.DevEui,
		dv.FkWallet,
		dv.Mode,
		dv.CreatedAt,
		dv.LastSeenAt,
		dv.ApplicationId,
		dv.Name,
	).Scan(&insertIndex)
	return insertIndex, errors.Wrap(err, "db/InsertDevice")
}
