package restaurantlikebiz

import (
	"context"
	"ocean-app-be/common"
	postlikemodel "ocean-app-be/module/postLike/model"
	"ocean-app-be/pubsub"
)

type UserLikePostStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postlikemodel.Like, error)

	Create(ctx context.Context, data *postlikemodel.Like) error
}

// type IncreaseLikeCountStore interface {
// 	IncreaseLikeCount(ctx context.Context, id int) error
// }

type userLikeRestaurantBiz struct {
	store UserLikePostStore
	// incStore IncreaseLikeCountStore
	pubsub pubsub.Pubsub
}

func NewUserLikeRestaurantBiz(
	store UserLikePostStore,
	// incStore IncreaseLikeCountStore,
	pubsub pubsub.Pubsub,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store: store,
		// incStore: incStore,
		pubsub: pubsub,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *postlikemodel.Like,
) error {
	likeExist, _ := biz.store.FindDataByCondition(ctx, map[string]interface{}{"user_id": data.UserId, "post_id": data.PostId})

	if likeExist != nil {
		return common.NewCustomError(nil, "user already like post", "AlreadyLikePost")
	}

	err := biz.store.Create(ctx, data)

	if err != nil {
		return postlikemodel.ErrCannotLikePost(err)
	}

	// side effect
	// go func() {
	// 	defer common.AppRecover()

	// 	job := asyncjob.NewJob(func(ctx context.Context) error {
	// 		return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	// 	})

	// 	_ = asyncjob.NewGroup(true, job).Run(ctx)
	// }()
	// New solution: use pubsub
	// Do not inject directly here, hard to unit test. Inject through struct instead
	biz.pubsub.Publish(ctx, common.TopicUserLikePost, pubsub.NewMessage(data))

	return nil
}
