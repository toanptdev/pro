package restaurantlikebussiness

import (
	"context"

	"rest-api/common"
	restaurantlikemodel "rest-api/modules/restaurantlike/model"
)

type ListUserLikeRestaurantStore interface {
	GetUserLikeRestaurant(
		ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type listUserLikeRestaurantBusiness struct {
	store ListUserLikeRestaurantStore
}

func NewListUserLikeRestaurantBusiness(store ListUserLikeRestaurantStore) *listUserLikeRestaurantBusiness {
	return &listUserLikeRestaurantBusiness{store: store}
}

func (business *listUserLikeRestaurantBusiness) ListUsers(
	ctx context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {
	users, err := business.store.GetUserLikeRestaurant(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(err, restaurantlikemodel.EntityName)
	}

	return users, nil
}
