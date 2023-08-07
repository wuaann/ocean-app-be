package userstorage

import (
	"ocean-app-be/common"
	"ocean-app-be/component/appcontext"
	usermodel "ocean-app-be/module/user/model"
)

func (s *sqlStore) Create(ctx appcontext.AppCtx, data *usermodel.UserCreate) error {
	db := s.db.Begin().Table(data.TableName())
	if err := db.Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	return nil
}
