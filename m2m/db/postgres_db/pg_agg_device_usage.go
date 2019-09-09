package postgres_db

import (
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	types "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

func (pgDbp *PGHandler) CreateAggDvUsgTable() error {
	_, err := pgDbp.DB.Exec(`
	
		CREATE TABLE IF NOT EXISTS agg_device_usage (
			id SERIAL PRIMARY KEY,
			fk_device INT REFERENCES device (id) NOT NULL,
			fk_agg_wallet_usage INT , -- REFERENCES agg_wallet_usage (id)   @@ to be added
			dl_cnt INT    DEFAULT 0 ,
			ul_cnt     INT DEFAULT 0,
			dl_cnt_free    INT DEFAULT 0,
			ul_cnt_free INT  DEFAULT 0,
			dl_size_sum  FLOAT    DEFAULT 0,
			ul_size_sum  FLOAT DEFAULT 0,
			start_at TIMESTAMP NOT NULL,
			duration_minutes   INT ,
			spend  NUMERIC(28,18) DEFAULT 0
		);		
	`)
	return errors.Wrap(err, "db/pg_agg_device_usage/CreateAggDvUsgTable")
}

func (pgDbp *PGHandler) InsertAggDvUsg(adu types.AggDvUsg) (insertIndex int64, err error) {
	err = pgDbp.DB.QueryRow(`
		INSERT INTO agg_device_usage (
			fk_device ,
			fk_agg_wallet_usage ,
			dl_cnt ,
			ul_cnt ,
			dl_cnt_free ,
			ul_cnt_free ,
			dl_size_sum ,
			ul_size_sum ,
			start_at ,
			duration_minutes ,
			spend  
			) 
		VALUES 
			($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		RETURNING id ;
	`,
		adu.FkDevice,
		adu.FkAggWalletUsg,
		adu.DlCnt,
		adu.UlCnt,
		adu.DlCntFree,
		adu.UlCntFree,
		adu.DlSizeSum,
		adu.UlSizeSum,
		adu.StartAt,
		adu.DurationMinutes,
		adu.Spend,
	).Scan(&insertIndex)
	return insertIndex, errors.Wrap(err, "db/pg_agg_device_usage/InsertAggDvUsg")
}
