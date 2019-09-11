package types

import "time"

type AggDvUsg struct {
	Id              int64     `db:"id"`
	FkDevice        int64     `db:"fk_device"` // fk in App server
	FkAggWalletUsg  int64     `db:"fk_agg_wallet_usg"`
	DlCnt           int64     `db:"dl_cnt"`
	UlCnt           int64     `db:"ul_cnt"`
	DlCntFree       int64     `db:"dl_cnt_free"`
	UlCntFree       int64     `db:"ul_cnt_free"`
	DlSizeSum       float64   `db:"dl_size_sum"`
	UlSizeSum       float64   `db:"ul_size_sum"`
	StartAt         time.Time `db:"start_at"`
	DurationMinutes float64   `db:"duration_minutes"`
	Spend           float64   `db:"spend"`
}
