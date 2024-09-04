package listings

import (
	"github.com/rtk-tickets/example-app-global/listings/gametime"
	"github.com/rtk-tickets/example-app-global/listings/seatgeek"
	"github.com/rtk-tickets/example-app-global/listings/ticketmaster"
	"github.com/rtk-tickets/example-app-global/models"
)

var DefaultListingsService = multiListingsService{
	marketplaceSvcs: []ListingsService{
		&seatgeek.SeatGeekListingsService{},
		&ticketmaster.TicketmasterListingsService{},
		&gametime.GameTimeListingsService{},
	},
}

type multiListingsService struct {
	marketplaceSvcs []ListingsService
}

func (mls *multiListingsService) GetListings() []models.Listing {
	ret := []models.Listing{}
	for _, service := range mls.marketplaceSvcs {
		ret = append(ret, service.GetListings()...)
	}

	for _, listing := range ret {
		if shouldBuy(listing) {
			listing.Buy()
		}
	}

	return ret
}

type ListingsService interface {
	GetListings() []models.Listing
}

func shouldBuy(_ models.Listing) bool {
	return true
}
