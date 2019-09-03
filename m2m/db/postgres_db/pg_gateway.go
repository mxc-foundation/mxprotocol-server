package postgres_db

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Gateway struct {
	id          int64     `db:"id"`
	Mac         string    `db:"mac"` // fk in AS (App Server)
	FkGatewayNs int64     `db:"fk_gateway_ns"`
	FkWallet    int64     `db:"fk_wallet"`
	Mode        string    `db:"mode"`
	CreatedAt   time.Time `db:"created_at"`
	LastSeenAt  time.Time `db:"last_seen_at"`
	OrgId       int64     `db:"org_id"`
	Description string    `db:"description"`
	Name        string    `db:"name"`
}

func (pgDbp *PGHandler) CreateGatewayTable() error {
	_, err := pgDbp.DB.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS 
				(SELECT 1 FROM pg_type WHERE typname = 'gateway_mode') 
			THEN
				CREATE TYPE gateway_mode AS ENUM (
					'INACTIVE',
					 'FREE_GATEWAYS_LIMITED',
					 'WHOLE_NETWORK',
					 'DELETED'
		);
		END IF;
			CREATE TABLE IF NOT EXISTS gateway (
			id SERIAL PRIMARY KEY,
			mac VARCHAR(128) NOT NULL,
			fk_gateway_ns INT NOT NULL,
			fk_wallet INT REFERENCES wallet (id) NOT NULL,
			mode GATEWAY_MODE    NOT NULL,
			created_at     TIMESTAMP,
			last_seen_at    TIMESTAMP,
			org_id INT,
			description  varchar(128),    
			name  varchar(128)    
		);

		END$$;
		
	`)
	return errors.Wrap(err, "db/pg_gateway/CreateGateway")
}

func (pgDbp *PGHandler) InsertGateway(gw Gateway) (insertIndex int64, err error) {
	err = pgDbp.DB.QueryRow(`
		INSERT INTO gateway (
			mac ,
			fk_gateway_ns,
			fk_wallet,
			mode,
			created_at,
			last_seen_at,
			org_id,
			description,
			name
			) 
		VALUES 
			($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING id ;
	`,
		gw.Mac,
		gw.FkGatewayNs,
		gw.FkWallet,
		gw.Mode,
		gw.CreatedAt,
		gw.LastSeenAt,
		gw.OrgId,
		gw.Description,
		gw.Name,
	).Scan(&insertIndex)
	return insertIndex, errors.Wrap(err, "db/pg_gateway/InsertGateway")
}

func (pgDbp *PGHandler) GetGatewayMode(gwId int64) (gwMode string, err error) {
	err = pgDbp.DB.QueryRow(
		`SELECT
			mode
		FROM
			gateway 
		WHERE
			id = $1 
		;`, gwId).Scan(&gwMode)
	return gwMode, errors.Wrap(err, "db/pg_gateway/GetGatewayMode")
}

func (pgDbp *PGHandler) SetGatewayMode(gwId int64, gwMode string) (err error) {
	_, err = pgDbp.DB.Exec(`
		UPDATE
			gateway 
		SET
			mode = $1
		WHERE
			id = $2
		;
		`, gwMode, gwId)

	return errors.Wrap(err, "db/pg_gateway/SetGatewayMode")

}

func (pgDbp *PGHandler) GetGatewayIdByMac(mac string) (gwId int64, err error) {
	err = pgDbp.DB.QueryRow(
		`SELECT
			id
		FROM
			gateway 
		WHERE
			mac = $1 
		ORDER BY id DESC 
		LIMIT 1  
		;`, mac).Scan(&gwId)
	return gwId, errors.Wrap(err, "db/pg_gateway/GetGatewayIdByMac")
}

func (pgDbp *PGHandler) UpdateGatewayLastSeen(gwId int64, newTime time.Time) (err error) {
	_, err = pgDbp.DB.Exec(`
		UPDATE
			gateway
		SET
			last_seen_at = $1
		WHERE
			id = $2
		;
		`, newTime, gwId)

	return errors.Wrap(err, "db/pg_gateway/UpdateGatewayLastSeen")
}

func (pgDbp *PGHandler) GetGatewayProfile(gwId int64) (gw Gateway, err error) {

	err = pgDbp.DB.QueryRow(
		`SELECT
			*
		FROM
			gateway 
		WHERE
			id = $1 
		;`, gwId).Scan(
		&gw.id,
		&gw.Mac,
		&gw.FkGatewayNs,
		&gw.FkWallet,
		&gw.Mode,
		&gw.CreatedAt,
		&gw.LastSeenAt,
		&gw.OrgId,
		&gw.Description,
		&gw.Name)
	return gw, errors.Wrap(err, "db/pg_gateway/GetGatewayProfile")
}

func (pgDbp *PGHandler) GetGatewayListOfWallet(walletId int64, offset int64, limit int64) (gwList []Gateway, err error) {
	rows, err := pgDbp.DB.Query(
		`SELECT
			*
		FROM
			gateway 
		WHERE
			fk_wallet = $1 
		ORDER BY id DESC
		LIMIT $2 
		OFFSET $3
	;`, walletId, limit, offset)

	defer rows.Close()

	// res := make([]WithdrawHistRet, 0)
	var gw Gateway

	for rows.Next() {
		rows.Scan(
			&gw.id,
			&gw.Mac,
			&gw.FkGatewayNs,
			&gw.FkWallet,
			&gw.Mode,
			&gw.CreatedAt,
			&gw.LastSeenAt,
			&gw.OrgId,
			&gw.Description,
			&gw.Name)

		gwList = append(gwList, gw)
	}
	return gwList, errors.Wrap(err, "db/pg_gateway/GetGatewayListOfWallet")
}

func (pgDbp *PGHandler) GetGatewayRecCnt(walletId int64) (recCnt int64, err error) {

	err = pgDbp.DB.QueryRow(`
		SELECT
			COUNT(*)
		FROM
			gateway 
		WHERE
			fk_wallet = $1 
	`, walletId).Scan(&recCnt)

	return recCnt, errors.Wrap(err, "db/pg_gateway/GetGatewayRecCnt")
}
