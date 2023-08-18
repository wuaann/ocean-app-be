package userbiz

import (
	"context"
	"ocean-app-be/common"
	"ocean-app-be/component/appcontext"
	"ocean-app-be/component/tokenprovider"
	usermodel "ocean-app-be/module/user/model"
)

type LoginStore interface {
	Find(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBiz struct {
	appCtx        appcontext.AppCtx
	userStore     LoginStore
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBiz(userStore LoginStore, provider tokenprovider.Provider, hasher Hasher, expiry int) *loginBiz {
	return &loginBiz{
		userStore:     userStore,
		tokenProvider: provider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	user, err := biz.userStore.Find(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	hashedPassword := biz.hasher.Hash(data.SaltedPassword + user.Salt)

	if user.SaltedPassword != hashedPassword {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}
	payload := tokenprovider.TokenPayload{
		UserId: user.UserId,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)

	}
	account := usermodel.NewAccount(accessToken, refreshToken)

	return account, nil
}
