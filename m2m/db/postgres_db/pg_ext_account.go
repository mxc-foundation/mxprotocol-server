package postgres_db

import (
	"time"

	"github.com/ethereum/go-ethereum/log"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type ExtAccount struct {
	Id                 int64     `db:"id"`
	FkWallet           int64     `db:"fk_wallet"`
	FkExtCurrency      int64     `db:"fk_ext_currency"`
	AccountAdr         string    `db:"account_adr"`
	InsertTime         time.Time `db:"insert_time"`
	Status             string    `db:"status"`
	LatestCheckedBlock int64     `db:"latest_checked_block"`
}

type ExtAccountHistRet struct {
	AccountAdr     string
	InsertTime     time.Time
	Status         string
	ExtCurrencyAbv string
}

func  CreateExtAccountTable() error {
	_, err := PgDB.Exec(`
		CREATE TABLE IF NOT EXISTS ext_account (
			id SERIAL PRIMARY KEY,
			fk_wallet INT REFERENCES wallet(id) NOT NULL,
			fk_ext_currency INT REFERENCES ext_currency (id) NOT NULL,
			account_adr varchar(128) NOT NULL,
			insert_time TIMESTAMP NOT NULL,
			status FIELD_STATUS NOT NULL,
			latest_checked_block INT DEFAULT 0
		);
		
	`)
	return errors.Wrap(err, "db/CreateExtAccountTable")
}

func  InsertExtAccount(ea ExtAccount) (insertIndex int64, err error) {

	alreadyExist, errAlreadyExist := alreadyExistActiveAcnt(ea.AccountAdr, ea.FkExtCurrency)
	if errAlreadyExist != nil {
		return 0, errors.Wrap(errAlreadyExist, "db/InsertExtAccount")
	}
	if alreadyExist {
		return 0, errors.Wrap(errors.New("Account Address is already active!"), "db/InsertExtAccount")
	}

	err = PgDB.QueryRow(`
	INSERT INTO ext_account (
			fk_wallet,
			fk_ext_currency,
			account_adr,
			insert_time,
			status,
			latest_checked_block)
		VALUES (
			$1, $2,	$3,	$4,	'ACTIVE',	$5
		)
		RETURNING id;
	`,
		ea.FkWallet,
		ea.FkExtCurrency,
		ea.AccountAdr,
		ea.InsertTime,
		ea.LatestCheckedBlock).Scan(&insertIndex)

	if err == nil {
		ea.Id = insertIndex
		err2 := changeStatus2ArcOldRowExtAcnt(ea)
		if err2 != nil {
			return insertIndex, errors.Wrap(err, "db/InsertExtAccount")
		}
	}

	return insertIndex, errors.Wrap(err, "db/InsertExtAccount")
}

func  alreadyExistActiveAcnt(acntAdr string, extCurrId int64) (bool, error) {

	var nRow int64

	err := PgDB.QueryRow(`
		select 
			COALESCE(SUM(1),0) as num_rep
		FROM
			ext_account 
		WHERE
			account_adr = $1
		AND
			status = 'ACTIVE'		
		AND 
			fk_ext_currency = $2
		;
	
	`, acntAdr, extCurrId).Scan(&nRow)

	res := nRow != 0

	return res, errors.Wrap(err, "db/alreadyExistActiveAcnt")
}

func  changeStatus2ArcOldRowExtAcnt(ea ExtAccount) (err error) {
	_, err = PgDB.Exec(`
	UPDATE 
		ext_account 
	SET 
		status = 'ARC'
	WHERE
		fk_wallet = $1 
		AND
		fk_ext_currency = $2
		AND
		id <> $3   
	;
	`,
		ea.FkWallet,
		ea.FkExtCurrency,
		ea.Id)

	return errors.Wrap(err, "db/changeStatus2ArcOldRowExtAcnt")
}

func  GetSuperNodeExtAccountAdr(extCurrAbv string) (string, error) {

	var res string

	err := PgDB.QueryRow(`
		select 
			ea.account_adr
		from
			wallet w ,ext_account ea,ext_currency ec
		WHERE
			w.id = ea.fk_wallet
			AND
			w.type = 'SUPER_ADMIN'
			AND
			ea.fk_ext_currency = ec.id
			AND
			ea.status = 'ACTIVE'
			AND
			ec.abv = $1
		ORDER BY ea.id DESC  
		LIMIT 1 
		;
	`, extCurrAbv).Scan(&res)

	return res, errors.Wrap(err, "db/GetSuperNodeExtAccountAdr")
}

func  GetSuperNodeExtAccountId(extCurrAbv string) (int64, error) {
	var res int64

	err := PgDB.QueryRow(`
		select 
			ea.id
		from
			wallet w ,ext_account ea,ext_currency ec
		WHERE
			w.id = ea.fk_wallet
			AND
			w.type = 'SUPER_ADMIN' 
			AND
			ea.fk_ext_currency = ec.id
			AND
			ea.status = 'ACTIVE'
			AND
			ec.abv = $1
		ORDER BY ea.id DESC  
		LIMIT 1 
		;
	`, extCurrAbv).Scan(&res)

	return res, errors.Wrap(err, "db/GetSuperNodeExtAccountId")
}

func  GetUserExtAccountAdr(walletId int64, extCurrAbv string) (string, error) {

	var res string

	err := PgDB.QueryRow(`
		select 
			ea.account_adr
		from
			wallet w,ext_account ea,ext_currency ec
		WHERE
			w.id = ea.fk_wallet AND
			w.type = 'USER' AND
			ea.fk_ext_currency = ec.id AND
			ea.status = 'ACTIVE' AND
			w.id = $1 AND
			ec.abv = $2
		ORDER BY ea.id DESC 
		LIMIT 1 
		;
	
	`, walletId, extCurrAbv).Scan(&res)

	return res, errors.Wrap(err, "db/GetUserExtAccountAdr")
}

func  GetUserExtAccountId(walletId int64, extCurrAbv string) (int64, error) {

	var res int64

	err := PgDB.QueryRow(`
		select 
			ea.id
		from
			wallet w,ext_account ea,ext_currency ec
		WHERE
			w.id = ea.fk_wallet AND
			w.type = 'USER' AND
			ea.fk_ext_currency = ec.id AND
			ea.status = 'ACTIVE' AND
			w.id = $1 AND
			ec.abv = $2
		ORDER BY ea.id DESC 
		LIMIT 1 
		;
	
	`, walletId, extCurrAbv).Scan(&res)

	return res, errors.Wrap(err, "db/GetUserExtAccountId")
}

func  GetExtAccountIdByAdr(acntAdr string, extCurrAbv string) (int64, error) {

	var res int64

	err := PgDB.QueryRow(`
		select 
			ea.id
		from
			ext_account ea, ext_currency ec
		WHERE
			ea.fk_ext_currency = ec.id
		AND
			ea.account_adr = $1
		AND
			ea.status = 'ACTIVE'		
		AND 
			ec.abv = $2
		ORDER BY ea.id DESC 
		LIMIT 1 
		;
	
	`, acntAdr, extCurrAbv).Scan(&res)

	return res, errors.Wrap(err, "db/GetExtAccountIdByAdr")
}

func  GetLatestCheckedBlock(extAcntId int64) (int64, error) {

	var res int64

	err := PgDB.QueryRow(`
		SELECT 
			latest_checked_block 
		FROM 
			 ext_account 
		WHERE
			 id = $1
	
	`, extAcntId).Scan(&res)

	return res, errors.Wrap(err, "db/GetLatestCheckedBlock")
}

func  UpdateLatestCheckedBlock(extAcntId int64, updatedBlockNum int64) error {

	_, err := PgDB.Exec(`
		UPDATE ext_account 
		SET 
		latest_checked_block = $1
		WHERE
		id = $2;
	
	`, updatedBlockNum, extAcntId)

	return errors.Wrap(err, "db/UpdateLatestCheckedBlock")
}

func  GetExtAcntHist(walletId int64, offset int64, limit int64) ([]ExtAccountHistRet, error) {

	rows, err := PgDB.Query(
		`select
			ea.account_adr,
			ea.insert_time,
			ea.status,
			ec.abv AS ext_currency_abv
		from
			ext_account ea,
			wallet w, 
			ext_currency ec
		WHERE
			ea.fk_ext_currency = ec.id AND
			ea.fk_wallet = w.id AND
			w.id = $1
		ORDER BY ea.insert_time DESC
		LIMIT $2
		OFFSET $3
		;`, walletId, limit, offset)

	defer rows.Close()

	res := make([]ExtAccountHistRet, 0)
	var extAcntVal ExtAccountHistRet
	var insertTime string

	for rows.Next() {
		rows.Scan(
			&extAcntVal.AccountAdr,
			&insertTime,
			&extAcntVal.Status,
			&extAcntVal.ExtCurrencyAbv,
		)
		if conTime, errTime := time.Parse(timeLayout, insertTime); errTime == nil {
			extAcntVal.InsertTime = conTime
		} else {
			log.Debug("db/GetExtAcntHist Unable to convert time: ", err)
		}

		res = append(res, extAcntVal)
	}
	return res, errors.Wrap(err, "db/GetExtAcntHist")
}

func  GetExtAcntHistRecCnt(walletId int64) (recCnt int64, err error) {

	err = PgDB.QueryRow(`
		SELECT
			COUNT(*)
		FROM
			ext_account ea,
			wallet w
		WHERE
			ea.fk_wallet = w.id AND
			w.id = $1
		;
	`, walletId).Scan(&recCnt)

	return recCnt, errors.Wrap(err, "db/GetExtAcntHistRecCnt")
}
