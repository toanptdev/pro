package userstorage

import (
	"context"
	"gorm.io/gorm"
	"rest-api/common"
	usermodel "rest-api/modules/users/model"
)

func (s *sqlStore) FindUser(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error) {
	db := s.db.Begin()

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var user *usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrorRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return user, nil
}
