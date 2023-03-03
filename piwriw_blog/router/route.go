package router

import (
	"net/http"
	"piwriw_blog/logger"
	"piwriw_blog/middlewares"
	//"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// SetupRouter 路由
func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()

	//r.Use(logger.GinLogger(), logger.GinRecovery(true), m iddlewares.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.Use(middlewares.Cors())
//
// 	r.LoadHTMLFiles("./template/admin.html", "./template/index.html")
// 	r.StaticFile("/favicon.ico", "./favicon.ico")
//
// 	r.Static("/static", "./static")
//
// 	r.GET("/", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.html", nil)
// 	})
//
// 	r.GET("/admin", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "admin.html", nil)
// 	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//注册路由
	UserRouter(r)
	ArticleRouter(r)
	CategoryRouter(r)
	ProfileRouter(r)
	CommentRouter(r)
	AdminRouter(r)

	pprof.Register(r) // 注册pprof相关路由

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
