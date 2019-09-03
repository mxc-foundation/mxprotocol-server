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
	return errors.Wrap(err, "db/pg_device/CreateDevice")
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
	return insertIndex, errors.Wrap(err, "db/pg_device/InsertDevice")
}

func (pgDbp *PGHandler) GetDeviceMode(dvId int64) (dvMode string, err error) {
	err = pgDbp.DB.QueryRow(
		`SELECT
			mode
		FROM
			device 
		WHERE
			id = $1 
		;`, dvId).Scan(&dvMode)
	return dvMode, errors.Wrap(err, "db/pg_device/GetDeviceMode")
}

func (pgDbp *PGHandler) SetDeviceMode(dvId int64, dvMode string) (err error) {
	_, err = pgDbp.DB.Exec(`
		UPDATE
			device 
		SET
			mode = $1
		WHERE
			id = $2
		;
		`, dvMode, dvId)

	return errors.Wrap(err, "db/pg_device/SetDeviceMode")

}

func (pgDbp *PGHandler) GetDeviceIdByDevEui(devEui string) (devId int64, err error) {
	err = pgDbp.DB.QueryRow(
		`SELECT
			id
		FROM
			device 
		WHERE
			dev_eui = $1 
			ORDER BY id DESC 
			LIMIT 1  
		;`, devEui).Scan(&devId)
	return devId, errors.Wrap(err, "db/pg_device/GetDeviceIdByDevEui")
}

func (pgDbp *PGHandler) UpdateDeviceLastSeen(dvId int64, newTime time.Time) (err error) {
	_, err = pgDbp.DB.Exec(`
		UPDATE
			device 
		SET
			last_seen_at = $1
		WHERE
			id = $2
		;
		`, newTime, dvId)

	return errors.Wrap(err, "db/pg_device/UpdateDeviceLastSeen")
}

func (pgDbp *PGHandler) GetDeviceProfile(dvId int64) (dv Device, err error) {

	err = pgDbp.DB.QueryRow(
		`SELECT
			*
		FROM
			device 
		WHERE
			id = $1 
		;`, dvId).Scan(
		&dv.id,
		&dv.DevEui,
		&dv.FkWallet,
		&dv.Mode,
		&dv.CreatedAt,
		&dv.LastSeenAt,
		&dv.ApplicationId,
		&dv.Name)
	return dv, errors.Wrap(err, "db/pg_device/GetDeviceProfile")
}

func (pgDbp *PGHandler) GetDeviceListOfWallet(walletId int64, offset int64, limit int64) (dvList []Device, err error) {
	rows, err := pgDbp.DB.Query(
		`SELECT
			*
		FROM
			device 
		WHERE
			fk_wallet = $1 
		ORDER BY id DESC
		LIMIT $2 
		OFFSET $3
	;`, walletId, limit, offset)

	defer rows.Close()

	// res := make([]WithdrawHistRet, 0)
	var dv Device

	for rows.Next() {
		rows.Scan(
			&dv.id,
			&dv.DevEui,
			&dv.FkWallet,
			&dv.Mode,
			&dv.CreatedAt,
			&dv.LastSeenAt,
			&dv.ApplicationId,
			&dv.Name,
		)

		dvList = append(dvList, dv)
	}
	return dvList, errors.Wrap(err, "db/pg_device/GetDeviceListOfWallet")
}

func (pgDbp *PGHandler) GetDeviceRecCnt(walletId int64) (recCnt int64, err error) {

	err = pgDbp.DB.QueryRow(`
		SELECT
			COUNT(*)
		FROM
			device 
		WHERE
			fk_wallet = $1 
	`, walletId).Scan(&recCnt)

	return recCnt, errors.Wrap(err, "db/pg_device/GetDeviceRecCnt")
}
