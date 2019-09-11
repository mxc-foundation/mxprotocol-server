package postgres_db

import (
	"time"

	"github.com/pkg/errors"
)

type internalTxInterface struct{}

var PgInternalTx internalTxInterface

type FieldStatus string // db:field_status
const (
	ACTIVE FieldStatus = "ACTIVE"
	ARC    FieldStatus = "ARC"
)

type PaymentCategory string // db:payment_category
const (
	UPLINK                PaymentCategory = "UPLINK"
	DOWNLINK              PaymentCategory = "DOWNLINK"
	PURCHASE_SUBSCRIPTION PaymentCategory = "PURCHASE_SUBSCRIPTION"
	GATEWAY_INCOME        PaymentCategory = "GATEWAY_INCOME"
	TOP_UP                PaymentCategory = "TOP_UP"
	WITHDRAW              PaymentCategory = "WITHDRAW"
)

type InternalTx struct {
	Id             int64     `db:"id"`
	FkWalletSender int64     `db:"fk_wallet_sender"`
	FkWalletRcvr   int64     `db:"fk_wallet_receiver"`
	PaymentCat     string    `db:"payment_cat"`
	TxInternalRef  int64     `db:"tx_internal_ref"` // reference to the id of corresponding table to PaymentCat
	Value          float64   `db:"value"`
	TimeTx         time.Time `db:"timestamp"`
}

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
			'UPLINK',
			'DOWNLINK',
			'PURCHASE_SUBSCRIPTION',
			'GATEWAY_INCOME',
			'TOP_UP',
			'WITHDRAW'
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

func (*internalTxInterface) InsertInternalTx(it InternalTx) (insertIndex int64, err error) {
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
