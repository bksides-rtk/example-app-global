package db

import (
	"database/sql"

	"github.com/rtk-tickets/example-app-global/logging"
	"github.com/rtk-tickets/example-app-global/models"
)

type dbIface interface {
	Exec(string) (sql.Result, error)
}

var Db dbIface

func DoThing1() {
	logging.Info("Doing thing 1")
	Db.Exec("...")
}

func DoThing2(listings []models.Listing) {
	logging.Info("Doing thing 2")
	Db.Exec("...")
}
