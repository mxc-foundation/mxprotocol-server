package postgres_db

import (
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type aggWalletUsageInterface struct{}

var PgAggWalletUsage aggWalletUsageInterface

func (*aggWalletUsageInterface) CreateAggWltUsgTable() error {
	_, err := PgDB.Exec(`
	
		CREATE TABLE IF NOT EXISTS agg_wallet_usage (
			id SERIAL PRIMARY KEY,
			fk_wallet INT REFERENCES wallet (id) NOT NULL,
			dl_cnt INT    DEFAULT 0 ,
			ul_cnt     INT DEFAULT 0,
			dl_cnt_free    INT DEFAULT 0,
			ul_cnt_free INT  DEFAULT 0,
			dl_size_sum  FLOAT    DEFAULT 0,
			ul_size_sum  FLOAT DEFAULT 0,
			start_at TIMESTAMP NOT NULL,
			duration_minutes   INT ,
			spend  NUMERIC(28,18) DEFAULT 0,
			income  NUMERIC(28,18) DEFAULT 0,
			balance_increase  NUMERIC(28,18) DEFAULT 0,
			updated_balance  NUMERIC(28,18) DEFAULT 0

		);		
	`)
	return errors.Wrap(err, "db/pg_agg_wallet_usage/CreateAggWltUsgTable")
}

func (*aggWalletUsageInterface) InsertAggWltUsg(awu types.AggWltUsg) (insertIndex int64, err error) {
	err = PgDB.QueryRow(`
		INSERT INTO agg_wallet_usage (
			fk_wallet ,
			dl_cnt ,
			ul_cnt ,
			dl_cnt_free ,
			ul_cnt_free ,
			dl_size_sum ,
			ul_size_sum ,
			start_at ,
			duration_minutes ,
			spend ,
			income,
			balance_increase,
			updated_balance
			) 
		VALUES 
			($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
		RETURNING id ;
	`,
		awu.FkWallet,
		awu.DlCnt,
		awu.UlCnt,
		awu.DlCntFree,
		awu.UlCntFree,
		awu.DlSizeSum,
		awu.UlSizeSum,
		awu.StartAt,
		awu.DurationMinutes,
		awu.Spend,
		awu.Income,
		awu.BalanceIncrease,
		awu.Income-awu.Spend,
	).Scan(&insertIndex)
	return insertIndex, errors.Wrap(err, "db/pg_agg_wallet_usage/InsertAggWltUsg")
}

func (*aggWalletUsageInterface)GetWalletUsageHist(rogId int64) ([]types.AggWltUsg, error){
	return []types.AggWltUsg{{}}, nil
}