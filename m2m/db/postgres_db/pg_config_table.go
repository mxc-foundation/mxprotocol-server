package postgres_db

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type configTableInterface struct{}

type ConfigKey string

const (
	LowBalanceWarningKey          = "low_balance_warning"
	DownlinkFeeKey                = "downlink_fee"
	TransactionPercentageShareKey = "transaction_percentage_share"
)

type Config struct {
	LowBalanceWarning          *int
	DownlinkFee                *int
	TransactionPercentageShare *int
}

type row struct {
	Key   ConfigKey
	Value string
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

	return errors.Wrap(err, "db/pg_congif_table/CreateDlPktTable")
}

func (*configTableInterface) Insert(config *Config, ignoreDuplicateKey bool) (err error) {
	query := "INSERT INTO config_table (key, value, updated_time) VALUES "
	values := getValues(config)

	for i := 0; i < len(values); i += 3 {
		query += fmt.Sprintf("($%d, $%d, $%d),", i+1, i+2, i+3)
	}

	//trim the last ,
	query = query[0 : len(query)-1]

	if ignoreDuplicateKey {
		query += " ON CONFLICT (key) DO NOTHING"
	}

	//prepare the statement
	stmt, err := PgDB.Prepare(query)

	if err != nil {
		return errors.Wrap(err, "db/pg_congif_table/Insert")
	}

	//format all vals at once
	_, err = stmt.Exec(values...)

	return errors.Wrap(err, "db/pg_congif_table/Insert")
}

func (*configTableInterface) Get() (config *Config, err error) {
	result, err := PgDB.Query(`
		SELECT key, value
		FROM config_table
	`)

	if err != nil {
		return nil, errors.Wrap(err, "db/pg_congif_table/Get")
	}

	defer result.Close()

	rows := make([]*row, 0)

	for result.Next() {
		r := &row{}

		result.Scan(
			&r.Key,
			&r.Value,
		)

		rows = append(rows, r)
	}

	config = &Config{}

	err = setValues(rows, config)

	if err != nil {
		return nil, errors.Wrap(err, "db/pg_congif_table/Get")
	}

	return config, nil
}

func (*configTableInterface) GetOne(key ConfigKey) (value string, err error) {
	err = PgDB.QueryRow(`
		SELECT
			value
		FROM
				config_table
		WHERE
			key = $1
		`, key).Scan(&value)

	return value, errors.Wrap(err, "db/pg_congif_table/GetOne")
}

func (*configTableInterface) Update(config *Config) (err error) {
	query := `UPDATE config_table AS t
		SET
			key = c.key,
			value = c.value,
			updated_time = c.updated_time
		FROM (values `

	values := getValues(config)

	for i := 0; i < len(values); i += 3 {
		query += fmt.Sprintf("($%d, $%d, $%d::timestamptz),", i+1, i+2, i+3)
	}

	//trim the last ,
	query = query[0 : len(query)-1]

	query += `)
		AS c(key, value, updated_time)
		WHERE c.key = t.key`

	//prepare the statement
	stmt, err := PgDB.Prepare(query)

	if err != nil {
		return errors.Wrap(err, "db/pg_congif_table/Update")
	}

	//format all vals at once
	_, err = stmt.Exec(values...)

	return errors.Wrap(err, "db/pg_congif_table/Update")
}

func (*configTableInterface) UpdateOne(key ConfigKey, value string) (err error) {
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
	return errors.Wrap(err, "db/pg_congif_table/UpdateOne")
}

func getValues(config *Config) []interface{} {
	updatedTime := time.Now().UTC()
	values := []interface{}{}

	if config.LowBalanceWarning != nil {
		values = append(values, LowBalanceWarningKey, config.LowBalanceWarning, updatedTime)
	}

	if config.DownlinkFee != nil {
		values = append(values, DownlinkFeeKey, config.DownlinkFee, updatedTime)
	}

	if config.TransactionPercentageShare != nil {
		values = append(values, TransactionPercentageShareKey, config.TransactionPercentageShare, updatedTime)
	}

	return values
}

func setValues(rows []*row, config *Config) error {
	for _, r := range rows {
		switch true {
		case r.Key == LowBalanceWarningKey:
			value, err := strconv.Atoi(r.Value)
			if err != nil {
				return err
			}
			config.LowBalanceWarning = &value
		case r.Key == DownlinkFeeKey:
			value, err := strconv.Atoi(r.Value)
			if err != nil {
				return err
			}
			config.DownlinkFee = &value
		case r.Key == TransactionPercentageShareKey:
			value, err := strconv.Atoi(r.Value)
			if err != nil {
				return err
			}
			config.TransactionPercentageShare = &value
		}
	}

	return nil
}
