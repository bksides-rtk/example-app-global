package main

import (
	dbPkg "github.com/rtk-tickets/example-app-global/db"
	"github.com/rtk-tickets/example-app-global/logging"
)

func main() {
	db := dbPkg.InitDB()
	logger := logging.InitLogger()

	dbPkg.DoThing1(logger, db)

	listings := listings.GetListings()

	dbPkg.DoThing2(logger, db, listings)
}
