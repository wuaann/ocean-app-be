package restaurantlikebiz

import (
	"context"
	"ocean-app-be/common"
	postlikemodel "ocean-app-be/module/postLike/model"
	"ocean-app-be/pubsub"
)

type UserUnlikeRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postlikemodel.Like, error)
	Delete(ctx context.Context, userId, restaurantId int) error
}

// type DecreaseLikeCountStore interface {
// 	DecreaseLikeCount(ctx context.Context, id int) error
// }

type userUnlikeRestaurantBiz struct {
	store UserUnlikeRestaurantStore
	// descStore DecreaseLikeCountStore
	pubsub pubsub.Pubsub
}

func NewUserUnlikeRestaurantBiz(
	store UserUnlikeRestaurantStore,
	// descStore DecreaseLikeCountStore,
	pubsub pubsub.Pubsub,
) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{
		store: store,
		// descStore: descStore,
		pubsub: pubsub,
	}
}

func (biz *userUnlikeRestaurantBiz) UnlikeRestaurant(
	ctx context.Context,
	userId,
	restaurantId int,
) error {
	_, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"user_id": userId, "restaurant_id": restaurantId})

	if err != nil {
		return common.NewCustomError(nil, "user has not like restaurant", "NotLikeRestaurant")
	}

	if err := biz.store.Delete(ctx, userId, restaurantId); err != nil {
		return postlikemodel.ErrCannotUnlikePost(err)
	}

	// side effect
	// go func() {
	// 	defer common.AppRecover()

	// 	job := asyncjob.NewJob(func(ctx context.Context) error {
	// 		return biz.descStore.DecreaseLikeCount(ctx, restaurantId)
	// 	})

	// 	_ = asyncjob.NewGroup(true, job).Run(ctx)
	// }()
	// New solution: use pubsub
	// Do not inject directly here, hard to unit test. Inject through struct instead
	biz.pubsub.Publish(ctx, common.TopicUserUnlikePost, pubsub.NewMessage(&postlikemodel.Like{
		PostId: restaurantId,
		UserId: userId,
	}))

	return nil
}
