package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"piwriw_blog/models"
	"piwriw_blog/service"
)

func GetProfile(c *gin.Context) {
	id := c.Param("id")
	profile, err := service.GetProfile(id)
	if err != nil {
		zap.L().Error("service.GetMeProfile() is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, profile)
}
func UpdateProfileHandler(c *gin.Context) {
	id := c.Param("id")
	m := new(models.Profile)
	err := c.ShouldBindJSON(m)
	if err != nil {
		zap.L().Error("UpdateProfileHandler with valid params", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	err = service.UpdateProfile(id, m)
	if err != nil {
		zap.L().Error("service.UpdateProfile(id, m) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}
