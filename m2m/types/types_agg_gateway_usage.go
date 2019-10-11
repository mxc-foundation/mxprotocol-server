package types

import "time"

type AggGwUsg struct {
	Id              int64     `db:"id"`
	FkGateway       int64     `db:"fk_gateway"` // fk in App server
	FkAggPeriod     int64     `db:"fk_agg_period"`
	FkAggWalletUsg  int64     `db:"fk_agg_wallet_usg"`
	DlCnt           int64     `db:"dl_cnt"`
	UlCnt           int64     `db:"ul_cnt"`
	DlCntFree       int64     `db:"dl_cnt_free"`
	UlCntFree       int64     `db:"ul_cnt_free"`
	DlSizeSum       float64   `db:"dl_size_sum"`
	UlSizeSum       float64   `db:"ul_size_sum"`
	Income          float64   `db:"income"`
	StartAt         time.Time `db:"start_at"`
	DurationMinutes float64   `db:"duration_minutes"`
}
