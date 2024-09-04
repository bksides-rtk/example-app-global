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

type DbIface interface {
	Exec(string, ...any) (sql.Result, error)
}

func (dbs *DBService) DoThing1() {
	dbs.loggingService.Info("Doing thing 1")
	dbs.db.Exec("...")
}

func (dbs *DBService) DoThing2(listings []models.Listing) {
	dbs.loggingService.Info("Doing thing 2")
	dbs.db.Exec("...")
}

func InitDBService(loggingService logging.LoggingService) DBService {
	return DBService{
		loggingService: loggingService,
		db:             &sql.DB{},
	}
}
