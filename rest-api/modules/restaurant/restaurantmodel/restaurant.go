package restaurantmodel

import (
	"errors"
	"strings"

	"rest-api/common"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name"`
	Addr            string `json:"addr" gorm:"column:addr"`
}

func (r Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdates struct {
	Name *string `json:"name" gorm:"column:name"`
	Addr *string `json:"addr" gorm:"column:addr"`
}

func (r RestaurantUpdates) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	Id   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Addr string `json:"addr" gorm:"column:addr"`
}

func (r RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (r *RestaurantCreate) Validate() error {
	r.Name = strings.TrimSpace(r.Name)
	if len(r.Name) == 0 {
		return errors.New("restaurant name cant be empty")
	}

	return nil
}
