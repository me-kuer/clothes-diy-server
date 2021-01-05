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
	AppID     = "wx65ae05da4969cd7b"
	AppSecret = "770753c16b9b1a7f208edb55b0df2f6a"
)

type ResponseInfo struct {
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
	SessionKey string `json:"session_key"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

func UserLogin(c *gin.Context) {
	code := c.Query("code")

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
	has, err3 := db.Cols("id").Where("openid=?", resInfo.Openid).Get(&user)
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
		user = models.Users{
			Openid:       resInfo.Openid,
			Unionid:      resInfo.Unionid,
			SessionKey:   resInfo.SessionKey,
			RegisterTime: strconv.FormatInt(time.Now().Unix(), 10),
		}
		_, err4 := db.InsertOne(&user)
		if err4 != nil {
			log.Error(err4.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  err4.Error(),
			})
			return
		}
	}

	// 生成token
	token, err5 := jwt.SetUserId(user.Id).SetExpireTime(2 * 60 * 60).EncodeToken()
	if err5 != nil {
		log.Error(err5.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err5.Error(),
		})
		return
	}
	// 下发token
	c.Header("Authorization", token)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
	})
}

// 保存用户信息
func SaveUserInfo(c *gin.Context) {
	nickname := c.Query("nickname")
	headPic := c.Query("head_pic")
	// 获取user_id
	userId, has := c.Get("user_id")
	if !has {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "user_id不存在",
		})
		return
	}

	// 修改
	var user = models.Users{
		Nickname: nickname,
		HeadPic:  headPic,
	}
	_,err := db.Id(userId).Update(&user)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	// 返回成功
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "用户信息修改成功",
	})
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
