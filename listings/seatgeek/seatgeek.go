package seatgeek

import "github.com/rtk-tickets/example-app-global/models"

type SeatGeekListingsService struct{}

func (sgls *SeatGeekListingsService) GetListings() []models.Listing {
	return []models.Listing{
		{
			Buy: sgls.Buy,
		},
	}
}

func (sgls *SeatGeekListingsService) Buy() {
	// do something
}
