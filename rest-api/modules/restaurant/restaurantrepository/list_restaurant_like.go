package restaurantrepository

import (
	"context"
	"log"

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

type LikeStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type restaurantLikeRepository struct {
	restaurantStore ListRestaurantStore
	likeStore       LikeStore
}

func NewRestaurantLikeRepository(restaurantStore ListRestaurantStore, likeStore LikeStore) *restaurantLikeRepository {
	return &restaurantLikeRepository{
		restaurantStore: restaurantStore,
		likeStore:       likeStore,
	}
}

func (r *restaurantLikeRepository) GetListRestaurantLike(ctx context.Context, conditions map[string]interface{}, filter *restaurantmodel.Filter, paging *common.Paging, moreKeys ...string) ([]restaurantmodel.Restaurant, error) {
	result, err := r.restaurantStore.ListDataByCondition(ctx, nil, filter, paging, "User")
	if err != nil {
		return nil, err
	}

	ids := make([]int, len(result))

	for k, v := range result {
		ids[k] = v.ID
	}

	mapResLike, err := r.likeStore.GetRestaurantLikes(ctx, ids)

	if err != nil {
		log.Println("cannot get restaurant like: ", err)
	}

	if mapResLike != nil {
		for k, v := range result {
			result[k].LikedCount = mapResLike[v.ID]
		}
	}

	return result, nil
}
