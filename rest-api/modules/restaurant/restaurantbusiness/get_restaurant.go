package restaurantbusiness

import (
	"context"
	"errors"
	"rest-api/modules/restaurant/restaurantmodel"
)

type GetRestaurantStore interface {
	FindDataByCondition(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBusiness struct {
	store GetRestaurantStore
}

func NewGetRestaurantBusiness(store GetRestaurantStore) *getRestaurantBusiness {
	return &getRestaurantBusiness{store: store}
}

func (g *getRestaurantBusiness) GetRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	data, err := g.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	if data.Status == 0 {
		return nil, errors.New("restaurant has been deleted")
	}

	return data, nil
}
