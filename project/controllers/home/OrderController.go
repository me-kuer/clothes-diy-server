package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "创建订单成功",
	})
}
