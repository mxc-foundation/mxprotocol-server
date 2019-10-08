package types

import "time"

type ConfigTable struct {
	Id         int64     `db:"id"`
	Key        string    `db:"key"`
	Value      string    `db:"value"`
	UpdateTime time.Time `db:"update_time"`
}
