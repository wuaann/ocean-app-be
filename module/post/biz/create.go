package postbiz

import (
	"context"
	"ocean-app-be/common"
	postmodel "ocean-app-be/module/post/model"
)

type CreatePostStore interface {
	Create(ctx context.Context, data *postmodel.Post) error
}

type createPostBiz struct {
	store CreatePostStore
}

func NewCreatePostBiz(store CreatePostStore) *createPostBiz {
	return &createPostBiz{store: store}
}

func (biz *createPostBiz) Create(ctx context.Context, data *postmodel.Post) error {
	data.Status = 1
	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrDB(err)
	}
	return nil
}
