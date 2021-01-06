package router

import (
	"diy-server/controllers/admin"
	"diy-server/controllers/home"
	"diy-server/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(middlewares.Cors())

	// 配置静态文件目录
	router.StaticFS("/static", http.Dir("./static"))
	router.StaticFS("/upload", http.Dir("./upload"))

	// Homer（前端） 分组
	router.GET("/user/login", home.UserLogin)
	// 获取首页
	router.GET("/goods/list", home.GetGoodsList)
	// 获取当前选中的商品属性
	router.GET("/goods/detail", home.GetGoodsDetail)
	// 获取颜色对应的图片
	router.GET("/goods/color", home.GetColor)
	// 获取图案列表
	router.GET("/goods/picture", home.GetPicture)

	homeGroup := router.Group("/")
	homeGroup.Use(middlewares.CheckUserToken())
	{
		// 保存用户 nickname & head_pic
		homeGroup.POST("/user/saveinfo", home.SaveUserInfo)



		// 文件上传
		homeGroup.POST("/upload/base64", home.UploadBase64)
		homeGroup.POST("/upload/file", home.UploadFile)

		// 地址操作
		homeGroup.POST("/addr/add", home.AddAddress)
		homeGroup.POST("/addr/update", home.UpdateAddr)
		homeGroup.POST("/addr/list", home.GetAddrList)

		// 创建订单
		homeGroup.POST("/order/create", home.CreateOrder)
		// 订单列表
		homeGroup.GET("/order/list", home.GetOrderList)
		// 订单详情
		homeGroup.GET("/order/detail", home.OrderDetail)
		// 付款
		homeGroup.GET("/order/pay", home.OrderPay)
	}
	// Admin（后台） 分组
	router.GET("/admin/user/login", admin.AdminLogin)
	adminGroup := router.Group("/admin")
	adminGroup.Use(middlewares.CheckAdminToken())
	{
		// 文件上传
		//homeGroup.POST("/upload/base64", home.UploadBase64)
		//homeGroup.POST("/upload/file", home.UploadFile)
	}
	//

	// 启动并监听8080端口
	router.Run(":8090")
}
