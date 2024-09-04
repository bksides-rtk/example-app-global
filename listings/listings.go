package listings

import (
	"github.com/rtk-tickets/example-app-global/listings/gametime"
	"github.com/rtk-tickets/example-app-global/listings/seatgeek"
	"github.com/rtk-tickets/example-app-global/listings/ticketmaster"
	"github.com/rtk-tickets/example-app-global/models"
)

type ListingsService interface {
	GetListings() []models.Listing
}

var ListingsServices []ListingsService = []ListingsService{
	&gametime.GameTimeListingsService{},
	&seatgeek.SeatGeekListingsService{},
	&ticketmaster.TicketmasterListingsService{},
}

func GetListings() []models.Listing {
	ret := []models.Listing{}
	for _, service := range ListingsServices {
		ret = append(ret, service.GetListings()...)
	}

	return ret
}
