package postgres_db

import (
	"time"

	"github.com/pkg/errors"
)

type configTableInterface struct{}

var PgConfigTable configTableInterface

func (*configTableInterface) CreateConfigTable() error {
	_, err := PgDB.Exec(`
	
	CREATE TABLE IF NOT EXISTS config_table (
		key VARCHAR(128) PRIMARY KEY,
		value VARCHAR(128) NOT NULL,
		updated_time TIMESTAMP
	);
	
`)

	return errors.Wrap(err, "db/pg_congif_table/CreateConfigTable")
}

func (*configTableInterface) InsertConfig(key string, value string) (err error) {
	_, err = PgDB.Exec(`
	INSERT INTO config_table 
		(key,
		value,
		updated_time)
	VALUES ($1, $2, $3)
`,
		key,
		value,
		time.Now().UTC(),
	)
	return errors.Wrap(err, "db/pg_congif_table/InsertConfig")
}

func (*configTableInterface) UpdateConfig(key string, value string) (err error) {
	_, err = PgDB.Exec(`
	UPDATE config_table 
	SET
		value = $1 , 
		updated_time = $2
	WHERE
		key = $3;
`,
		value,
		time.Now().UTC(),
		key,
	)
	return errors.Wrap(err, "db/pg_congif_table/UpdateConfig")
}

func (*configTableInterface) GetConfig(key string) (val string, err error) {
	err = PgDB.QueryRow(`
	SELECT 
		value
	FROM
	    config_table 
	WHERE
		key = $1
	;`, key).Scan(&val)
	return val, errors.Wrap(err, "db/pg_congif_table/GetConfig")
}
