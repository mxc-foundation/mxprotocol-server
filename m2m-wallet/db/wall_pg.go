package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type Wall struct {
	id      int `db:"id"`
	fkOrgLa int `db:"fk_org_la"`
}

func kmyFunc() {
	fmt.Println("my func2")

	// db, err := sqlx.Connect("postgres", "m2m_db:@localhost/m2m_database?sslmode=disable")
	db, err := sqlx.Open("postgres", "m2m_db:@localhost/m2m_database?sslmode=disable")

	// db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=m2m_db password= dbname=m2m_database sslmode=disable")
	// db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=m2m_db dbname=m2m_database sslmode=disable")

	// Connect to a database and verify with a ping.
	// sqlx.Open()   Open is the same as sql.Open, but returns an *sqlx.DB instead.

	fmt.Println(db)
	fmt.Println("err:", err)

	a, err2 := db.Exec(`
	insert into wall (
		id,
		fk_org_la,
	) values ($1, $2)`,
		1,
		2,
	)
	fmt.Println("a:", a)
	fmt.Println("err2:", err2)

}

func InsertWall(db sqlx.Ext, w *Wall) error {

	// if err := w.Validate(); err != nil {
	// 	return errors.Wrap(err, "validate error")
	// }

	// dpID, err := uuid.NewV4()
	// if err != nil {
	// 	return errors.Wrap(err, "new uuid v4 error")
	// }

	// now := time.Now()  => based on the time of the server
	// w.CreatedAt = now
	// w.UpdatedAt = now

	_, err := db.Exec(`
        insert into wall (
            id,
            fk_org_la,
        ) values ($1, $2)`,
		w.id,
		w.fkOrgLa,
	)
	if err != nil {
		return handlePSQLError(Insert, err, "insert error")
	}

	log.WithFields(log.Fields{
		"wallet_id:": w.id,
	}).Info("wallet created")

	return nil
}

// func InsertWall(db sqlx.Ext, w *Wall) error {

// 	if err := w.Validate(); err != nil {
// 		return errors.Wrap(err, "validate error")
// 	}

// 	// now := time.Now()  => based on the time of the server
// 	// w.CreatedAt = now
// 	// w.UpdatedAt = now

// 	_, err := db.Exec(`
// 	create table wall (id int, fk_org_la int);
// 	`
// 	)
// 	if err != nil {
// 		return handlePSQLError(Insert, err, "insert error")
// 	}

// 	log.WithFields(log.Fields{
// 		"wallet_id:": w.id,
// 	}).Info("wallet created")

// 	return nil
// }
