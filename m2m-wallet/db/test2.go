package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type Wall struct {
	id      int `db:"id"`
	fkOrgLa int `db:"fk_org_la"`
}

func createWall(dbp *sql.DB) {
	cr, errCr := dbp.Exec(`create table wall (id int, fk_org_la int);`)
	fmt.Println("cr:", cr)
	fmt.Println("errCr:", errCr)
}

func createTemplate(dbp *sql.DB) {
	cr, errCr := dbp.Exec(`create table w (id int);`)
	fmt.Println("cr:", cr)
	fmt.Println("errCr:", errCr)
}

func insertTemplate(dbp *sql.DB) {

	ins, errIns := dbp.Exec(`insert into w (id) values (45) returning id;`)

	fmt.Println("ins:", ins)
	fmt.Println("errIns:", errIns)

}

func updateTemplate(dbp *sql.DB) {

	ins, errIns := dbp.Exec(`update w set id = id +1 ;`)

	fmt.Println("update:", ins)
	fmt.Println("errUpdate:", errIns)

}

func selectTemplate(dbp *sql.DB) {

	query := dbp.QueryRow(`select * from w;`)

	fmt.Println("select:", query)
	// fmt.Println("errSelect:", errIns)
}

func kmyFunc() {

	// db, err := sql.Open("postgres", "m2m_db:@172.18.0.4:5432/m2m_database?sslmode=disable")
	// db, err := sql.Open("postgres", "postgres://postgres@postgres:5432/postgres?sslmode=disable")

	db, err := sql.Open("postgres", "postgres://m2m_db@postgres:5432/m2m_database?sslmode=disable")
	createTemplate(db)
	insertTemplate(db)
	updateTemplate(db)
	selectTemplate(db)

	// db, err := sqlx.Connect("postgres", "m2m_db:@localhost/m2m_database?sslmode=disable")
	// db, err := sqlx.Open("postgres", "m2m_db:@localhost/m2m_database?sslmode=disable")

	// db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=m2m_db password= dbname=m2m_database sslmode=disable")
	// db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=m2m_db dbname=m2m_database sslmode=disable")

	// Connect to a database and verify with a ping.
	// sqlx.Open()   Open is the same as sql.Open, but returns an *sqlx.DB instead.

	fmt.Println(db)
	fmt.Println("err db:", err)

	// ins, errIns := db.Exec(`insert into w (name) values (34);`)

	// // a, err2 := db.Exec(`
	// // insert into wall (
	// // 	id,
	// // 	fk_org_la,
	// // ) values ($1, $2)`,
	// // 	1,
	// // 	2,
	// // )
	// fmt.Println("ins:", ins)
	// fmt.Println("errIns:", errIns)

}

func InsertWall(db *sql.DB, w *Wall) error {

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

// func InsertWall(db *sql.DB, w *Wall) error {

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
