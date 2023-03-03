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

func UserLoginHandler(c *gin.Context) {
	m := new(models.ParamUser)
	if err := c.ShouldBindJSON(m); err != nil {
		zap.L().Error("UserLoginHandler with valid params", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	respUser, err := service.Login(m)
	if err != nil {
		zap.L().Error("service.Login(m) is failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		} else if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, respUser)
}
func RegisterUserHandler(c *gin.Context) {
	m := new(models.ParamUser)
	err := c.ShouldBindJSON(m)
	if err != nil {
		zap.L().Error("AddUserHandler with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	err = service.Register(m)
	if err != nil {
		zap.L().Error("service.Register(m) is failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		} else if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}
func GetAllUsersHandler(c *gin.Context) {
	pagenum, pagesize := getPageInfo(c)
	users, err := service.GetAllUsers(pagenum, pagesize)
	if err != nil {
		zap.L().Error("service.GetAllUsers(pagenum, pagesize) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, users)
}

func DeleteUserByIdHandler(c *gin.Context) {
	id := c.Param("id")
	err := service.DeleteUserById(id)
	if err != nil {
		zap.L().Error("service.DeleteUserById(id) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseError(c, CodeSuccess)
}
func GetUserByIdHandler(c *gin.Context) {
	id := c.Param("id")
	user, err := service.GetUserById(id)
	if err != nil {
		zap.L().Error("service.GetUserById(id) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, user)
}
func UpdateUserByIdHandler(c *gin.Context) {
	id := c.Param("id")
	m := new(models.ParamEditUser)
	if err := c.ShouldBindJSON(m); err != nil {
		zap.L().Error("UpdateUserHandler with valid params", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	err := service.UpdateUserById(id, m)
	if err != nil {
		zap.L().Error("service.UpdateUserById(id, m) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}
func UpdateUserPasswordHandler(c *gin.Context) {
	password := c.Query("password")
	id := c.Param("id")
	err := service.UpdateUserPassword(password, id)
	if err != nil {
		zap.L().Error("service.UpdateUserPassword(password, id) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}
func GetUserListHandler(c *gin.Context) {
	pagenum, pagesize := getPageInfo(c)
	username := c.Query("username")
	resppage, err := service.GetUserList(pagenum, pagesize, username)
	if err != nil {
		zap.L().Error("service.GetUserList(pagenum, pagesize, username) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, resppage)
}
