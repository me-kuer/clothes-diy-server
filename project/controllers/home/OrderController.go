package home

import (
	"diy-server/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

// 创建订单
func CreateOrder(c *gin.Context) {
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
	// 获取的字段 【goods_id，尺码，正面款式, 反面款式， 正面图案，反面图案，正面合成图，反面合成图，】
	goodsId, _ := strconv.Atoi(c.Query("goods_id"))
	size := c.Query("size")
	front := c.Query("front")
	contrary := c.Query("contrary")
	frontPicture := c.Query("front_picture")
	contraryPicture := c.Query("contrary_picture")
	frontCompose := c.Query("front_compose")
	contraryCompose := c.Query("contrary_compose")
	consigneeName := c.Query("consignee_name")
	consigneeTel := c.Query("consignee_tel")
	consigneeSex := c.Query("consignee_sex")
	consigneeProvince := c.Query("consignee_province")
	consigneeCity := c.Query("consignee_city")
	consigneeRegion := c.Query("consignee_region")
	consigneeDetail := c.Query("consignee_detail")

	// 通过goods_id 查询该商品的 名称, 价格
	var goodsInfo models.Goods
	exists, err := db.Id(goodsId).Get(&goodsInfo)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	if !exists {
		msg := "该商品不存在"
		log.Error(msg)
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  msg,
		})
		return
	}

	// 创建订单
	var order = models.Orders{
		OutTradeNo: getOrderNo(),
		UserId:            userId.(int),
		GoodsId:           goodsId,
		GoodsName:         goodsInfo.Name,
		Total:             goodsInfo.Price,
		Size:              size,
		Front:             front,
		FrontPicture:      frontPicture,
		FrontCompose:      frontCompose,
		Contrary:          contrary,
		ContraryPicture:   contraryPicture,
		ContraryCompose:   contraryCompose,
		ConsigneeName:     consigneeName,
		ConsigneeTel:      consigneeTel,
		ConsigneeSex:      consigneeSex,
		ConsigneeProvince: consigneeProvince,
		ConsigneeCity:     consigneeCity,
		ConsigneeRegion:   consigneeRegion,
		ConsigneeDetail:   consigneeDetail,
		CreateTime:        strconv.FormatInt(time.Now().Unix(), 10),
	}
	fmt.Println(order)
	_, err2 := db.InsertOne(&order)
	if err2 != nil {
		log.Error(err2.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg": err2.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建订单成功",
		"order_id": order.Id,
	})
}

func GetOrderList(c *gin.Context) {
	// 获取user_id
	userId, has := c.Get("user_id")
	if !has {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "user_id不存在",
		})
		return
	}

	status := c.DefaultQuery("status", "")
	refreshTime := c.DefaultQuery("refresh_time", "0")
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	offset := (page - 1) * pageSize

	var where = "user_id=? and create_time>?"
	if status != "" {
		where += " and status=" + status
	}
	var orders []*models.Orders
	err := db.Cols("id", "status", "goods_name", "front_compose", "contrary_compose", "total", "create_time").
		Where(where, userId.(int), refreshTime).
		OrderBy("id desc").
		Limit(pageSize, offset).
		Find(&orders)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	// 获取总数
	total, err2 := db.Where(where, userId.(int)).Count(&orders)
	if err2 != nil {
		log.Error(err2.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err2.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"total": total,
		"data":  orders,
	})
}

func OrderDetail(c *gin.Context) {
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

	id := c.Query("id")

	var order models.Orders
	has, err := db.Cols(
		"id",
		"front_compose",
		"contrary_compose",
		"goods_name",
		"out_trade_no",
		"total",
		"note",
		"status",
		"consignee_name",
		"consignee_tel",
		"consignee_sex",
		"consignee_province",
		"consignee_city",
		"consignee_region",
		"consignee_detail",
		"create_time",
	).
		ID(id).
		Where("user_id=?", userId.(int)).
		Get(&order)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	// 如果没有该条记录
	if !has {
		msg := "该订单不存在"
		log.Error(msg)
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  msg,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": order,
	})
}

// 订单支付
func OrderPay(c *gin.Context) {

}

//生成24位订单号
//前面17位代表时间精确到毫秒，中间3位代表进程id，最后4位代表序号
func getOrderNo() string {
	t := time.Now()
	s := t.Format("20060102150405")
	m := t.UnixNano()/1e6 - t.UnixNano()/1e9*1e3
	ms := sup(m, 3)
	p := os.Getpid() % 1000
	ps := sup(int64(p), 3)
	var num int64
	i := atomic.AddInt64(&num, 1)
	r := i % 10000
	rs := sup(r, 4)
	n := fmt.Sprintf("%s%s%s%s", s, ms, ps, rs)
	return n
}

//对长度不足n的数字前面补0
func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}
