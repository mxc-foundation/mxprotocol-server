package db

import "database/sql"

var DbError struct {
	NoRowQueryRes error
}

func dbErrorInit() {
	DbError.NoRowQueryRes = sql.ErrNoRows
}
