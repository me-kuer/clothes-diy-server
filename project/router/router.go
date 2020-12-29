package router

import (
	"diy-server/controllers/admin"
	"diy-server/controllers/home"
	"diy-server/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(middlewares.Cors())

	router.GET("/user/login", home.UserLogin)

	homeGroup := router.Group("/")
	homeGroup.Use(middlewares.CheckUserToken())
	{
		// 获取首页
		homeGroup.GET("/goods/list", home.GoodsList)
		// 创建订单
		homeGroup.GET("/order/create", home.CreateOrder)
		// 文件上传
		homeGroup.POST("/upload/base64", home.UploadBase64)
		homeGroup.POST("/upload/file", home.UploadFile)

	}

	router.GET("/admin/user/login", admin.AdminLogin)
	adminGroup := router.Group("/admin")
	adminGroup.Use(middlewares.CheckAdminToken())
	{
		// 文件上传
		//homeGroup.POST("/upload/base64", home.UploadBase64)
		//homeGroup.POST("/upload/file", home.UploadFile)
	}


	// 启动并监听8080端口
	router.Run(":8090")
}
