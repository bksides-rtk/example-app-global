package listings

import (
	gametimeBuyer "github.com/rtk-tickets/example-app-global/buyers/gametime"
	seatgeekBuyer "github.com/rtk-tickets/example-app-global/buyers/seatgeek"
	ticketmasterBuyer "github.com/rtk-tickets/example-app-global/buyers/ticketmaster"
	"github.com/rtk-tickets/example-app-global/listings/gametime"
	"github.com/rtk-tickets/example-app-global/listings/seatgeek"
	"github.com/rtk-tickets/example-app-global/listings/ticketmaster"
	"github.com/rtk-tickets/example-app-global/models"
)

func GetListings() []models.Listing {
	ret := gametime.GetListingsGametime()

	ret = append(ret, seatgeek.GetListingsSeatgeek()...)

	ret = append(ret, ticketmaster.GetListingsTicketmaster()...)

	for _, listing := range ret {
		if shouldBuy(listing) {
			switch listing.Marketplace {
			case "gametime":
				gametimeBuyer.BuyGametimeListing(listing.MarketplaceID)
			case "seatgeek":
				seatgeekBuyer.BuySeatgeekListing(listing.MarketplaceID)
			case "ticketmaster":
				ticketmasterBuyer.BuyTicketmasterListing(listing.MarketplaceID)
			}
		}
	}

	return ret
}

func shouldBuy(_ models.Listing) bool {
	return true
}
