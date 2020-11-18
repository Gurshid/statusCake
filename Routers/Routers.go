package Routers

import (
	"../Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("urls", Controllers.ListUrl)
	r.POST("urls", Controllers.AddNewUrl)
	r.GET("urls/:id", Controllers.GetUrl)
	r.PUT("urls/:id", Controllers.UpdateUrl)
	r.DELETE("urls/:id", Controllers.DeleteUrl)

	return r
}
