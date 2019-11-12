package postgres_db

import (
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
					 'UNKNOWN_PKT'
		);
		END IF;
			CREATE TABLE IF NOT EXISTS dl_pkt (
			id SERIAL PRIMARY KEY,
			dl_id_ns VARCHAR (128),
			fk_device INT REFERENCES device (id) NOT NULL,
			fk_gateway INT REFERENCES gateway (id) NOT NULL,
			nonce INT,
			token_dl_frm1 INT,
			token_dl_frm2 INT,
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
			dl_id_ns,
			fk_device,
			fk_gateway,
			nonce ,
			token_dl_frm1,
			token_dl_frm2,
			created_at,
			size ,
			category
			)
		VALUES
			($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING id ;
	`,
		dlPkt.DlIdNs,
		dlPkt.FkDevice,
		dlPkt.FkGateway,
		dlPkt.Nonce,
		dlPkt.TokenDlFrm1,
		dlPkt.TokenDlFrm2,
		dlPkt.CreatedAt,
		dlPkt.Size,
		dlPkt.Category,
	).Scan(&insertIndex)
	return insertIndex, errors.Wrap(err, "db/pg_dl_pkt/InsertDlPkt")
}

func (*dlPacketInterface) GetAggDlPktDeviceWallet(startIndDlPkt, endIndDlPkt int64) (walletId []int64, count []int64, err error) {
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
		dlp.id BETWEEN
			$1
		AND
			$2
	GROUP BY
		dv.fk_wallet;
	`, startIndDlPkt, endIndDlPkt)

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

func (*dlPacketInterface) GetAggDlPktGatewayWallet(startIndDlPkt, endIndDlPkt int64) (walletId []int64, count []int64, err error) {
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
		dlp.id BETWEEN
			$1 
		AND
			$2
	GROUP BY
		gw.fk_wallet;
	`, startIndDlPkt, endIndDlPkt)

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

func (*dlPacketInterface) GetAggDlPktFreeWallet(startIndDlPkt, endIndDlPkt int64) (walletId []int64, count []int64, err error) {
	rows, err := PgDB.Query(`
	SELECT
		dv.fk_wallet as wallet_id,
		count(*)
	FROM
		dl_pkt dlp,
		device dv,
		gateway gw
	WHERE
		dlp.fk_device = dv.id
		AND
		dlp.fk_gateway = gw.id
		AND
		dv.fk_wallet = gw.fk_wallet
		AND
		dlp.id BETWEEN
			$1
		AND
			$2
	GROUP BY
		dv.fk_wallet;
	`, startIndDlPkt, endIndDlPkt)

	if err != nil {
		return nil, nil, errors.Wrap(err, "db/pg_dl_packet/GetAggDlPktFreeWallet")
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

func (*dlPacketInterface) GetLastReceivedDlPktId() (latestId int64, err error) {
	err = PgDB.QueryRow(`
		SELECT
			MAX (id)
		FROM
			dl_pkt
		;
	`).Scan(&latestId)
	return latestId, errors.Wrap(err, "db/pg_dl_pkt/GetLastReceivedDlPktId")
}
