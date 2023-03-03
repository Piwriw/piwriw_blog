package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"piwriw_blog/models"
	"piwriw_blog/service"
)

func AddArticleHandler(c *gin.Context) {
	m := new(models.ParamArticle)
	if err := c.ShouldBindJSON(m); err != nil {
		zap.L().Error("AddArticle with valid params", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	err := service.AddArticle(m)
	if err != nil {
		zap.L().Error(" service.AddArticle(m) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}
func GetAllArticleHandler(c *gin.Context) {
	page, limit := getPageInfo(c)
	respArt, err := service.GetAllArticles(page, limit)
	if err != nil {
		zap.L().Error("service.GetAllArticles()", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, respArt)
}

func GetArticleById(c *gin.Context) {
	articleId := c.Param("id")
	article, err := service.GetArticleById(articleId)
	if err != nil {
		zap.L().Error("service.AddArticleById(articleId) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, article)
}
func GetArticleListByCId(c *gin.Context) {
	cid := c.Param("cid")
	pagenum, pagesize := getPageInfo(c)
	art, err := service.GetArticlesListByCId(cid, pagenum, pagesize)
	if err != nil {
		zap.L().Error("service.GetArticlesListByCId(cid, pagenum, pagesize) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, art)
}

func UpdateArticleHandler(c *gin.Context) {
	articleId := c.Param("id")
	m := new(models.ParamArticle)
	if err := c.ShouldBindJSON(m); err != nil {
		zap.L().Error("UpdateArticleHandler with valid params", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	_, err := service.UpdateArticle(articleId, m)
	if err != nil {
		zap.L().Error("service.UpdateArticle(articleId, m) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)

}
func GetArticleByCateId(c *gin.Context) {
	id := c.Param("id")
	article, err := service.GetArticleByCateId(id)
	if err != nil {
		zap.L().Error("service.GetArticleByCateId(id) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, article)
}
func DeleteArticleById(c *gin.Context) {
	id := c.Param("id")
	_, err := service.DeleteArticleById(id)
	if err != nil {
		zap.L().Error("service.DeleteArticleById(id) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}
func GetArticleListByTitle(c *gin.Context) {
	title := c.Query("title")
	pagenum, pagesize := getPageInfo(c)
	if title == "" {
		zap.L().Error("GetArticleListByTitle with valid params")
		ResponseError(c, CodeInvalidParam)
		return
	}
	resppage, err := service.GetArticleListByTitle(pagenum, pagesize, title)
	if err != nil {
		zap.L().Error("service.GetArticleListByTitle(pagenum, pagesize, m) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, resppage)
}
