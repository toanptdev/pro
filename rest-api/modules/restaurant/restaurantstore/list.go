package restaurantstore

import (
	"context"
	"rest-api/common"
	"rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) ListDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	db := s.db

	var results []restaurantmodel.Restaurant

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions).Where("status in (1)")

	if filter != nil {
		if filter.CityID > 0 {
			db = db.Where("city_id = ?", filter.CityID)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Order("id desc").Find(&results).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return results, nil
}
