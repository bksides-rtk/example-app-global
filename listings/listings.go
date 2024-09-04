package listings

import (
	"github.com/rtk-tickets/example-app-global/listings/gametime"
	"github.com/rtk-tickets/example-app-global/listings/seatgeek"
	"github.com/rtk-tickets/example-app-global/listings/ticketmaster"
	"github.com/rtk-tickets/example-app-global/models"
)

func GetListings() []models.Listing {
	ret := gametime.GetListingsGametime()

	ret = append(ret, seatgeek.GetListingsSeatgeek()...)

	ret = append(ret, ticketmaster.GetListingsTicketmaster()...)

	return ret
}
