package userstorage

import (
	"context"
	"gorm.io/gorm"
	"ocean-app-be/common"
	usermodel "ocean-app-be/module/user/model"
)

func (s *sqlStore) Find(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &user, nil
}
