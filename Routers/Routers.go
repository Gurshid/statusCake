package Routers

import (
	"../Controllers"
	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()
	urlController := Controllers.New()
	r.GET("urls", urlController.ListUrl)
	r.POST("urls", urlController.AddNewUrl)
	r.GET("urls/:id", urlController.GetUrl)
	r.PUT("urls/:id", urlController.UpdateUrl)
	r.DELETE("urls/:id", urlController.DeleteUrl)

	return r
}
