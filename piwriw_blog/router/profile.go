package router

import (
	"github.com/gin-gonic/gin"
	"piwriw_blog/controller"
)

func ProfileRouter(r *gin.Engine) {
	profileRouter := r.Group("/profile")
	{
		profileRouter.GET("/:id", controller.GetProfile)
	}

}
