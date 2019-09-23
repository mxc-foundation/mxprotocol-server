package postgres_db

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type deviceInterface struct{}

var PgDevice deviceInterface

func (*deviceInterface) CreateDeviceTable() error {
	_, err := PgDB.Exec(`
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

func (*deviceInterface) InsertDevice(dv types.Device) (insertIndex int64, err error) {
	err = PgDB.QueryRow(`
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

func (*deviceInterface) GetDeviceMode(dvId int64) (dvMode types.DeviceMode, err error) {
	err = PgDB.QueryRow(
		`SELECT
			mode
		FROM
			device 
		WHERE
			id = $1 
		;`, dvId).Scan(&dvMode)
	return dvMode, errors.Wrap(err, "db/pg_device/GetDeviceMode")
}

func (*deviceInterface) SetDeviceMode(dvId int64, dvMode types.DeviceMode) (err error) {
	_, err = PgDB.Exec(`
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

func (*deviceInterface) GetDeviceIdByDevEui(devEui string) (devId int64, err error) {
	err = PgDB.QueryRow(
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

func (*deviceInterface) UpdateDeviceLastSeen(dvId int64, newTime time.Time) (err error) {
	_, err = PgDB.Exec(`
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

func (*deviceInterface) GetDeviceProfile(dvId int64) (dv types.Device, err error) {

	err = PgDB.QueryRow(
		`SELECT
			*
		FROM
			device 
		WHERE
			id = $1 
		;`, dvId).Scan(
		&dv.Id,
		&dv.DevEui,
		&dv.FkWallet,
		&dv.Mode,
		&dv.CreatedAt,
		&dv.LastSeenAt,
		&dv.ApplicationId,
		&dv.Name)
	return dv, errors.Wrap(err, "db/pg_device/GetDeviceProfile")
}

func (*deviceInterface) GetDeviceListOfWallet(orgId int64, offset int64, limit int64) (dvList []types.Device, err error) {
	walletId, err := PgWallet.GetWalletIdFromOrgId(orgId)
	if err != nil {
		return dvList, errors.Wrap(err, "db/pg_device/GetDeviceListOfWallet")
	}

	rows, err := PgDB.Query(
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
	var dv types.Device

	for rows.Next() {
		rows.Scan(
			&dv.Id,
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

func (*deviceInterface) GetDeviceRecCnt(walletId int64) (recCnt int64, err error) {

	err = PgDB.QueryRow(`
		SELECT
			COUNT(*)
		FROM
			device 
		WHERE
			fk_wallet = $1 
	`, walletId).Scan(&recCnt)

	return recCnt, errors.Wrap(err, "db/pg_device/GetDeviceRecCnt")
}
