package home

import (
	"diy-server/models"
	"fmt"
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

// 获取单个颜色详情
func GetColor(c *gin.Context) {
	id := c.Query("id")
	var colorInfo models.GoodsImg
	has, err := db.Cols("front", "contrary").ID(id).Get(&colorInfo)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	if !has {
		msg := "该颜色信息不存在"
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  msg,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": colorInfo,
	})
}

// 获取图案列表
func GetPicture(c *gin.Context) {
	id := c.Query("id")

	var goods models.Goods
	has, err := db.Cols("picture_list").ID(id).Get(&goods)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	if !has {
		msg := "该商品不存在"
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  msg,
		})
		return
	}
	var pictureList = make([]*models.Picture, 0)
	err2 := db.Where(fmt.Sprintf("id in (%s)", goods.PictureList)).Find(&pictureList)
	if err2 != nil {
		log.Error(err2.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err2.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": pictureList,
	})
}
