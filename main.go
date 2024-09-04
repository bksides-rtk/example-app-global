package main

import (
	dbPkg "github.com/rtk-tickets/example-app-global/db"
	"github.com/rtk-tickets/example-app-global/listings"
	"github.com/rtk-tickets/example-app-global/listings/gametime"
	"github.com/rtk-tickets/example-app-global/listings/seatgeek"
	"github.com/rtk-tickets/example-app-global/listings/ticketmaster"
	"github.com/rtk-tickets/example-app-global/logging"
)

func main() {
	db := dbPkg.InitDB()
	logger := logging.InitLogger()

	dbPkg.DoThing1(logger, db)

	listings := listings.GetListings([]listings.ListingsService{
		&gametime.GameTimeListingsService{},
		&seatgeek.SeatGeekListingsService{},
		&ticketmaster.TicketmasterListingsService{},
	})

	dbPkg.DoThing2(logger, db, listings)
}
