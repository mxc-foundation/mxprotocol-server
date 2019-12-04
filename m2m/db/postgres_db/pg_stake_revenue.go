package postgres_db

import (
	"fmt"
	"time"

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

func (*stakeRevenueInterface) CreateStakeRevenueFunctions() error {
	_, err := PgDB.Exec(`

	CREATE OR REPLACE FUNCTION stake_revenue_exec (


		v_fk_stake_revenue_period INT,
		v_fk_stake INT,		
		v_revenue_amount NUMERIC(28,18),
		v_time TIMESTAMP,
		v_fk_wallet_supernode_income INT,
		v_fk_wallet_user INT,
		v_payment_cat PAYMENT_CATEGORY
	) RETURNS  NUMERIC(28,18)
	LANGUAGE plpgsql
	AS $$

	declare stake_rev_id INT;
	declare updated_wlt_balance NUMERIC(28,18);

	BEGIN

	

	UPDATE
		wallet 
	SET
		balance = balance + v_revenue_amount,
		tmp_balance = tmp_balance + v_revenue_amount
		
	WHERE
		id = v_fk_wallet_user
	RETURNING balance INTO updated_wlt_balance
	;

	UPDATE
		wallet 
	SET
		balance = balance - v_revenue_amount,
		tmp_balance = tmp_balance - v_revenue_amount
		
	WHERE
		id = v_fk_wallet_supernode_income
	;
	 

	INSERT INTO
		stake_revenue (
		fk_stake_revenue_period,
		fk_stake ,
		r_fk_wallet,
		revenue_amount ,
		updated_balance )
	VALUES
		(
		v_fk_stake_revenue_period ,
		v_fk_stake,
		v_fk_wallet_user,		
		v_revenue_amount,
		updated_wlt_balance
	)
	RETURNING id INTO stake_rev_id;


INSERT INTO internal_tx (
		fk_wallet_sender,
		fk_wallet_receiver,
		payment_cat,
		tx_internal_ref,
		value,
		time_tx )
	VALUES (
		v_fk_wallet_supernode_income,
		v_fk_wallet_user,
		v_payment_cat,
		stake_rev_id,
		v_revenue_amount,
		v_time)
		;

	RETURN stake_rev_id;

	END;
	$$;

	`)

	return errors.Wrap(err, "db/pg_stake_revenue/CreateStakeRevenueFunctions")
}

func (*stakeRevenueInterface) InsertStakeRevenue(stakeId int64, stakeRevenuePeriodId int64, revenueAmount float64) (insertIndex int64, err error) {

	userWalletId, err := PgStake.GetStakeWalletId(stakeId)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("db/pg_stake_revenue/InsertStakeRevenue  stakeId: %d, stakeRevenuePeriodId: %d;  Unable to get walletId! ", stakeId, stakeRevenuePeriodId))
	}

	supernodeIncomeWltId, err := PgWallet.GetWalletIdSuperNodeIncome()
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("db/pg_stake_revenue/InsertStakeRevenue  stakeId: %d, stakeRevenuePeriodId: %d;  Unable to get WalletIdSuperNodeIncome! ", stakeId, stakeRevenuePeriodId))
	}

	err = PgDB.QueryRow(`
	SELECT
		stake_revenue_exec (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		);`,

		stakeRevenuePeriodId,
		stakeId,
		revenueAmount,
		time.Now().UTC(),
		supernodeIncomeWltId,
		userWalletId,
		types.STAKE_REVENUE,
	).Scan(&insertIndex)

	return insertIndex, errors.Wrap(err, fmt.Sprintf("db/pg_stake_revenue/InsertStakeRevenue  stakeId: %d, stakeRevenuePeriodId: %d", stakeId, stakeRevenuePeriodId))
}

func (*stakeRevenueInterface) GetStakeRevenueHistory(walletId int64, offset int64, limit int64) (stakeRevenueHists []types.StakeRevenueHist, err error) {

	rows, err := PgDB.Query(
		`SELECT 
			s.fk_wallet, 
			s.amount, 
			s.status, 
			s.start_stake_time, 
			unstake_time, 
			srp.staking_period_start, 
			srp.staking_period_end, 
			supernode_income, 
			income_to_stake_portion, 
			sr.revenue_amount, 
			sr.updated_balance
		FROM 
			stake s 
		LEFT JOIN  
			stake_revenue sr
		INNER JOIN
			stake_revenue_period srp
		ON 
			sr.fk_stake_revenue_period = srp.id  
		ON  
			s.id = sr.fk_stake_revenue_period
		WHERE 
			s.fk_wallet = $1
		ORDER BY 
			sr.id DESC,
			unstake_time DESC, 
			s.start_stake_time DESC
		LIMIT $2 
		OFFSET $3
		;
		
	;`, walletId, limit, offset)

	if err != nil {
		return stakeRevenueHists, errors.Wrap(err, fmt.Sprintf("db/pg_stake_revenue/GetStakeRevenueHistory  walletId: %d", walletId))
	}

	defer rows.Close()

	srh := types.StakeRevenueHist{}

	for rows.Next() {
		rows.Scan(
			&srh.WalletId,
			&srh.StakeAmount,
			&srh.StakeStatus,
			&srh.StartStakeTime,
			&srh.UnstakeTime,
			&srh.StakingPeriodStart,
			&srh.StakingPeriodEnd,
			&srh.SuperNodeIncome,
			&srh.IncomeToStakePortion,
			&srh.RevenueAmount,
			&srh.UpdatedBalance,
		)

		stakeRevenueHists = append(stakeRevenueHists, srh)
	}

	return stakeRevenueHists, errors.Wrap(err, fmt.Sprintf("db/pg_stake_revenue/GetStakeRevenueHistory  walletId: %d", walletId))
}

func (*stakeRevenueInterface) GetStakeRevenueHistoryCnt(walletId int64) (recCnt int64, err error) {

	err = PgDB.QueryRow(`
		SELECT 
			COUNT(*) 
		FROM 
			stake s 
		LEFT JOIN  
			stake_revenue sr 
		ON  
			s.id = sr.fk_stake_revenue_period
		WHERE 
			s.fk_wallet = $1
		;
	`, walletId).Scan(&recCnt)

	return recCnt, errors.Wrap(err, fmt.Sprintf("db/pg_stake_revenue/GetStakeRevenueHistoryCnt  walletId: %d", walletId))
}
