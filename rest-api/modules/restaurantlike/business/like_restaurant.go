package restaurantlikebussiness

import (
	"context"
	"errors"
	"rest-api/common"
	"rest-api/component/asyncjob"

	restaurantlikemodel "rest-api/modules/restaurantlike/model"
)

type LikeRestaurantStore interface {
	Find(ctx context.Context, userID int, restaurantID int) (*restaurantlikemodel.Like, error)
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type IncreaseRestaurantLikeStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type likeRestaurantBusiness struct {
	store     LikeRestaurantStore
	incrStore IncreaseRestaurantLikeStore
}

func NewLikeRestaurantBusiness(store LikeRestaurantStore, incrStore IncreaseRestaurantLikeStore) *likeRestaurantBusiness {
	return &likeRestaurantBusiness{store: store, incrStore: incrStore}
}

func (l *likeRestaurantBusiness) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	like, err := l.store.Find(ctx, data.UserID, data.RestaurantID)
	if err != nil {
		if err != common.ErrorRecordNotFound {
			return common.ErrCantGetEntity(restaurantlikemodel.EntityName, err)
		}
	}

	if like != nil {
		return common.ErrEntityExisted(errors.New("user already like restaurant"), restaurantlikemodel.EntityName)
	}

	if err := l.store.Create(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	// side effect
	job := asyncjob.NewJob(func(ctx context.Context) error {
		return l.incrStore.IncreaseLikeCount(ctx, data.RestaurantID)
	})
	_ = asyncjob.NewGroup(true, job).Run(ctx)

	return nil
}
