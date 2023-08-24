package ginpost

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ocean-app-be/common"
	"ocean-app-be/component/appcontext"
	postbiz "ocean-app-be/module/post/biz"
	postrepo "ocean-app-be/module/post/repo"
	postlikestorage "ocean-app-be/module/postLike/storage"

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
		likeStore := postlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		repo := postrepo.NewPostRepo(store, likeStore)
		biz := postbiz.NewListPostBiz(repo)

		result, err := biz.ListPost(c.Request.Context(), &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeID.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
