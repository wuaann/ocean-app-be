package storage

import (
	"context"
	"ocean-app-be/common"
)

func (store *sqlStore) Create(ctx context.Context, data *common.Image) error {
	db := store.db.Table(common.Image{}.TableName())

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
