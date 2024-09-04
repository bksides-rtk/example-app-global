package ticketmaster

import "github.com/rtk-tickets/example-app-global/models"

type TicketmasterListingsService struct{}

func (tmls *TicketmasterListingsService) GetListings() []models.Listing {
	return []models.Listing{
		{},
	}
}
