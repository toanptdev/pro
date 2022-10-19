package restaurantlikemodel

type Filter struct {
	RestaurantID int `json:"restaurant_id" form:"restaurant_id"`
	UserID       int `json:"user_id" form:"user_id"`
}
