package seatgeek

import "github.com/rtk-tickets/example-app-global/models"

type SeatgeekListingsService struct{}

func (sgls *SeatgeekListingsService) GetListings() []models.Listing {
	return []models.Listing{
		{},
	}
}
