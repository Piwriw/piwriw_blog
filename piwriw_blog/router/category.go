package router

import (
	"github.com/gin-gonic/gin"
	"piwriw_blog/controller"
)

func CategoryRouter(r *gin.Engine) {
	categoryGroup := r.Group("/category")
	{
		categoryGroup.GET("/all", controller.GetCategoriesHandler)
	}

}
