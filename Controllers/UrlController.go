package Controllers

import (
	"../ApiHelpers"
	"../Models"
	"github.com/gin-gonic/gin"
)

var service = Models.New()

func ListUrl(c *gin.Context) {
	urlMaster, err := service.GetAllUrl()
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, urlMaster)
	}
}

func AddNewUrl(c *gin.Context) {
	var urlMaster Models.UrlMaster
	c.BindJSON(&urlMaster)
	err := service.AddNewUrl(&urlMaster)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, urlMaster)
	}

}

func GetUrl(c *gin.Context) {
	id := c.Params.ByName("id")
	var urlMaster Models.UrlMaster
	err := service.GetUrl(&urlMaster, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, urlMaster)
	}
}

func UpdateUrl(c *gin.Context) {
	var urlMaster Models.UrlMaster
	id := c.Params.ByName("id")
	err := service.GetUrl(&urlMaster, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	}
	c.BindJSON(&urlMaster)
	err = service.UpdateUrl(&urlMaster)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, urlMaster)
	}
}

func DeleteUrl(c *gin.Context) {
	var urlMaster Models.UrlMaster
	id := c.Params.ByName("id")
	err := service.DeleteUrl(&urlMaster, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, urlMaster)
	}
}
