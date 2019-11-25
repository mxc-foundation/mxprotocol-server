package types

import "time"

type AggPeriod struct {
	Id              int64     `db:"id"`
	StartAt         time.Time `db:"start_at"`
	DurationMinutes int64     `db:"duration_minutes"`
}

type AggregationStatus string

const (
	AGG_IN_PROCESS AggregationStatus = "IN_PROCESS"
	AGG_COMPLETED  AggregationStatus = "COMPLETED"
)
