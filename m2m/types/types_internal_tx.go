package types

import "time"

type InternalTx struct {
	Id             int64     `db:"id"`
	FkWalletSender int64     `db:"fk_wallet_sender"`
	FkWalletRcvr   int64     `db:"fk_wallet_receiver"`
	PaymentCat     string    `db:"payment_cat"`
	TxInternalRef  int64     `db:"tx_internal_ref"` // reference to the id of corresponding table to PaymentCat
	Value          float64   `db:"value"`
	TimeTx         time.Time `db:"timestamp"`
}

type PaymentCategory string // db:payment_category
const (
	WITHDRAW_FEE_SN_INCOME PaymentCategory = "WITHDRAW_FEE_SN_INCOME"
	DOWNLINK_AGG_SN_INCOME PaymentCategory = "DOWNLINK_AGG_SN_INCOME"
	DOWNLINK_AGGREGATION   PaymentCategory = "DOWNLINK_AGGREGATION"
	PURCHASE_SUBSCRIPTION  PaymentCategory = "PURCHASE_SUBSCRIPTION"
	BUY_SUBSCRIPTION       PaymentCategory = "BUY_SUBSCRIPTION"
	TOP_UP                 PaymentCategory = "TOP_UP"
	WITHDRAW              PaymentCategory = "WITHDRAW"
	STAKE_REVENUE         PaymentCategory = "STAKE_REVENUE"
	INSERT_STAKE          PaymentCategory = "INSERT_STAKE"
	UNSTAKE               PaymentCategory = "UNSTAKE"
)
