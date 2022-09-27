package restaurantbusiness

import (
	"context"

	"rest-api/common"
	"rest-api/modules/restaurant/restaurantmodel"
)

type ListRestaurantStore interface {
	ListDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBusiness struct {
	store ListRestaurantStore
}

func NewListRestaurantBusiness(store ListRestaurantStore) *listRestaurantBusiness {
	return &listRestaurantBusiness{store: store}
}

func (l *listRestaurantBusiness) ListRestaurant(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := l.store.ListDataByCondition(ctx, nil, filter, paging)
	if err != nil {
		return nil, err
	}

	return result, nil
}
