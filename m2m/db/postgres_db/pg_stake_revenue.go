package postgres_db

import (
	"fmt"

	"github.com/pkg/errors"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type stakeRevenueInterface struct{}

var PgStakeRevenue stakeRevenueInterface

func (*stakeRevenueInterface) CreateStakeRevenueTable() error {
	_, err := PgDB.Exec(
		`
		CREATE TABLE IF NOT EXISTS stake_revenue (
			id SERIAL PRIMARY KEY,
			fk_stake_revenue_period INT REFERENCES stake_revenue_period (id) NOT NULL,
			fk_stake INT REFERENCES stake(id) NOT NULL,
			r_fk_wallet INT REFERENCES wallet (id),
			revenue_amount NUMERIC(28,18) DEFAULT 0 NOT NULL  CHECK (revenue_amount >= 0),
			updated_balance NUMERIC(28,18)
		);	
	`)
	return errors.Wrap(err, "db/pg_stake_revenue/CreateStakeRevenueTable")
}

func (*stakeRevenueInterface) InsertStakeRevenue(stakeId int64, stakeReveneuPeriodId int64, revenueAmount float64) (insertIndex int64, err error) {

	/*
		TODO

		by a single operation:
			get fk_wallet
			insert stake_revenue
			change balance/tmp_balance wallet
			change balance supernode
			insert internal_tx row
	*/

	return insertIndex, errors.Wrap(err, fmt.Sprintf("db/pg_stake_revenue/InsertStakeRevenue  stakeId: %d, stakeReveneuPeriodId: %d", stakeId, stakeReveneuPeriodId))
}

func (*stakeRevenueInterface) GetStakeReveneuHistory(walletId int64, offset int64, limit int64) (stakeRevenueHists []types.StakeRevenueHist, err error) {
	return stakeRevenueHists, errors.Wrap(err, fmt.Sprintf("db/pg_stake_revenue/GetStakeReveneuHistory  walletId: %d", walletId))
}

func (*stakeRevenueInterface) GetStakeReveneuHistoryCnt(walletId int64) (recCnt int64, err error) {
	return recCnt, errors.Wrap(err, fmt.Sprintf("db/pg_stake_revenue/GetStakeReveneuHistoryCnt  walletId: %d", walletId))
}
