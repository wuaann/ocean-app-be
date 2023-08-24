package postmodel

import "ocean-app-be/common"

const EntityName = "Post"

type Post struct {
	common.PSModel `json:",inline"`
	Caption        string             `json:"caption" gorm:"column:caption"`
	Images         *common.Image      `json:"images,omitempty" gorm:"column:images;type:json"`
	LikedCount     int                `json:"liked_count" gorm:"column:liked_count"`
	User           *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (Post) TableName() string {
	return "posts"
}
func (data *Post) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypePost)
}

type PostCreate struct {
	common.PSModel `json:",inline"`
	UserID         int           `json:"user_id" gorm:"column:user_id"`
	Caption        string        `json:"caption" gorm:"column:caption"`
	Images         *common.Image `json:"images,omitempty" gorm:"column:images;type:json"`
	LikedCount     int           `json:"liked_count" gorm:"column:liked_count"`
}

func (PostCreate) TableName() string {
	return "posts"
}

func (data *PostCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypePost)
}
