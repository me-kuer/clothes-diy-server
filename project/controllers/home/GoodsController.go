package home

import (
	"diy-server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GoodsInfo struct {
	models.Goods `xorm:"extends"`
	ColorId      int    `xorm:"'color_id'" json:"color_id,omitempty"`
	Front        string `xorm:"'front'" json:"front,omitempty"`
}

func GetGoodsList(c *gin.Context) {

	var goodsList = make([]*GoodsInfo, 0)

	err := db.Table("goods").
		Alias("g").
		Select("g.id, g.name, g.price, i.id as color_id, i.front").
		Join("left", "goods_img as i", "g.id=i.goods_id").
		GroupBy("g.id").
		OrderBy("g.id desc").
		Find(&goodsList)
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
		"data": goodsList,
	})
}

func GetGoodsDetail(c *gin.Context) {
	id := c.Query("id")

	var imgs = make([]*models.GoodsImg, 0)
	err := db.Where("goods_id=?", id).Find(&imgs)

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
		"data": imgs,
	})
}
