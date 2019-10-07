package types

import "time"

type TopupHistRet struct {
	AcntSender  string
	AcntRcvr    string
	ExtCurrency string
	Value       float64
	TxAprvdTime time.Time
	TxHash      string
}
