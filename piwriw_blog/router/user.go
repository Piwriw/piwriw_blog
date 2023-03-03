package router

import (
	"github.com/gin-gonic/gin"
	"piwriw_blog/controller"
)

func UserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login/front", controller.UserLoginHandler)
		userGroup.POST("/add", controller.RegisterUserHandler)
	}

}
