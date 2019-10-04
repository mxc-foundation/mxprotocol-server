package postgres_db

import (
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type walletInterface struct{}

var PgWallet walletInterface

type wallet struct {
	Id      int64  `db:"id"`
	FkOrgLa int64  `db:"fk_org_la"`
	TypeW   string `db:"type"`
	// Balance is updated during the aggregations (containing internal_tx reference)
	Balance float64 `db:"balance"`
	// Tmp balance is updated per transactions and is uesd while the balanc is not updated.
	// During the aggregation, TmpBalance will get updated to value of Balance
	TmpBalance float64 `db:"tmp_balance"`
}

func (*walletInterface) CreateWalletTable() error {
	_, err := PgDB.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS 
				(SELECT 1 FROM pg_type WHERE typname = 'wallet_type') 
			THEN
				CREATE TYPE WALLET_TYPE AS ENUM (
					'SUPER_ADMIN',
 					'USER'
		);
		END IF;
			CREATE TABLE IF NOT EXISTS wallet (
	
			id SERIAL PRIMARY KEY,
			fk_org_la INT UNIQUE NOT NULL, -- foreign_key LoRa app server DB
			type WALLET_TYPE NOT NULL,
			balance NUMERIC(28,18) NOT NULL DEFAULT 0,
			tmp_balance NUMERIC(28,18) DEFAULT 0
		);

		END$$;
		
	`)
	return errors.Wrap(err, "db/CreateWalletTable")
}

func (*walletInterface) InsertWallet(orgId int64, walletType types.WalletType) (insertIndex int64, err error) {
	w := wallet{
		FkOrgLa:    orgId,
		TypeW:      string(walletType),
		Balance:    0.0,
		TmpBalance: 0.0,
	}

	err = PgDB.QueryRow(`
		INSERT INTO wallet (
			fk_org_la ,
			type,
			balance,
			tmp_balance ) 
		VALUES 
			($1,$2,$3,$4)
		RETURNING id ;
	`,
		w.FkOrgLa,
		w.TypeW,
		w.Balance,
		w.TmpBalance).Scan(&insertIndex)

	// fmt.Println(val, err)
	return insertIndex, errors.Wrap(err, "db/InsertWallet")
}

func (*walletInterface) GetWalletIdFromOrgId(orgIdLora int64) (int64, error) {
	id := int64(0)
	err := PgDB.QueryRow(
		`SELECT id
		FROM wallet
		WHERE
			fk_org_la = $1;`,
		orgIdLora).Scan(&id)

	return id, errors.Wrap(err, "db/GetWalletIdFromOrgId")
}

func (*walletInterface) GetWalletBalance(walletId int64) (float64, error) {
	balance := float64(0)
	err := PgDB.QueryRow(
		`SELECT balance
		FROM wallet
		WHERE
			id = $1;`,
		walletId).Scan(&balance)

	return balance, errors.Wrap(err, "db/GetWalletBalance")
}

func (*walletInterface) GetWalletTmpBalance(walletId int64) (float64, error) {
	balance := float64(0)
	err := PgDB.QueryRow(
		`SELECT tmp_balance
		FROM wallet
		WHERE
			id = $1;`,
		walletId).Scan(&balance)

	return balance, errors.Wrap(err, "db/GetWalletTmpBalance")
}

func (*walletInterface) GetWalletIdofActiveAcnt(acntAdr string, externalCur string) (walletId int64, err error) {

	err = PgDB.QueryRow(
		`select 
			w.id as wallet_id 
			from
			wallet w,ext_account ea,ext_currency ec
		WHERE
			w.id = ea.fk_wallet AND
			w.type = 'USER' AND
			ea.fk_ext_currency = ec.id AND
			ea.status = 'ACTIVE' AND
			account_adr = $1 AND
			ec.abv = $2 
		ORDER BY ea.id DESC 
		LIMIT 1 
		;`, acntAdr, externalCur).Scan(&walletId)

	return walletId, errors.Wrap(err, "db/GetWalletIdofActiveAcnt")
}

func getWalletIdofActiveAcntSuperAdmin(acntAdr string, externalCur string) (walletId int64, err error) {

	err = PgDB.QueryRow(
		`select 
			w.id as wallet_id 
			from
			wallet w,ext_account ea,ext_currency ec
		WHERE
			w.id = ea.fk_wallet AND
			w.type = 'SUPER_ADMIN' AND
			ea.fk_ext_currency = ec.id AND
			ea.status = 'ACTIVE' AND
			account_adr = $1 AND
			ec.abv = $2 
		ORDER BY ea.id DESC 
		LIMIT 1 
		;`, acntAdr, externalCur).Scan(&walletId)

	return walletId, errors.Wrap(err, "db/getWalletIdofActiveAcntSuperAdmin")
}

func (*walletInterface) GetWalletIdSuperNode() (walletId int64, err error) {

	err = PgDB.QueryRow(
		`SELECT
			id
		FROM
			wallet 
		WHERE
			type = 'SUPER_ADMIN' 
			ORDER BY id DESC 
			LIMIT 1  -- altough only one super node exists
		;`).Scan(&walletId)

	return walletId, errors.Wrap(err, "db/GetWalletIdSuperNode")
}

func (*walletInterface) UpdateBalanceByWalletId(walletId int64, newBalance float64) error {
	_, err := PgDB.Exec(`
	UPDATE
		wallet 
	SET
		balance = $1
	WHERE
		id = $2
	;
	`, newBalance, walletId)

	return errors.Wrap(err, "db/UpdateBalanceByWalletId")
}

//ToDo
func (*walletInterface) TmpBalanceUpdatePktTx(dvId, gwId int64, amount float64) error {
	return nil
}
