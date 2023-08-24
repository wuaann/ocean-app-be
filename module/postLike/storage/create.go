package postlikestorage

import (
	"context"
	"ocean-app-be/common"
	postlikemodel "ocean-app-be/module/postLike/model"
)

func (s *sqlStore) Create(ctx context.Context, data *postlikemodel.Like) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
