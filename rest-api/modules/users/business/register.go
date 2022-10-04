package userbusiness

import (
	"context"
	"rest-api/common"
	usermodel "rest-api/modules/users/model"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBusiness(registerStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		registerStorage: registerStorage,
		hasher:          hasher,
	}
}

func (r *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, err := r.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil && err != common.ErrorRecordNotFound {
		return common.ErrDB(err)
	}
	if user != nil {
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = r.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"
	data.Status = 1

	if err := r.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err, usermodel.EntityName)
	}

	return nil
}
