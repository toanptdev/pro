package restaurantstore

import (
	"context"
	"gorm.io/gorm"
	"rest-api/common"
	"rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) FindDataByCondition(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	db := s.db
	var restaurant *restaurantmodel.Restaurant

	for v := range moreKeys {
		db = db.Preload(moreKeys[v])
	}

	if err := db.Where(conditions).First(&restaurant).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrorRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return restaurant, nil
}
