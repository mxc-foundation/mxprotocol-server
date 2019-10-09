package types

import "time"

type DlCategory string

const (
	JOIN_ANS    DlCategory = "JOIN_ANS"
	MAC_COMMAND DlCategory = "MAC_COMMAND"
	PAYLOAD     DlCategory = "PAYLOAD"
	UNKNOWN_PKT DlCategory = "UNKNOWN"
)

type DlPkt struct {
	Id          int64      `db:"id"`
	DlIdNs      string     `db:"dl_id_ns"`
	FkDevice    int64      `db:"dev_eui"` // fk in App server
	FkGateway   int64      `db:"fk_gateway"`
	Nonce       int64      `db:"nonce"`
	TokenDlFrm1 int64      `db:"token_dl_frm1"`
	TokenDlFrm2 int64      `db:"token_dl_frm2"`
	CreatedAt   time.Time  `db:"created_at"`
	Size        float64    `db:"size"`
	Category    DlCategory `db:"category"`
}
