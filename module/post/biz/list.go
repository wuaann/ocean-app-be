package postbiz

import (
	"context"
	"ocean-app-be/common"
	postmodel "ocean-app-be/module/post/model"
)

type ListPostRepo interface {
	ListPost(
		ctx context.Context,
		paging *common.Paging,
		moreKeys ...string,
	) ([]postmodel.Post, error)
}

type listPostBiz struct {
	repo ListPostRepo
}

func NewListPostBiz(repo ListPostRepo) *listPostBiz {
	return &listPostBiz{repo: repo}
}

func (biz *listPostBiz) ListPost(
	ctx context.Context,
	paging *common.Paging,
	moreKeys ...string,
) ([]postmodel.Post, error) {

	result, err := biz.repo.ListPost(ctx, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(postmodel.EntityName, err)
	}
	return result, nil
}
