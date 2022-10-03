package restaurantlikestorage

import (
	"context"
	"rest-api/common"

	restaurantlikemodel "rest-api/modules/restaurantlike/model"
)

func (s *sqlStore) GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	type sqlData struct {
		RestaurantID int `gorm:"column:restaurant_id"`
		LikeCount    int `gorm:"column:count"`
	}

	var listLike []sqlData

	if err := s.db.
		Table(restaurantlikemodel.Like{}.TableName()).
		Select("restaurant_id, count(user_id) as count").
		Where("restaurant_id in ?", ids).
		Group("restaurant_id").
		Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, v := range listLike {
		result[v.RestaurantID] = v.LikeCount
	}

	return result, nil
}
