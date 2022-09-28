package restaurantbusiness

import (
	"context"
	"rest-api/common"
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
		if err == common.ErrorRecordNotFound {
			return nil, common.ErrCantGetEntity(restaurantmodel.EntityName, err)
		}
		return nil, common.ErrCantGetEntity(restaurantmodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	return data, nil
}
