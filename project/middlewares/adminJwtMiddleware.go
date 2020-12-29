package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 管理员校验jwt
func CheckAdminToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token = c.GetHeader("Authorization")
		// 如果没有token 则表示
		if token == "" {
			// 返回未登录
			c.JSON(http.StatusOK, gin.H{
				"code": "401",
				"msg":  "对不起，您还没有登录",
			})
		}
		// 开始校验token
		claims, err := jwt.DecodeToken(token)
		if err != nil {
			log.Error(err.Error())
			// 返回未登录
			c.JSON(http.StatusOK, gin.H{
				"code": "401",
				"msg":  "token校验失败",
			})
		}
		// 判断 admin_id 是否存在
		if claims == nil || claims.AdminId <= 0 {
			// 返回未登录
			c.JSON(http.StatusOK, gin.H{
				"code": "500",
				"msg":  "token 解码失败",
			})
		}
		// 判断 token 是否过期
		if claims != nil && claims.ExpiresAt < time.Now().Unix() {
			// 返回未登录
			c.JSON(http.StatusOK, gin.H{
				"code": "401",
				"msg":  "登录状态已失效",
			})
		}
		// 如果token距离过期还有30分钟，则重新颁发token
		// 公式： 过期时间 - 当前时间 <= 30分钟
		if claims != nil && (claims.ExpiresAt-time.Now().Unix() <= 30*60) {
			newToken, err2 := jwt.SetAdminId(claims.AdminId).SetExpireTime(2 * 60 * 60).EncodeToken()
			if err2 != nil {
				log.Error(err2.Error())
				c.JSON(http.StatusOK, gin.H{
					"code": "401",
					"msg":  err2.Error(),
				})
			}
			// 将新token保存到响应header头中
			c.Header("Authorization", newToken)
		}

		c.Next()
	}
}

