package types

import "time"

type Device struct {
	Id            int64     `db:"id"`
	DevEui        string    `db:"dev_eui"` // fk in App server
	FkWallet      int64     `db:"fk_wallet"`
	Mode          string    `db:"mode"`
	CreatedAt     time.Time `db:"created_at"`
	LastSeenAt    time.Time `db:"last_seen_at"`
	ApplicationId int64     `db:"application_id"`
	Name          string    `db:"name"`
}
