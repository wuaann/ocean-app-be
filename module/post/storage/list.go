package poststorage

import (
	"context"
	"ocean-app-be/common"
	postmodel "ocean-app-be/module/post/model"
)

func (store *sqlStore) ListPostByCondition(
	ctx context.Context,
	conditions map[string]interface{},

	paging *common.Paging,
	moreKeys ...string,
) ([]postmodel.Post, error) {

	var result []postmodel.Post

	db := store.db.Table(postmodel.Post{}.TableName()).Where(conditions)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			db = db.Where("id< ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil

}
