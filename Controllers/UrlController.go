package Controllers

import (
	"../ApiHelpers"
	"../Models"
	"github.com/gin-gonic/gin"
)
type UrlController struct {
	urlService Models.UrlService
}

func New() UrlController {
	return UrlController{urlService: Models.New()}
}

func (u *UrlController) ListUrl(c *gin.Context) {
	urlMaster, err := u.urlService.GetAllUrl()
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, urlMaster)
	}
}

func (u *UrlController) AddNewUrl(c *gin.Context) {
	var urlMaster Models.UrlMaster
	c.BindJSON(&urlMaster)
	err := u.urlService.AddNewUrl(&urlMaster)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, urlMaster)
	}

}

func (u *UrlController) GetUrl(c *gin.Context) {
	id := c.Params.ByName("id")
	var urlMaster Models.UrlMaster
	err := u.urlService.GetAndCheckUrl(&urlMaster, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, urlMaster)
	}
}

func (u *UrlController) UpdateUrl(c *gin.Context) {
	var urlMaster Models.UrlMaster
	id := c.Params.ByName("id")
	err := u.urlService.GetUrl(&urlMaster, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	}
	c.BindJSON(&urlMaster)
	err = u.urlService.UpdateUrl(&urlMaster, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, urlMaster)
	}
}

func (u *UrlController) DeleteUrl(c *gin.Context) {
	var urlMaster Models.UrlMaster
	id := c.Params.ByName("id")
	err := u.urlService.DeleteUrl(&urlMaster, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, urlMaster)
	}
}
