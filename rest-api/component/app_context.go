package component

import (
	"gorm.io/gorm"
	"os"
	"rest-api/component/tokenprovider"
	"strconv"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetSecret() string
	NewTokenConfig() (*tokenprovider.TokenConfig, error)
}

type appCtx struct {
	db     *gorm.DB
	secret string
}

func NewAppContext(db *gorm.DB, secret string) *appCtx {
	return &appCtx{db: db, secret: secret}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetSecret() string {
	return ctx.secret
}

func (ctx *appCtx) NewTokenConfig() (*tokenprovider.TokenConfig, error) {
	atExp, err := strconv.Atoi(os.Getenv("accessTokenExpiry"))
	if err != nil {
		return nil, err
	}
	rtExp, err := strconv.Atoi(os.Getenv("refreshTokenExpiry"))
	if err != nil {
		return nil, err
	}

	return &tokenprovider.TokenConfig{
		AccessTokenExpiry:  atExp,
		RefreshTokenExpiry: rtExp,
	}, nil
}
