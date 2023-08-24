package middleware

import (
	"github.com/gin-gonic/gin"
	"ocean-app-be/common"
	"ocean-app-be/component/appcontext"
)

func Recover(ac appcontext.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err) // gin has it's owned recovery
					return
				}

				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err) // gin has it's owned recovery
				return
			}
		}()

		c.Next()
	}
}
