package gametime

import "github.com/rtk-tickets/example-app-global/models"

type GametimeListingsService struct{}

var gametimeListingsServiceInstance GametimeListingsService

func (gtls *GametimeListingsService) GetListings() []models.Listing {
	return []models.Listing{
		{},
	}
}
