package usermodel

import (
	"errors"
	"rest-api/common"
	"rest-api/component/tokenprovider"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email"`
	Password        string `json:"-" gorm:"column:password"`
	LastName        string `json:"last_name" gorm:"column:last_name"`
	FirstName       string `json:"first_name" gorm:"column:first_name"`
	Role            string `json:"role" gorm:"column:role;"`
	Salt            string `json:"-" gorm:"column:salt;"`
}

func (u *User) GetUserID() int {
	return u.ID
}

func (u *User) GetRole() string {
	return u.Role
}

func (u *User) GetEmail() string {
	return u.Email
}

func (User) TableName() string {
	return "users"
}

func (u *User) Mask() {
	u.GenerateUID(common.DBTypeUser)
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email"`
	Password        string `json:"password" gorm:"column:password"`
	LastName        string `json:"last_name" gorm:"column:last_name"`
	FirstName       string `json:"first_name" gorm:"column:first_name"`
	Role            string `json:"-" gorm:"column:role;"`
	Salt            string `json:"-" gorm:"column:salt;"`
}

func (u UserCreate) TableName() string {
	return User{}.TableName()
}

func (u *UserCreate) Mask() {
	u.GenerateUID(common.DBTypeUser)
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email"`
	Password string `json:"password" form:"password" gorm:"column:password"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

type Account struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token"`
}

func NewAccount(at, rt *tokenprovider.Token) *Account {
	return &Account{
		AccessToken:  at,
		RefreshToken: rt,
	}
}

var (
	ErrUserNameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUserNameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email already exist"),
		"email already existed",
		"ErrEmailExist",
	)
)
