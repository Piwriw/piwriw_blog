package router

import (
	"github.com/gin-gonic/gin"
	"piwriw_blog/controller"
	"piwriw_blog/middlewares"
)

func CommentRouter(r *gin.Engine) {
	commentRouter := r.Group("/commentfront")
	{
		commentRouter.GET("/:id", controller.GetAllCommentList)
		commentRouter.Use(middlewares.JWTAuthMiddleware())
		commentRouter.POST("/add", controller.AddCommentHandler)
	}

}
