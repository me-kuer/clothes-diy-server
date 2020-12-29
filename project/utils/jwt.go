package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 指定加密密钥
var jwtSecret = []byte("diy-server")

// 定义Jwt结构体
type Jwt struct {
	adminId    int
	userId     int
	expireAt int64
}

//Claim是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	UserId  int `json:"user_id"`
	AdminId int `json:"admin_id"`
	jwt.StandardClaims
}

// 编码token
func (jwtInstance *Jwt) EncodeToken() (string, error) {

	claims := Claims{
		UserId:  jwtInstance.userId,
		AdminId: jwtInstance.adminId,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: jwtInstance.expireAt,
			// 指定token发行人
			Issuer: "mantong",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 设置user_id
func (jwtInstance *Jwt) SetUserId(UserId int) *Jwt {
	jwtInstance.userId = UserId
	return jwtInstance
}

// 设置admin_id
func (jwtInstance *Jwt) SetAdminId(AdminId int) *Jwt {
	jwtInstance.userId = AdminId
	return jwtInstance
}

// 设置过期时间
func (jwtInstance *Jwt) SetExpireTime(ExpireTime int64) *Jwt {
	//设置token有效时间
	nowTime := time.Now().Unix()
	jwtInstance.expireAt = nowTime + ExpireTime
	return jwtInstance
}

// 解码token
// 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func (jwtInstance *Jwt)  DecodeToken(token string) (*Claims, error) {

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims == nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
