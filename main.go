package main

import (
	"github.com/rtk-tickets/example-app-global/db"
	"github.com/rtk-tickets/example-app-global/listings"
)

func main() {
	db.DoThing1()

	listings := listings.GetListings()

	db.DoThing2(listings)
}
