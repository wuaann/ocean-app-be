package postlikestorage

import (
	"ocean-app-be/common"
	postlikemodel "ocean-app-be/module/postLike/model"
)

import (
	"context"
)

func (s *sqlStore) Delete(ctx context.Context, userId, postId int) error {
	db := s.db.Table(postlikemodel.Like{}.TableName())

	if err := db.
		Where("user_id = ? and post_id = ?", userId, postId).
		Delete(nil).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
