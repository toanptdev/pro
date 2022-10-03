package restaurantbusiness

import (
	"context"
	"rest-api/common"
	"rest-api/modules/restaurant/restaurantmodel"
)

type ListRestaurantLikeRepository interface {
	GetListRestaurantLike(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBusiness struct {
	repository ListRestaurantLikeRepository
}

func NewListRestaurantBusiness(repository ListRestaurantLikeRepository) *listRestaurantBusiness {
	return &listRestaurantBusiness{repository: repository}
}

func (l *listRestaurantBusiness) ListRestaurant(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := l.repository.GetListRestaurantLike(ctx, filter, paging)
	if err != nil {
		return nil, err
	}

	return result, nil
}
