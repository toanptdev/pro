package restaurantmodel

import (
	"errors"
	"strings"

	"rest-api/common"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name"`
	Addr            string             `json:"addr" gorm:"column:addr"`
	LikedCount      int                `json:"liked_count" gorm:"column:like_count"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false;foreignKey:ID;references:ID"`
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
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name"`
	Addr            string `json:"addr" gorm:"column:addr"`
	OwnerID         int    `json:"-" gorm:"column:owner_id"`
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

func (r *Restaurant) Mask(isAdminOwner bool) {
	r.GenerateUID(common.DBTypeRestaurant)

	if r.User != nil {
		r.User.Mask()
	}
}
