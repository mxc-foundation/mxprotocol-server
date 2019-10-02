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
			dl_cnt_dv INT    DEFAULT 0 ,
			dl_cnt_dv_free INT    DEFAULT 0 ,
			ul_cnt_dv     INT DEFAULT 0,
			ul_cnt_dv_free INT DEFAULT 0,
			dl_cnt_gw    INT DEFAULT 0,
			dl_cnt_gw_free INT  DEFAULT 0,
			ul_cnt_gw INT  DEFAULT 0,
			ul_cnt_gw_free INT  DEFAULT 0,
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
			fk_wallet,
			dl_cnt_dv,
			dl_cnt_dv_free,
			ul_cnt_dv,
			ul_cnt_dv_free,
			dl_cnt_gw,
			dl_cnt_gw_free,
			ul_cnt_gw,
			ul_cnt_gw_free,
			start_at,
			duration_minutes,
			spend,
			income,
			balance_increase,
			updated_balance
			) 
		VALUES 
			($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)
		RETURNING id ;
	`,
		awu.FkWallet,
		awu.DlCntDv,
		awu.DlCntDvFree,
		awu.UlCntDv,
		awu.UlCntDvFree,
		awu.DlCntGw,
		awu.DlCntGwFree,
		awu.UlCntGw,
		awu.UlCntGwFree,
		awu.StartAt,
		awu.DurationMinutes,
		awu.Spend,
		awu.Income,
		awu.BalanceIncrease,
		awu.UpdatedBalance,
	).Scan(&insertIndex)
	return insertIndex, errors.Wrap(err, "db/pg_agg_wallet_usage/InsertAggWltUsg")
}

func (*aggWalletUsageInterface) GetWalletUsageHist(rogId int64) ([]types.AggWltUsg, error) {
	return []types.AggWltUsg{{}}, nil
}
