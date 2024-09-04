package main

import (
	"github.com/rtk-tickets/example-app-global/db"
	"github.com/rtk-tickets/example-app-global/listings"
)

func main() {
	db.DefaultDBService.DoThing1()

	listings := listings.GetListings()

	db.DefaultDBService.DoThing2(listings)
}
