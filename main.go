package main

import (
	"database/sql"

	"github.com/rtk-tickets/common/util/logging"
)

var logger logging.Logger
var db *sql.DB

func main() {
	doThing1()
	doThing2()
}

func doThing1() {
	logger.Infof("doing thing 1")
	db.Exec("...")
}

func doThing2() {
	logger.Infof("doing thing 2")
	db.Exec("...")
}
