package ginpost

import (
	"github.com/gin-gonic/gin"
	"ocean-app-be/common"
	"ocean-app-be/component/appcontext"

	poststorage "ocean-app-be/module/post/storage"
)

func ListPost(appCtx appcontext.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		store := poststorage.NewSqlStore(appCtx.GetMainDBConnection())
		//likeStore := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		//repo := postrepo.NewPostRepo(store, )
		//biz := postbiz.NewListPostBiz(repo)
	}
}
