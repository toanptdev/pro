package restaurantstore

import (
	"context"

	"rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdates) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}
