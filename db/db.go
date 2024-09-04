package db

import (
	"database/sql"

	"github.com/rtk-tickets/example-app-global/logging"
)

type DbIface interface {
	Exec(string, ...any) (sql.Result, error)
}

func DoThing1(logger logging.Logger, db DbIface) {
	logging.Info(logger, "Doing thing 1")
	db.Exec("...")
}

func DoThing2(logger logging.Logger, db DbIface) {
	logging.Info(logger, "Doing thing 2")
	db.Exec("...")
}

func InitDB() DbIface {
	return &sql.DB{}
}
