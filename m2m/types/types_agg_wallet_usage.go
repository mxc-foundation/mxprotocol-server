package types

import "time"

type AggWltUsg struct {
	Id              int64     `db:"id"`
	FkWallet        int64     `db:"fk_wallet"`
	DlCnt           int64     `db:"dl_cnt"`
	UlCnt           int64     `db:"ul_cnt"`
	DlCntFree       int64     `db:"dl_cnt_free"`
	UlCntFree       int64     `db:"ul_cnt_free"`
	DlSizeSum       float64   `db:"dl_size_sum"`
	UlSizeSum       float64   `db:"ul_size_sum"`
	StartAt         time.Time `db:"start_at"`
	DurationMinutes float64   `db:"duration_minutes"`
	Spend           float64   `db:"spend"`
	Income          float64   `db:"income"`
	BalanceIncrease float64   `db:"balance_increase"`
	UpdatedBalance  float64   `db:"updated_balance"`
}
