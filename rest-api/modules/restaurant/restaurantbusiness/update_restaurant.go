package restaurantbusiness

import (
	"context"
	"errors"
	"rest-api/modules/restaurant/restaurantmodel"
)

type UpdateRestaurantStore interface {
	FindDataByCondition(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
	Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdates) error
}

type updateRestaurantBusiness struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantStore(store UpdateRestaurantStore) *updateRestaurantBusiness {
	return &updateRestaurantBusiness{store: store}
}

func (u *updateRestaurantBusiness) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdates) error {
	restaurant, err := u.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if restaurant.Status == 0 {
		return errors.New("data has been deleted")
	}

	if err := u.store.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}
