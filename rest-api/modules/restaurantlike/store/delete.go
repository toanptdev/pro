package restaurantlikestorage

import (
	"context"
	"rest-api/common"
	restaurantlikemodel "rest-api/modules/restaurantlike/model"
)

func (s *sqlStore) Delete(ctx context.Context, userID int, restaurantID int) error {
	db := s.db

	if err := db.Table(restaurantlikemodel.Like{}.TableName()).Where("user_id = ? and restaurant_id = ?", userID, restaurantID).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
