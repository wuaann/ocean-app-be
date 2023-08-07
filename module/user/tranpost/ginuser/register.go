package ginuser

import (
	"github.com/gin-gonic/gin"
	"ocean-app-be/common"
	"ocean-app-be/component/appcontext"
	"ocean-app-be/component/hasher"
	userbiz "ocean-app-be/module/user/biz"
	usermodel "ocean-app-be/module/user/model"
	userstorage "ocean-app-be/module/user/storage"
)

func RegisterHandler(appCtx appcontext.AppCtx) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data usermodel.UserCreate
		db := appCtx.GetMainDBConnection()
		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := userstorage.NewSqlStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBiz(store, md5)
		if err := biz.Register(ctx.Request.Context(), &data); err != nil {

		}
	}
}
