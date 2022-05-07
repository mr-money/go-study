package Handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/shockerli/cvt"
	"go-study/Model"
	"time"
)

type ApiLoginClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const tokenExpire = int64(time.Hour * 24) //设置过期时间

var keySecret = []byte("Go-Gin-Study123") //加盐秘钥

//
// ApiLoginToken
// @Description: 登录生成jwt
// @param user
// @return string
// @return error
//
func ApiLoginToken(user Model.User) (string, error) {
	// 创建api登录声明
	claims := ApiLoginClaims{
		user.Name, // 自定义字段
		jwt.StandardClaims{
			Audience:  user.Name,             // 受众
			ExpiresAt: tokenExpire,           // 失效时间
			Id:        cvt.String(user.Uuid), // 编号
			IssuedAt:  time.Now().Unix(),     // 签发时间
			Issuer:    "admin",               // 签发人
			NotBefore: time.Now().Unix(),     // 生效时间
			Subject:   "login",               // 主题
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(keySecret)
}

//
// ParseToken
// @Description: 解密jwt
// @param token
// @return *jwt.StandardClaims
// @return error
//
func ParseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return keySecret, nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}