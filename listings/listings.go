package listings

import (
	"github.com/rtk-tickets/example-app-global/models"
)

type ListingsService interface {
	GetListings() []models.Listing
}

func GetListings(listingsServices []ListingsService) []models.Listing {
	ret := []models.Listing{}
	for _, service := range listingsServices {
		ret = append(ret, service.GetListings()...)
	}

	return ret
}
