package db

import "database/sql"

type DbIface interface {
	Exec(string, ...any) (sql.Result, error)
}

func DoThing1(db DbIface) {
	db.Exec("...")
}

func DoThing2(db DbIface) {
	db.Exec("...")
}

func InitDB() DbIface {
	return &sql.DB{}
}
