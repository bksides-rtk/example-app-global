package main

import (
	"github.com/rtk-tickets/example-app-global/db"
	"github.com/rtk-tickets/example-app-global/logging"
)

func main() {
	doThing1()
	doThing2()
}

func doThing1() {
	logging.Infof("doing thing 1")
	db.DoThing1()
}

func doThing2() {
	logging.Infof("doing thing 2")
	db.DoThing2()
}
