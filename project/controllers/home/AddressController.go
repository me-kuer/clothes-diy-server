package home

import (
	"diy-server/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 获取地址列表
func GetAddrList(c *gin.Context) {
	userId, has := c.Get("user_id")
	if !has {
		msg := "user_id不存在"
		log.Error(msg)
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  msg,
		})
		return
	}

	var addrList []*models.Address

	err := db.Where("user_id=?", userId).OrderBy("id desc").Find(&addrList)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  addrList,
	})

}

// 修改
func UpdateAddr(c *gin.Context) {
	userId, has := c.Get("user_id")
	if !has {
		msg := "user_id不存在"
		log.Error(msg)
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  msg,
		})
		return
	}
	// 获取POST数据
	name := c.Query("name")
	tel := c.Query("tel")
	province := c.Query("province")
	city := c.Query("city")
	region := c.Query("region")
	detail := c.Query("detail")
	sex, _ := strconv.Atoi(c.Query("sex"))

	var addr = models.Address{
		Name:     name,
		Tel:      tel,
		Sex:      int8(sex),
		Province: province,
		City:     city,
		Region:   region,
		Detail:   detail,
	}

	_, err := db.Where("user_id=?", userId).Update(&addr)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "保存成功",
	})
}

// 添加
func AddAddress(c *gin.Context) {
	userId, has := c.Get("user_id")
	if !has {
		msg := "user_id不存在"
		log.Error(msg)
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  msg,
		})
		return
	}

	// 获取POST数据
	name := c.Query("name")
	tel := c.Query("tel")
	province := c.Query("province")
	city := c.Query("city")
	region := c.Query("region")
	detail := c.Query("detail")
	sex, _ := strconv.Atoi(c.Query("sex"))

	var addr = models.Address{
		Name:     name,
		Tel:      tel,
		Sex:      int8(sex),
		Province: province,
		City:     city,
		Region:   region,
		Detail:   detail,
		UserId:   userId.(int),
	}

	_, err := db.InsertOne(&addr)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}
