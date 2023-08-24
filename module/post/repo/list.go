package postrepo

import (
	"context"
	"log"
	"ocean-app-be/common"
	postmodel "ocean-app-be/module/post/model"
)

type ListPostStore interface {
	ListPostByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		paging *common.Paging,
		moreKeys ...string,
	) ([]postmodel.Post, error)
}

type LikeStore interface {
	GetPostLikes(
		ctx context.Context,
		ids []int,
	) (map[int]int, error)
}

type listPostStore struct {
	store     ListPostStore
	likeStore LikeStore
}

func NewPostRepo(
	store ListPostStore,
	likeStore LikeStore,
) *listPostStore {
	return &listPostStore{store: store, likeStore: likeStore}
}

func (biz *listPostStore) ListPostByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	paging *common.Paging,
	moreKeys ...string,
) ([]postmodel.Post, error) {
	result, err := biz.store.ListPostByCondition(ctx, conditions, paging, "User")

	if err != nil {
		return nil, common.ErrCannotListEntity(postmodel.EntityName, err)
	}

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	mapResLike, err := biz.likeStore.GetPostLikes(ctx, ids)

	if err != nil {
		log.Println("cannot get restaurant likes:", err)
	}

	if v := mapResLike; v != nil {
		for i, item := range result {
			result[i].LikedCount = mapResLike[item.Id]
		}
	}

	return result, nil
}
