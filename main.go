package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"ocean-app-be/component/appcontext"
	"ocean-app-be/component/uploadprovider"
	middleware "ocean-app-be/midlewares"
	"ocean-app-be/module/post/transpost/ginpost"
	ginupload "ocean-app-be/module/upload/transport/gin"
	"ocean-app-be/module/user/tranpost/ginuser"
	"os"
)

func main() {

	dsn := os.Getenv("POSTGRES_CONN_STRING")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	secretKey, ok := os.LookupEnv("SECRET_KEY")

	if !ok {
		log.Fatalln("Missing Secret Key string.")
	}

	db.Debug()

	bucketName, ok := os.LookupEnv("BUCKET_NAME")
	if !ok {
		log.Fatalln("Missing firebase Region string.")
	}

	fibKey, ok := os.LookupEnv("FIREBASE_KEY")
	if !ok {
		log.Fatalln("Missing firebase API Key string.")
	}

	Cloudinary := uploadprovider.NewFirebaseProvider(bucketName, fibKey)

	appCtx := appcontext.NewAppCtx(db, secretKey, Cloudinary)

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://127.0.0.1:3000")
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("v1")
	{
		v1.POST("register", ginuser.RegisterHandler(appCtx))
		v1.POST("login", ginuser.LoginHandler(appCtx))

		v1.POST("upload", ginupload.UploadHandler(appCtx))

		post := v1.Group("post", middleware.RequiredAuth(appCtx))
		{
			post.POST("", ginpost.CreatePost(appCtx))
			post.GET("", ginpost.ListPost(appCtx))

		}

	}
	r.Run()
}
