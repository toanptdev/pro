package restaurantlikestorage

import (
	"context"
	"gorm.io/gorm"
	"rest-api/common"
	restaurantlikemodel "rest-api/modules/restaurantlike/model"
)

func (s *sqlStore) Find(ctx context.Context, userID int, restaurantID int) (*restaurantlikemodel.Like, error) {
	var restaurantLike *restaurantlikemodel.Like

	db := s.db

	if err := db.Where("user_id = ? and restaurant_id = ?", userID, restaurantID).First(&restaurantLike).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrorRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return restaurantLike, nil
}
