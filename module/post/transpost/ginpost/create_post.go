package ginpost

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ocean-app-be/common"
	"ocean-app-be/component/appcontext"
	postbiz "ocean-app-be/module/post/biz"
	postmodel "ocean-app-be/module/post/model"
	poststorage "ocean-app-be/module/post/storage"
)

func CreatePost(appCtx appcontext.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data postmodel.Post
		db := appCtx.GetMainDBConnection()

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := poststorage.NewSqlStore(db)

		biz := postbiz.NewCreatePostBiz(store)

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data.UserID = requester.GetUserId()

		if err := biz.Create(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(true)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeID.String()))

	}
}
