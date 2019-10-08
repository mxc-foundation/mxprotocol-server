package postgres_db

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type gatewayInterface struct{}

var PgGateway gatewayInterface

func (*gatewayInterface) CreateGatewayTable() error {
	_, err := PgDB.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS 
				(SELECT 1 FROM pg_type WHERE typname = 'gateway_mode') 
			THEN
				CREATE TYPE gateway_mode AS ENUM (
					'GW_INACTIVE',
					 'GW_FREE_GATEWAYS_LIMITED',
					 'GW_WHOLE_NETWORK',
					 'GW_DELETED'
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

func (*gatewayInterface) InsertGateway(gw types.Gateway) (insertIndex int64, err error) {
	err = PgDB.QueryRow(`
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

func (*gatewayInterface) GetGatewayMode(gwId int64) (gwMode types.GatewayMode, err error) {
	err = PgDB.QueryRow(
		`SELECT
			mode
		FROM
			gateway 
		WHERE
			id = $1 
		;`, gwId).Scan(&gwMode)
	return gwMode, errors.Wrap(err, "db/pg_gateway/GetGatewayMode")
}

func (*gatewayInterface) SetGatewayMode(gwId int64, gwMode types.GatewayMode) (err error) {
	_, err = PgDB.Exec(`
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

func (*gatewayInterface) GetGatewayIdByMac(mac string) (gwId int64, err error) {
	err = PgDB.QueryRow(
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

func (*gatewayInterface) UpdateGatewayLastSeen(gwId int64, newTime time.Time) (err error) {
	_, err = PgDB.Exec(`
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

func (*gatewayInterface) GetGatewayProfile(gwId int64) (gw types.Gateway, err error) {

	err = PgDB.QueryRow(
		`SELECT
			*
		FROM
			gateway 
		WHERE
			id = $1 
		;`, gwId).Scan(
		&gw.Id,
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

func (*gatewayInterface) GetGatewayListOfWallet(walletId int64, offset int64, limit int64) (gwList []types.Gateway, err error) {

	rows, err := PgDB.Query(
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

	var gw types.Gateway

	for rows.Next() {
		rows.Scan(
			&gw.Id,
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

func (*gatewayInterface) GetGatewayRecCnt(walletId int64) (recCnt int64, err error) {

	err = PgDB.QueryRow(`
		SELECT
			COUNT(*)
		FROM
			gateway 
		WHERE
			fk_wallet = $1 
	`, walletId).Scan(&recCnt)

	return recCnt, errors.Wrap(err, "db/pg_gateway/GetGatewayRecCnt")
}

func (*gatewayInterface) GetFreeGwList(walletId int64) (gwId []int64, gwMac []string, err error) {

	rows, err := PgDB.Query(
		`SELECT
			id as gwId, mac
		FROM
			gateway 
		WHERE
			fk_wallet = $1 
	;`, walletId)

	defer rows.Close()

	var id int64
	var mac string

	for rows.Next() {
		rows.Scan(
			&id,
			&mac,
		)

		gwId = append(gwId, id)
		gwMac = append(gwMac, mac)
	}
	return gwId, gwMac, errors.Wrap(err, "db/pg_gateway/GetFreeGwList")

}

func (*gatewayInterface) GetWalletIdOfGateway(gwId int64) (gwWalletId int64, err error) {
	err = PgDB.QueryRow(
		`SELECT
			fk_wallet
		FROM
			gateway 
		WHERE	
			id = $1 
			ORDER BY id DESC 
			LIMIT 1  
		;`, gwId).Scan(&gwWalletId)
	return gwWalletId, errors.Wrap(err, "db/pg_gateway/GetWalletIdOfGateway")
}
