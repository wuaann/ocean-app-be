package postlikestorage

import (
	"context"
	"ocean-app-be/common"
	postlikemodel "ocean-app-be/module/postLike/model"

	"gorm.io/gorm"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*postlikemodel.Like, error) {
	var result postlikemodel.Like

	if err := s.db.Where(conditions).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
