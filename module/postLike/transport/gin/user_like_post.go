package postlikegin

import (
	"net/http"
	"ocean-app-be/common"
	"ocean-app-be/component/appcontext"
	restaurantlikebiz "ocean-app-be/module/postLike/biz"
	postlikemodel "ocean-app-be/module/postLike/model"
	postLikestorage "ocean-app-be/module/postLike/storage"

	"github.com/gin-gonic/gin"
)

// POST /v1/restaurants/:id/like

func UserLikeRestaurantHandler(appCtx appcontext.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := postlikemodel.Like{
			PostId: int(uid.GetLocalID()),
			UserId: requester.GetUserId(),
		}

		store := postLikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		// incStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebiz.NewUserLikeRestaurantBiz(store, appCtx.GetPubsub())

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
