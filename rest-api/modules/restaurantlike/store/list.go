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

func (s *sqlStore) GetUserLikeRestaurant(
	ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	var results []restaurantlikemodel.Like

	db := s.db

	db = db.Table(restaurantlikemodel.Like{}.TableName()).Where(conditions)

	if filter != nil {
		if filter.RestaurantID > 0 {
			db = db.Where("restaurant_id = ?", filter.RestaurantID)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	//for i := range moreKeys {
	//	db = db.Preload(moreKeys[i])
	//}

	db = db.Preload("User")

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&results).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(results))

	for k, v := range results {
		users[k] = *v.User
	}
	return users, nil
}
