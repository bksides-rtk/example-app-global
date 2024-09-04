package db

import "database/sql"

type dbIface interface {
	Exec(string) (sql.Result, error)
}

var Db dbIface

func DoThing1() {
	Db.Exec("...")
}

func DoThing2() {
	Db.Exec("...")
}
