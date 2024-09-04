package db

import (
	"database/sql"

	"github.com/rtk-tickets/example-app-global/logging"
	"github.com/rtk-tickets/example-app-global/models"
)

var db *sql.DB

func DoThing1() {
	logging.Info("Doing thing 1")
	db.Exec("...")
}

func DoThing2(listings []models.Listing) {
	logging.Info("Doing thing 2")
	db.Exec("...")
}
