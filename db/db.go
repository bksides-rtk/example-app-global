package db

import "database/sql"

var db *sql.DB

func DoThing1() {
	db.Exec("...")
}

func DoThing2() {
	db.Exec("...")
}
