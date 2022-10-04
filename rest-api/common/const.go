package common

const (
	DBTypeRestaurant = 1
	DBTypeUser       = 2
)

type Requester interface {
	GetRole() string
	GetEmail() string
	GetUserID() int
}
