package poststorage

import (
	"context"
	"ocean-app-be/common"
	postmodel "ocean-app-be/module/post/model"
)

func (s *sqlStore) Create(ctx context.Context, data *postmodel.PostCreate) error {
	db := s.db.Table(postmodel.Post{}.TableName())

	if err := db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
