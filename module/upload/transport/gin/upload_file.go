package ginupload

import (
	"github.com/gin-gonic/gin"
	"ocean-app-be/common"
	"ocean-app-be/component/appcontext"
	uploadbiz "ocean-app-be/module/upload/biz"
	uploadstorage "ocean-app-be/module/upload/storage"
)

func UploadHandler(appCtx appcontext.AppCtx) func(*gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()

		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "ocean-app")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)

		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		imgStore := uploadstorage.NewSQLStore(db)

		biz := uploadbiz.NewUploadBiz(appCtx.UploadProvider(), imgStore)

		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}
