package router

import (
	"ZShortUrl/controller"
	"ZShortUrl/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	user := controller.GetNewUser()
	r.GET("/:path", user.Go)
	r.POST("/genNewUrl", user.GenNewUrl)
	return r
}
