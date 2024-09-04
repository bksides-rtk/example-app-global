package db

import (
	"database/sql"

	"github.com/rtk-tickets/example-app-global/logging"
)

type DBService struct {
	loggingService logging.LoggingService
	db             DbIface
}

var DefaultDBService = DBService{
	loggingService: logging.DefaultLoggingService,
	db:             &sql.DB{},
}

type DbIface interface {
	Exec(string, ...any) (sql.Result, error)
}

func (dbs *DBService) DoThing1() {
	dbs.loggingService.Info("Doing thing 1")
	dbs.db.Exec("...")
}

func (dbs *DBService) DoThing2() {
	dbs.loggingService.Info("Doing thing 2")
	dbs.db.Exec("...")
}
