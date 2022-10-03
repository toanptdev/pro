package restaurantlikemodel

import (
	"time"
)

type Like struct {
	RestaurantID int        `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserID       int        `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at;"`
}

func (l Like) TableName() string {
	return "restaurant_likes"
}
