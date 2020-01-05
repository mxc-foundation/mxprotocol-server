package postgres_db

import (
	"github.com/pkg/errors"
	"github.com/mxc-foundation/mxprotocol-server/m2m/types"
)

type internalTxInterface struct{}

var PgInternalTx internalTxInterface

type FieldStatus string // db:field_status
const (
	ACTIVE FieldStatus = "ACTIVE"
	ARC    FieldStatus = "ARC"
)

func (*internalTxInterface) CreateInternalTxTable() error {
	_, err := PgDB.Exec(`
	DO $$
	BEGIN
		IF 
			NOT EXISTS 
			(SELECT 1 FROM pg_type WHERE typname = 'field_status')
		THEN
			CREATE TYPE FIELD_STATUS AS ENUM (
			'ACTIVE', 'ARC');
		END IF;		
		
		IF 
			NOT EXISTS 
			(SELECT 1 FROM pg_type WHERE typname = 'payment_category')
		THEN
		CREATE TYPE PAYMENT_CATEGORY AS ENUM (
			'DOWNLINK_AGGREGATION',
			'PURCHASE_SUBSCRIPTION',
			'BUY_SUBSCRIPTION',
			'TOP_UP',
			'WITHDRAW',
			'WITHDRAW_FEE_SN_INCOME',
			'DOWNLINK_AGG_SN_INCOME',
			'STAKE_REVENUE',
			'INSERT_STAKE',        
			'UNSTAKE'              
		);
		END IF;

		CREATE TABLE IF NOT EXISTS internal_tx (
			id SERIAL PRIMARY KEY,
			fk_wallet_sender INT REFERENCES wallet(id) NOT NULL,
			fk_wallet_receiver INT REFERENCES wallet(id) NOT NULL,
			payment_cat PAYMENT_CATEGORY,
			tx_internal_ref INT NOT NULL,
			value NUMERIC(28,18),
			time_tx  TIMESTAMP,
			CONSTRAINT payment_cat_tx_internal_ref UNIQUE (payment_cat, tx_internal_ref)
		);
	END$$;
	`)
	return errors.Wrap(err, "db/CreateInternalTxTable")
}

func (*internalTxInterface) InsertInternalTx(it types.InternalTx) (insertIndex int64, err error) {
	err = PgDB.QueryRow(`
	INSERT INTO internal_tx (
		fk_wallet_sender,
		fk_wallet_receiver,
		payment_cat,
		tx_internal_ref,
		value,
		time_tx )
		VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6)
		RETURNING id;
	`,
		it.FkWalletSender,
		it.FkWalletRcvr,
		it.PaymentCat,
		it.TxInternalRef,
		it.Value,
		it.TimeTx).Scan(&insertIndex)

	return insertIndex, errors.Wrap(err, "db/InsertInternalTx")
}
