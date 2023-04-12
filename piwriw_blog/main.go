package main

import (
	"fmt"
	"piwriw_blog/controller"
	"piwriw_blog/dao/mysql"
	"piwriw_blog/logger"
	"piwriw_blog/router"
	"piwriw_blog/setting"
)

// @title Piwriw Blog项目接口文档
// @version 1.0
// @description Go Piwriw Blog
// @termsOfService http://swagger.io/terms/

// @contact.name Piwriw
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	//if len(os.Args) < 2 {
	//	fmt.Println("need config file.eg: new-ec-dashboard config.yaml")
	//	return
	//}
	// 加载配置
	if err := setting.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	//if err := redis.Init(setting.Conf.RedisConfig); err != nil {
	//	fmt.Printf("init redis failed, err:%v\n", err)
	//	return
	//}
	//defer redis.Close()

	// 初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}
	// 注册路由
	r := router.SetupRouter(setting.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
