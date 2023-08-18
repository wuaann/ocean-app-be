package storage

import (
	"context"
	"ocean-app-be/common"
)

func (store *sqlStore) List(
	ctx context.Context,
	ids []int,
	moreKey ...string,
) ([]common.Image, error) {

	db := store.db.Table(common.Image{}.TableName())
	var result []common.Image
	if err := db.Where("id in (?)", ids).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil

}
