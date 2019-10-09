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
