package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 普通用户校验jwt
func CheckUserToken() gin.HandlerFunc {

	return func(c *gin.Context) {
		var token = c.GetHeader("Authorization")

		// 如果没有token 则表示
		if token == "" {
			// 返回未登录
			c.JSON(http.StatusOK, gin.H{
				"code": "401",
				"msg":  "对不起，您还没有登录",
			})
			// 终止继续执行下面程序
			c.Abort()
			return
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
			c.Abort()
			return
		}
		// 判断 user_id 是否存在
		if claims == nil || claims.UserId <= 0 {
			// 返回未登录
			c.JSON(http.StatusOK, gin.H{
				"code": "500",
				"msg":  "token 解码失败",
			})
			c.Abort()
			return
		}

		// 判断 token 是否过期
		if claims.ExpiresAt < time.Now().Unix() {
			// 返回未登录
			c.JSON(http.StatusOK, gin.H{
				"code": "401",
				"msg":  "登录状态已失效",
			})
			c.Abort()
			return
		}
		// 将解析的userId 保存到 gin.Context中
		c.Set("user_id", claims.UserId)

		// 如果token距离过期还有30分钟，则重新颁发token
		// 公式： 过期时间 - 当前时间 <= 30分钟
		if claims.ExpiresAt - time.Now().Unix() <= 30 * 60 {
			newToken, err2 := jwt.SetUserId(claims.UserId).SetExpireTime(2 * 60 * 60).EncodeToken()
			if err2 != nil {
				log.Error(err2.Error())
				c.JSON(http.StatusOK, gin.H{
					"code": "401",
					"msg":  err2.Error(),
				})
				c.Abort()
				return
			}
			// 将新token保存到响应header头中
			c.Header("Authorization", newToken)
		}

		c.Next()
	}
}
