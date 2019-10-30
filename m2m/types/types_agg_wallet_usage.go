package types

import "time"

type AggWltUsg struct {
	Id              int64     `db:"id"`
	FkWallet        int64     `db:"fk_wallet"`
	FkAggPeriod     int64     `db:"fk_agg_period"`
	DlCntDv         int64     `db:"dl_cnt_dv"`
	DlCntDvFree     int64     `db:"dl_cnt_dv_free"`
	UlCntDv         int64     `db:"ul_cnt_dv"`
	UlCntDvFree     int64     `db:"ul_cnt_dv_free"`
	DlCntGw         int64     `db:"dl_cnt_gw"`
	DlCntGwFree     int64     `db:"dl_cnt_gw_free"`
	UlCntGw         int64     `db:"ul_cnt_gw"`
	UlCntGwFree     int64     `db:"ul_cnt_gw_free"`
	Spend           float64   `db:"spend"`
	Income          float64   `db:"income"`
	BalanceDelta    float64   `db:"balance_delta"`
	UpdatedBalance  float64   `db:"updated_balance"`
	StartAt         time.Time `db:"start_at"`
	DurationMinutes int64     `db:"duration_minutes"`
}
