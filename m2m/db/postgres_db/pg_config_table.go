package postgres_db

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type configTableInterface struct{}

type Config struct {
	Key   string
	Value interface{}
}

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

func (*configTableInterface) InsertConfigs(data map[string]interface{}, ignoreDuplicateKey bool) (err error) {
	updatedTime := time.Now().UTC()

	query := "INSERT INTO config_table (key, value, updated_time) VALUES "
	values := []interface{}{}

	i := 0
	for key, value := range data {
		query += fmt.Sprintf("($%d, $%d, $%d),", i+1, i+2, i+3)
		values = append(values, key, value, updatedTime)
		i += 3
	}

	//trim the last ,
	query = query[0 : len(query)-1]

	if ignoreDuplicateKey {
		query += " ON CONFLICT (key) DO NOTHING"
	}

	//prepare the statement
	stmt, err := PgDB.Prepare(query)

	if err != nil {
		return errors.Wrap(err, "db/pg_congif_table/InsertConfigs")
	}

	//format all vals at once
	_, err = stmt.Exec(values...)

	return errors.Wrap(err, "db/pg_congif_table/InsertConfigs")
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

func (*configTableInterface) UpdateConfigs(data map[string]interface{}) (err error) {
	query := `UPDATE config_table AS t
		SET
			key = c.key,
			value = c.value,
			updated_time = c.updated_time
		FROM (values `

	updatedTime := time.Now().UTC()
	values := []interface{}{}

	i := 0
	for key, value := range data {
		query += fmt.Sprintf("($%d, $%d, $%d::timestamptz),", i+1, i+2, i+3)
		values = append(values, key, value, updatedTime)
		i += 3
	}

	//trim the last ,
	query = query[0 : len(query)-1]

	query += `)
		AS c(key, value, updated_time)
		WHERE c.key = t.key`

	//prepare the statement
	stmt, err := PgDB.Prepare(query)

	if err != nil {
		return errors.Wrap(err, "db/pg_congif_table/UpdateConfigs")
	}

	//format all vals at once
	_, err = stmt.Exec(values...)

	return errors.Wrap(err, "db/pg_congif_table/UpdateConfigs")
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

func (*configTableInterface) GetConfigs(keys []string) (configs []Config, err error) {
	query := `
	SELECT 
		key, value
	FROM
			config_table 
	WHERE
		key IN (`

	args := make([]interface{}, len(keys))

	for i, key := range keys {
		query += fmt.Sprintf("$%d,", i+1)
		args[i] = key
	}

	query = query[0:len(query)-1] + ")"

	rows, err := PgDB.Query(query, args...)

	if err != nil {
		return nil, errors.Wrap(err, "db/pg_congif_table/GetConfigs")
	}

	defer rows.Close()

	config := Config{}

	for rows.Next() {
		rows.Scan(
			&config.Key,
			&config.Value,
		)

		configs = append(configs, config)
	}

	return configs, errors.Wrap(err, "db/pg_congif_table/GetConfigs")
}
