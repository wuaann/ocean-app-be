package postlikegin

import (
	"net/http"
	"ocean-app-be/common"
	"ocean-app-be/component/appcontext"
	restaurantlikebiz "ocean-app-be/module/postLike/biz"
	postlikestorage "ocean-app-be/module/postLike/storage"

	"github.com/gin-gonic/gin"
)

// DELETE /v1/restaurants/:id/unlike

func UserUnlikeRestaurantHandler(appCtx appcontext.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := postlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		// descStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebiz.NewUserUnlikeRestaurantBiz(store, appCtx.GetPubsub())

		if err := biz.UnlikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
