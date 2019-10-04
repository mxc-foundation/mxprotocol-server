package types

import "time"

type AggWltUsg struct {
	Id              int64     `db:"id"`
	FkWallet        int64     `db:"fk_wallet"`
	DlCntDv         int64     `db:"dl_cnt_dv"`
	DlCntDvFree     int64     `db:"dl_cnt_dv_free"`
	UlCntDv         int64     `db:"ul_cnt_dv"`
	UlCntDvFree     int64     `db:"ul_cnt_dv_free"`
	DlCntGw         int64     `db:"dl_cnt_gw"`
	DlCntGwFree     int64     `db:"dl_cnt_gw_free"`
	UlCntGw         int64     `db:"ul_cnt_gw"`
	UlCntGwFree     int64     `db:"ul_cnt_gw_free"`
	StartAt         time.Time `db:"start_at"`
	DurationMinutes int64     `db:"duration_minutes"`
	Spend           float64   `db:"spend"`
	Income          float64   `db:"income"`
	BalanceIncrease float64   `db:"balance_increase"`
	UpdatedBalance  float64   `db:"updated_balance"`
}
