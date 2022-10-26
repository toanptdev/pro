package restaurantlikebussiness

import (
	"context"
	"rest-api/common"
	"rest-api/component/asyncjob"
	restaurantlikemodel "rest-api/modules/restaurantlike/model"
)

type UnlikeRestaurantStore interface {
	Find(ctx context.Context, userID int, restaurantID int) (*restaurantlikemodel.Like, error)
	Delete(ctx context.Context, userID int, restaurantID int) error
}

type DecreaseRestaurantLikeStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type unLikeRestaurantBusiness struct {
	store    UnlikeRestaurantStore
	decrLike DecreaseRestaurantLikeStore
}

func NewLUnLikeRestaurantBusiness(store UnlikeRestaurantStore, decrLike DecreaseRestaurantLikeStore) *unLikeRestaurantBusiness {
	return &unLikeRestaurantBusiness{store: store, decrLike: decrLike}
}

func (l *unLikeRestaurantBusiness) UnLikeRestaurant(ctx context.Context, userID int, restaurantID int) error {
	_, err := l.store.Find(ctx, userID, restaurantID)
	if err != nil {
		if err == common.ErrorRecordNotFound {
			return restaurantlikemodel.ErrUserHaveNotLikeRestaurant(err)
		}
		return common.ErrCantGetEntity(restaurantlikemodel.EntityName, err)
	}

	if err := l.store.Delete(ctx, userID, restaurantID); err != nil {
		return restaurantlikemodel.ErrCannotUnLikeRestaurant(err)
	}

	// side effect
	go func() {
		defer common.AppRecover()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return l.decrLike.DecreaseLikeCount(ctx, restaurantID)
		})
		_ = asyncjob.NewGroup(true, job).Run(ctx)
	}()

	return nil
}
