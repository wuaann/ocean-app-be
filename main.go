package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"ocean-app-be/component/appcontext"
	middleware "ocean-app-be/midleware"
	"ocean-app-be/module/user/tranpost/ginuser"
	"os"
)

func main() {
	dsn := os.Getenv("POSTGRES_CONN_STRING")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.Debug()

	appCtx := appcontext.NewAppCtx(db)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("v1")

	v1.POST("register", ginuser.RegisterHandler(appCtx))

	r.Run()
}
