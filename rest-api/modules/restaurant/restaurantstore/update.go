package restaurantstore

import (
	"context"
	"gorm.io/gorm"

	"rest-api/common"
	"rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdates) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Update("like_count", gorm.Expr("like_count + 1")).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Update("like_count", gorm.Expr("like_count - 1")).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
