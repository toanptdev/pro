package restaurantlikemodel

import (
	"rest-api/common"
	"time"
)

const EntityName = "UserLikeRestaurant"

type Like struct {
	RestaurantID int                `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserID       int                `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (l Like) TableName() string {
	return "restaurant_likes"
}
