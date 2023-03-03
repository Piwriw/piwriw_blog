package router

import (
	"github.com/gin-gonic/gin"
	"piwriw_blog/controller"
)

func ArticleRouter(r *gin.Engine) {
	articleGroup := r.Group("/article")
	{

		articleGroup.GET("/front", controller.GetAllArticleHandler)
		articleGroup.GET("/category/:cid", controller.GetArticleListByCId)
		articleGroup.GET("/:id", controller.GetArticleById)
		articleGroup.GET("/search", controller.GetArticleListByTitle)
		//articleGroup.PUT("/:id", controller.UpdateViewHandler)
	}

}
