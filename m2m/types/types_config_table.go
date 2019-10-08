package types

import "time"

type ConfigTable struct {
	Key        string    `db:"key"`
	Value      string    `db:"value"`
	UpdateTime time.Time `db:"update_time"`
}
