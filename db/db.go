package db

import (
	"database/sql"

	"github.com/rtk-tickets/example-app-global/logging"
	"github.com/rtk-tickets/example-app-global/models"
)

type DBService struct {
	loggingService logging.LoggingService
	db             DbIface
}

func (dbs *DBService) WithDB(db DbIface) *DBService {
	dbs.db = db
	return dbs
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

func (dbs *DBService) DoThing2(listings []models.Listing) {
	dbs.loggingService.Info("Doing thing 2")

	var newDbService = dbs

	switch db := dbs.db.(type) {
	case *sql.DB:
		tx, err := db.Begin()
		if err != nil {
			return
		}
		newDbService = dbs.WithDB(tx)
	}

	newDbService.db.Exec("...")
	newDbService.DoThing1()
}
