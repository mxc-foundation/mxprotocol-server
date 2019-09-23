package types

import "time"

type WithdrawHistRet struct {
	AcntSender  string
	AcntRcvr    string
	ExtCurrency string
	Value       float64
	WithdrawFee float64
	TxSentTime  time.Time `db:"tx_sent_time"`
	TxStatus    string    `db:"tx_status"`
	TxAprvdTime time.Time
	TxHash      string
}
