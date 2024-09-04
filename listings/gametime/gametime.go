package gametime

import "github.com/rtk-tickets/example-app-global/models"

type GameTimeListingsService struct{}

func (gtls *GameTimeListingsService) GetListings() []models.Listing {
	return []models.Listing{
		{
			Buy: gtls.Buy,
		},
	}
}

func (gtls *GameTimeListingsService) Buy() {
	// do something
}
