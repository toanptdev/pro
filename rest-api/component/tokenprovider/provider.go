package tokenprovider

import (
	"errors"
	"rest-api/common"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

type Token struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	Expiry    int       `json:"expiry"`
}

type TokenPayload struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
}

type TokenConfig struct {
	AccessTokenExpiry  int
	RefreshTokenExpiry int
}

func (t *TokenConfig) GetAtExp() int {
	return t.AccessTokenExpiry
}

func (t *TokenConfig) GetRtExp() int {
	return t.RefreshTokenExpiry
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrTokenNotFound",
	)

	ErrEncodingToken = common.NewCustomError(
		errors.New("error encoding token"),
		"error encoding token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = common.NewCustomError(
		errors.New("invalid token"),
		"invalid token",
		"ErrInvalidToken",
	)
)
