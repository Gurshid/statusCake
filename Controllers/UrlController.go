package Controllers

import (
	"../ApiHelpers"
	"../Models"
	"github.com/gin-gonic/gin"
)

func ListUrl(c *gin.Context) {
	var urlMaster []Models.UrlMaster
	err := Models.GetAllUrl(&urlMaster)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, urlMaster)
	}
}

func AddNewUrl(c *gin.Context) {
	var urlMaster Models.UrlMaster
	c.BindJSON(&urlMaster)
	err := Models.AddNewUrl(&urlMaster)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, urlMaster)
	}

}

func GetUrl(c *gin.Context) {
	id := c.Params.ByName("id")
	var urlMaster Models.UrlMaster
	err := Models.GetUrl(&urlMaster, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, urlMaster)
	}
}

func UpdateUrl(c *gin.Context) {
	var urlMaster Models.UrlMaster
	id := c.Params.ByName("id")
	err := Models.GetUrl(&urlMaster, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	}
	c.BindJSON(&urlMaster)
	err = Models.UpdateUrl(&urlMaster, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, urlMaster)
	}
}

func DeleteUrl(c *gin.Context) {
	var urlMaster Models.UrlMaster
	id := c.Params.ByName("id")
	err := Models.DeleteUrl(&urlMaster, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, urlMaster)
	}
}
