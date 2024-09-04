package main

import (
	dbPkg "github.com/rtk-tickets/example-app-global/db"
	"github.com/rtk-tickets/example-app-global/logging"
)

func main() {
	loggingService := logging.InitLoggingService()
	dbService := dbPkg.InitDBService(loggingService)

	dbService.DoThing1()
	dbService.DoThing2()
}
