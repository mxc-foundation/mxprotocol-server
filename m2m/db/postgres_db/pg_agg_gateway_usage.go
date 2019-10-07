package postgres_db

import (
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type aggGatewayUsageInterface struct{}

var PgAggGatewayUsage aggGatewayUsageInterface

func (*aggGatewayUsageInterface) CreateAggGwUsgTable() error {
	_, err := PgDB.Exec(`
	
		CREATE TABLE IF NOT EXISTS agg_gateway_usage (
			id SERIAL PRIMARY KEY,
			fk_gateway INT REFERENCES gateway (id) NOT NULL,
			fk_agg_wallet_usage INT  REFERENCES agg_wallet_usage (id) ,
			dl_cnt INT    DEFAULT 0 ,
			ul_cnt     INT DEFAULT 0,
			dl_cnt_free    INT DEFAULT 0,
			ul_cnt_free INT  DEFAULT 0,
			dl_size_sum  FLOAT    DEFAULT 0,
			ul_size_sum  FLOAT DEFAULT 0,
			start_at TIMESTAMP NOT NULL,
			duration_minutes   INT ,
			income  NUMERIC(28,18) DEFAULT 0
		);		
	`)
	return errors.Wrap(err, "db/pg_agg_gateway_usage/CreateAggGwUsgTable")
}

func (*aggGatewayUsageInterface) InsertAggGwUsg(agu types.AggGwUsg) (insertIndex int64, err error) {
	err = PgDB.QueryRow(`
		INSERT INTO agg_gateway_usage (
			fk_gateway ,
			fk_agg_wallet_usage ,
			dl_cnt ,
			ul_cnt ,
			dl_cnt_free ,
			ul_cnt_free ,
			dl_size_sum ,
			ul_size_sum ,
			start_at ,
			duration_minutes ,
			income  
			) 
		VALUES 
			($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		RETURNING id ;
	`,
		agu.FkGateway,
		agu.FkAggWalletUsg,
		agu.DlCnt,
		agu.UlCnt,
		agu.DlCntFree,
		agu.UlCntFree,
		agu.DlSizeSum,
		agu.UlSizeSum,
		agu.StartAt,
		agu.DurationMinutes,
		agu.Income,
	).Scan(&insertIndex)
	return insertIndex, errors.Wrap(err, "db/pg_agg_gateway_usage/InsertAggGwUsg")
}
