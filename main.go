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
	loggingService := logging.InitLoggingService()
	dbService := dbPkg.InitDBService(loggingService)

	dbService.DoThing1()

	listings := listings.GetListings([]listings.ListingsService{
		&gametime.GameTimeListingsService{},
		&seatgeek.SeatGeekListingsService{},
		&ticketmaster.TicketmasterListingsService{},
	})

	dbService.DoThing2(listings)
}
