package db

import (
	"database/sql"
	"sync"

	"github.com/rtk-tickets/example-app-global/logging"
	"github.com/rtk-tickets/example-app-global/models"
)

type dbIface interface {
	Exec(string, ...any) (sql.Result, error)
}

var mtx = &sync.Mutex{}
var Db dbIface

func DoThing1() {
	logging.Info("Doing thing 1")
	Db.Exec("...")
}

func DoThing2(listings []models.Listing) {
	logging.Info("Doing thing 2")

	mtx.Lock()
	defer mtx.Unlock()

	switch db := Db.(type) {
	case *sql.DB:
		defer func() { Db = db }()

		var err error
		Db, err = db.Begin()
		if err != nil {
			return
		}
	}

	Db.Exec("...")
	DoThing1()
}
