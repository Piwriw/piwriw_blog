package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"piwriw_blog/models"
	"piwriw_blog/service"
)

func GetCategoryListHandler(c *gin.Context) {
	pagenum, pagesize := getPageInfo(c)
	categoryList, err := service.GetCategoryList(pagenum, pagesize)
	if err != nil {
		zap.L().Error("service.GetCategoryList(pagenum, pagesize) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
	}
	ResponseSuccess(c, categoryList)
}
func GetCategoriesHandler(c *gin.Context) {
	category, err := service.GetAllCategory()
	if err != nil {
		zap.L().Error("service.GetAllCategory() is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, category)
}
func AddCategoryHandler(c *gin.Context) {
	m := new(models.ParamCategory)
	err := c.ShouldBindJSON(m)
	if err != nil {
		zap.L().Error("AddCategoryHandler with valid params", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	err = service.AddCategory(m)
	if err != nil {
		zap.L().Error("service.AddCategory(m) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}
func UpdateCategoryHandler(c *gin.Context) {
	m := new(models.ParamCategory)
	err := c.ShouldBindJSON(m)
	if err != nil {
		zap.L().Error("UpdateCategoryHandler with valid params", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	err = service.UpdateCategory(m)
	if err != nil {
		zap.L().Error("service.UpdateCategory(m) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}
func GetCategoryById(c *gin.Context) {
	id := c.Param("id")
	err, category := service.GetCategoryById(id)
	if err != nil {
		zap.L().Error("service.GetCategoryById(id) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, category)
}
func DeleteCategoryById(c *gin.Context) {
	id := c.Param("id")
	err := service.DeleteCategoryById(id)
	if err != nil {
		zap.L().Error("service.DeleteCategoryById(id) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}
