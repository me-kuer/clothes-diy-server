package home

import (
	"diy-server/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// 配置微信
const (
	AppID     = "wx4e68b7770373c61b"
	AppSecret = "5d60d4684d550a0659bbbeff47c40727"
)

type ResponseInfo struct {
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
	SessionKey string `json:"session_key"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

func UserLogin(c *gin.Context) {
	code := c.GetString("code")
	res, err := getOpenid(code)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	// 将json解析为map
	var resInfo ResponseInfo
	err2 := json.Unmarshal([]byte(res), &resInfo) //第二个参数要地址传递

	if err2 != nil {
		log.Error(err2.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err2.Error(),
		})
		return
	}
	// 判断微信授权的返回码
	if resInfo.Errcode != 0 {
		// 如果返回错误码，则返回给前端
		errMsg := "错误码：" + strconv.Itoa(resInfo.Errcode) + "，错误信息：" + resInfo.Errmsg
		log.Error(errMsg)
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  errMsg,
		})
		return
	}
	// 判断数据库中是否有该openid，如果没有则添加，如果已存在则直接返回token
	var user models.Users
	has, err3 := db.Cols("id").Where("openid=?", resMap["openid"].(string)).Get(&user)
	if err3 != nil {
		log.Error(err3.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err3.Error(),
		})
		return
	}
	// 判断是否有该条记录，如果没有则进行创建用户
	if !has {
		var user = models.Users{
			Openid: resInfo.Openid,
			Unionid: resInfo.Unionid,
			SessionKey: resInfo.SessionKey,
			RegisterTime: strconv.FormatInt(time.Now().Unix(),10),
		}
	}

}

// 保存用户信息
func SaveUserInfo(c *gin.Context) {

}

// 通过code 获取 openid, access_token
func getOpenid(code string) (string, error) {
	var url = "https://api.weixin.qq.com/sns/jscode2session?appid=" + AppID + "&secret=" + AppSecret + "&js_code=" + code + "&grant_type=authorization_code"
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	robots, err2 := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err2 != nil {
		return "", err2
	}
	return string(robots), nil
}
