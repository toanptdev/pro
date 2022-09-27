package restaurantmodel

type Filter struct {
	CityID int `json:"city_id,omitempty" form:"city_id"`
}
