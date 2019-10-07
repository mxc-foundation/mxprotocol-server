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
