package postbiz

import (
	"context"
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
