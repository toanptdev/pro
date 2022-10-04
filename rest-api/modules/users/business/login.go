package userbusiness

import (
	"context"
	"rest-api/common"
	"rest-api/component/tokenprovider"

	usermodel "rest-api/modules/users/model"
)

type LoginStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
}

type TokenConfig interface {
	GetAtExp() int
	GetRtExp() int
}

type loginBusiness struct {
	loginStore    LoginStore
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	tkCfg         TokenConfig
}

func NewLoginBusiness(
	loginStore LoginStore,
	tokenProvider tokenprovider.Provider,
	hasher Hasher,
	tkCfg TokenConfig,
) *loginBusiness {
	return &loginBusiness{
		loginStore:    loginStore,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		tkCfg:         tkCfg,
	}
}

// Login Flow

// 1. Find User By Email
// 2. Hash password from request with salt in db then compare
// 3. Issue Token for client
// 4. Return token

func (l *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	user, err := l.loginStore.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErrUserNameOrPasswordInvalid
	}

	passHashed := l.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrUserNameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserID: user.ID,
		Role:   user.Role,
	}

	accessToken, err := l.tokenProvider.Generate(payload, l.tkCfg.GetAtExp())
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := l.tokenProvider.Generate(payload, l.tkCfg.GetRtExp())
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(accessToken, refreshToken)

	return account, nil
}
