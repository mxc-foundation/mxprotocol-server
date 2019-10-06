package postgres_db

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type dlPacketInterface struct{}

var PgDlPacket dlPacketInterface

func (*dlPacketInterface) CreateDlPktTable() error {
	_, err := PgDB.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS 
				(SELECT 1 FROM pg_type WHERE typname = 'dl_category') 
			THEN
				CREATE TYPE dl_category AS ENUM (
					'JOIN_ANS',
					 'MAC_COMMAND',
					 'PAYLOAD',
					 'UNKNOWN'
		);
		END IF;
			CREATE TABLE IF NOT EXISTS dl_pkt (
			id SERIAL PRIMARY KEY,
			fk_device INT REFERENCES device (id) NOT NULL,
			fk_gateway INT REFERENCES gateway (id) NOT NULL,
			nonce INT,
			created_at     TIMESTAMP,
			size FLOAT ,
			category  dl_category
		);

		END$$;
		
	`)
	return errors.Wrap(err, "db/pg_dl_pkt/CreateDlPktTable")
}

func (*dlPacketInterface) InsertDlPkt(dlPkt types.DlPkt) (insertIndex int64, err error) {
	err = PgDB.QueryRow(`
		INSERT INTO dl_pkt (
			fk_device,
			fk_gateway,
			nonce ,
			created_at,
			size ,
			category
			)
		VALUES
			($1,$2,$3,$4,$5,$6)
		RETURNING id ;
	`,
		dlPkt.FkDevice,
		dlPkt.FkGateway,
		dlPkt.Nonce,
		dlPkt.CreatedAt,
		dlPkt.Size,
		dlPkt.Category,
	).Scan(&insertIndex)
	return insertIndex, errors.Wrap(err, "db/pg_dl_pkt/InsertDlPkt")
}

func (*dlPacketInterface) GetAggDlPktDeviceWallet(begin time.Time, durationMin int64) (walletId []int64, count []int64, err error) {
	rows, err := PgDB.Query(`
	SELECT 
		dv.fk_wallet as wallet_id,
		count(*)
	FROM
		dl_pkt dlp,
		device dv
	WHERE 
		dlp.fk_device = dv.id
	AND
		dlp.created_at BETWEEN
			$1 
		AND
			current_timestamp + ($2 * interval '1 minute')
	GROUP BY
		dv.fk_wallet;
	`, begin)

	if err != nil {
		return nil, nil, errors.Wrap(err, "db/pg_dl_packet/getAggDlPktWallet")
	}

	defer rows.Close()

	var wltIdVal, cntVal int64

	for rows.Next() {
		rows.Scan(
			&wltIdVal,
			&cntVal,
		)

		walletId = append(walletId, wltIdVal)
		count = append(count, cntVal)
	}

	return walletId, count, nil
}

func (*dlPacketInterface) GetAggDlPktGatewayWallet(begin time.Time, durationMin int64) (walletId []int64, count []int64, err error) {
	rows, err := PgDB.Query(`
	SELECT 
		gw.fk_wallet as wallet_id,
		count(*)
	FROM
		dl_pkt dlp,
		gateway gw
	WHERE 
		dlp.fk_gateway = gw.id
	AND
		dlp.created_at BETWEEN
			$1 
		AND
			current_timestamp + ($2 * interval '1 minute')
	GROUP BY
		gw.fk_wallet;
	`, begin)

	if err != nil {
		return nil, nil, errors.Wrap(err, "db/pg_dl_packet/GetAggDlPktGatewayWallet")
	}

	defer rows.Close()

	var wltIdVal, cntVal int64

	for rows.Next() {
		rows.Scan(
			&wltIdVal,
			&cntVal,
		)

		walletId = append(walletId, wltIdVal)
		count = append(count, cntVal)
	}

	return walletId, count, nil
}
