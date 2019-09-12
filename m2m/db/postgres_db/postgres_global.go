package postgres_db

import "database/sql"

var timeLayout = "2006-01-02T15:04:05.000000Z"

var PgDB *sql.DB
