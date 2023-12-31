package userbiz

import (
	"ocean-app-be/common"
	"ocean-app-be/component/appcontext"
	usermodel "ocean-app-be/module/user/model"
)

type RegisterStore interface {
	Find(ctx appcontext.AppCtx, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	Create(ctx appcontext.AppCtx, data *usermodel.UserCreate) error
}
type Hasher interface {
	Hash(data string) string
}
type registerBiz struct {
	registerStore RegisterStore
	hasher        Hasher
}

func NewRegisterBiz(registerStore RegisterStore, hasher Hasher) *registerBiz {
	return &registerBiz{
		registerStore: registerStore,
		hasher:        hasher,
	}
}
func (biz *registerBiz) Register(ctx appcontext.AppCtx, data *usermodel.UserCreate) error {
	user, _ := biz.registerStore.Find(ctx, map[string]interface{}{"username": data.Username})
	if user != nil {
		return usermodel.ErrEmailExisted
	}
	salt := common.GenSalt(50)
	data.SaltedPassword = biz.hasher.Hash(data.SaltedPassword + salt)
	data.Salt = salt
	data.Role = "normal"
	data.Status = 1
	if err := biz.registerStore.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil

}
