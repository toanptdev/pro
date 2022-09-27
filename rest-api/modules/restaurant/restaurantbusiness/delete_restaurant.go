package restaurantbusiness

import (
	"context"
	"errors"
	"rest-api/modules/restaurant/restaurantmodel"
)

type DeleteRestaurantStore interface {
	FindDataByCondition(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
	SoftDelete(ctx context.Context, id int) error
}

type deleteRestaurantBusiness struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantStore(store DeleteRestaurantStore) *deleteRestaurantBusiness {
	return &deleteRestaurantBusiness{store: store}
}

func (u *deleteRestaurantBusiness) DeleteRestaurant(ctx context.Context, id int) error {
	restaurant, err := u.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if restaurant.Status == 0 {
		return errors.New("data has been deleted")
	}

	if err := u.store.SoftDelete(ctx, id); err != nil {
		return err
	}

	return nil
}
