package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"piwriw_blog/dao/mysql"
	"piwriw_blog/models"
	"piwriw_blog/service"
)

// AdminLoginHandler : 后台管理登录
func AdminLoginHandler(c *gin.Context) {
	m := new(models.ParamLogin)
	if err := c.ShouldBindJSON(m); err != nil {
		zap.L().Error("AdminLoginHandler with valid params", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	user, err := service.AdminLogin(m)
	if err != nil {
		zap.L().Error("service.AdminLogin(m) is failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		} else if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, user)
}

func UploadImgHandler(c *gin.Context) {
	//file, err := c.FormFile("file")
	//if err != nil {
	//	zap.L().Error("UploadImgHandler with valid params", zap.Error(err))
	//	ResponseError(c, CodeInvalidParam)
	//	return
	//}
	//service.Add
	//ResponseSuccess(c, CodeSuccess)
}
