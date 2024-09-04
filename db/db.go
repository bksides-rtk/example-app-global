package db

import "database/sql"

var Db *sql.DB

func DoThing1() {
	Db.Exec("...")
}

func DoThing2() {
	Db.Exec("...")
}
