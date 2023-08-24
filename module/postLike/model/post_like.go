package postlikemodel

import (
	"ocean-app-be/common"
	"time"
)

const EntityName = "UserLikePost"

type Like struct {
	PostId    int                `json:"post_id" gorm:"column:post_id;"`
	UserId    int                `json:"user_id" gorm:"column:user_id;"`
	CreatedAt *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User      *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (Like) TableName() string {
	return "likes"
}

func (l *Like) GetRestaurantId() int {
	return l.PostId
}

func ErrCannotLikePost(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"Cannot like this post",
		"ErrCannotLikePost",
	)
}

func ErrCannotUnlikePost(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"Cannot unlike this post",
		"ErrCannotUnlikePost",
	)
}
