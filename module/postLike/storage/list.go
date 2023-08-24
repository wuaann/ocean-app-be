package postlikestorage

import (
	"context"
	"fmt"
	"github.com/btcsuite/btcd/btcutil/base58"
	"ocean-app-be/common"
	postlikemodel "ocean-app-be/module/postLike/model"
	"time"
)

const timeLayout = "2006-01-02T15:04:05.999999"

type sqlData struct {
	PostID    int `gorm:"column:post_id;"`
	LikeCount int `gorm:"column:count;"`
}

func (s *sqlStore) GetUsersLikePost(
	ctx context.Context,
	conditions map[string]interface{},
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	var result []postlikemodel.Like

	db := s.db

	db = db.Table(postlikemodel.Like{}.TableName()).Where(conditions)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	// for i := range moreKeys {
	// 	db = db.Preload(moreKeys[i])
	// }

	db = db.Preload("User")

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("created_at < ?", timeCreated)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(result))

	for i, item := range result {
		result[i].User.CreatedAt = item.CreatedAt
		result[i].User.UpdatedAt = nil
		users[i] = *result[i].User

		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))))
			paging.NextCursor = cursorStr
		}
	}

	return users, nil
}

func (s *sqlStore) GetPostLikes(
	ctx context.Context,
	ids []int,
) (map[int]int, error) {
	result := make(map[int]int)

	var listLike []sqlData

	if err := s.db.Table(postlikemodel.Like{}.TableName()).Select("post_id, count(post_id) as count").
		Where("post_id in (?)", ids).
		Group("post_id").
		Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.PostID] = item.LikeCount
	}

	return result, nil
}
