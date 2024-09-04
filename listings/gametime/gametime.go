package gametime

import "github.com/rtk-tickets/example-app-global/models"

type GameTimeListingsService struct{}

func (gtls *GameTimeListingsService) GetListings() []models.Listing {
	return []models.Listing{
		{},
	}
}
