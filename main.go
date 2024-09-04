package main

import (
	dbPkg "github.com/rtk-tickets/example-app-global/db"
	"github.com/rtk-tickets/example-app-global/logging"
)

func main() {
	db := dbPkg.InitDB()
	logger := logging.InitLogger()

	doThing1(db, logger)
	doThing2(db, logger)
}

func doThing1(db dbPkg.DbIface, logger logging.Logger) {
	logging.Info(logger, "doing thing 1")
	dbPkg.DoThing1(db)
}

func doThing2(db dbPkg.DbIface, logger logging.Logger) {
	logging.Info(logger, "doing thing 2")
	dbPkg.DoThing2(db)
}
