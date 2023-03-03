package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"piwriw_blog/models"
	"piwriw_blog/service"
)

func GetAllCommentList(c *gin.Context) {
	aid := c.Param("id")
	pagenum, pagesize := getPageInfo(c)
	respcomment, err := service.GetAllCommentList(pagenum, pagesize, aid)
	if err != nil {
		zap.L().Error("service.GetAllComments(pagenum, pagesize) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, respcomment)
}
func GetAllComments(c *gin.Context) {
	pagenum, pagesize := getPageInfo(c)
	respcomment, err := service.GetAllComments(pagenum, pagesize)
	if err != nil {
		zap.L().Error("service.GetAllComments(pagenum, pagesize) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, respcomment)
}

func CheckCommentHandler(c *gin.Context) {
	id := c.Param("id")
	status := c.Param("status")
	err := service.CheckComment(id, status)
	if err != nil {
		zap.L().Error("service.CheckComment(id, status) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}
func GetCommentStatusById(c *gin.Context) {
	id := c.Param("id")
	comment, err := service.GetCommentById(id)
	if err != nil {
		zap.L().Error("service.GetCommentById(id) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
	}
	ResponseSuccess(c, comment.Status)
}
func AddCommentHandler(c *gin.Context) {
	m := new(models.ParamComment)
	err := c.ShouldBindJSON(m)
	if err != nil {
		zap.L().Error("AddCommentHandler with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 从content 获取当前发送请求用户ID
	_, err = getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	err = service.AddComment(m)
	if err != nil {
		zap.L().Error("service.AddComment(m) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}
