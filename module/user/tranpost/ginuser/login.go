package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ocean-app-be/common"
	"ocean-app-be/component/appcontext"
	"ocean-app-be/component/hasher"
	"ocean-app-be/component/tokenprovider/jwt"
	userbiz "ocean-app-be/module/user/biz"
	usermodel "ocean-app-be/module/user/model"
	userstorage "ocean-app-be/module/user/storage"
)

func LoginHandler(appCtx appcontext.AppCtx) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := ctx.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := appCtx.GetMainDBConnection()

		tokenprovider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSqlStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewLoginBiz(store, tokenprovider, md5, 60*60*24*7)
		account, err := biz.Login(ctx.Request.Context(), &loginUserData)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
	return nil
}
