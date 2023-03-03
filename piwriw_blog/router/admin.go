package router

import (
	"github.com/gin-gonic/gin"
	"piwriw_blog/controller"
	"piwriw_blog/middlewares"
)

func AdminRouter(r *gin.Engine) {
	adminGroup := r.Group("/admin")
	{
		adminGroup.POST("/login", controller.AdminLoginHandler)

		adminGroup.Use(middlewares.JWTAuthMiddleware())
		adminGroup.GET("/article", controller.GetAllArticleHandler)
		adminGroup.GET("/category", controller.GetCategoriesHandler)
		adminGroup.GET("/category/list", controller.GetCategoryListHandler)
		adminGroup.POST("/article/add", controller.AddArticleHandler)
		adminGroup.POST("/upload", controller.UploadImgHandler)
		adminGroup.PUT("/article/:id", controller.UpdateArticleHandler)
		adminGroup.GET("/article/:id", controller.GetArticleById)
		adminGroup.DELETE("/article/:id", controller.DeleteArticleById)
		adminGroup.GET("/article/list/:id", controller.GetArticleByCateId)
		adminGroup.POST("/category/add", controller.AddCategoryHandler)
		adminGroup.PUT("/category/:id", controller.UpdateCategoryHandler)
		adminGroup.GET("/category/:id", controller.GetCategoryById)
		adminGroup.DELETE("/category/:id", controller.DeleteCategoryById)
		adminGroup.GET("/comment/list", controller.GetAllComments)
		adminGroup.GET("/profile/:id", controller.GetProfile)
		adminGroup.PUT("/profile/:id", controller.UpdateProfileHandler)
		adminGroup.GET("/users", controller.GetAllUsersHandler)
		adminGroup.PUT("/checkcomment/:id/:status", controller.CheckCommentHandler)
		adminGroup.GET("/comment/info/:id", controller.GetCommentStatusById)
		adminGroup.DELETE("/user/:id", controller.DeleteUserByIdHandler)
		adminGroup.GET("/user/:id", controller.GetUserByIdHandler)
		adminGroup.PUT("/user/:id", controller.UpdateUserByIdHandler)
		adminGroup.PUT("/user/changepw/:id", controller.UpdateUserPasswordHandler)
		adminGroup.GET("/user", controller.GetUserListHandler)
	}

}
