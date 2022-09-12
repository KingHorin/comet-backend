package middleware

import (
	"comet-server/pkg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var Jwtkey = []byte("TemporaryKey#2022-09-11")

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// CreateToken 生成token
func CreateToken(username string) (string, error) {
	claims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(), //过期时间戳，默认设置2小时
			IssuedAt:  time.Now().Unix(),                    //当前时间戳
			Issuer:    "comet",                              //颁发者签名
			Subject:   "userToken",                          //签名主题
		},
	}
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenStruct.SignedString(Jwtkey)
}

// CheckToken 验证token
func CheckToken(token string) (*MyClaims, bool) {
	tokenObj, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Jwtkey, nil
	})
	if key, _ := tokenObj.Claims.(*MyClaims); tokenObj.Valid {
		return key, true
	} else {
		return nil, false
	}
}

// JwtMiddleware jwt中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从请求头中获取token
		tokenStr := c.Request.Header.Get("Authorization")
		//用户不存在
		if tokenStr == "" {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "无token，非法访问"})
			c.Abort()
			return
		}
		//token格式错误
		tokenSlice := strings.SplitN(tokenStr, " ", 2)
		if len(tokenSlice) != 2 && tokenSlice[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "token格式错误"})
			c.Abort()
			return
		}
		//验证token
		tokenStruck, ok := CheckToken(tokenSlice[1])
		tokenInRedis, err := pkg.GetRD().HGet("token", tokenStruck.Username).Result()
		ok = ok && err == nil && tokenInRedis == tokenStr[7:]
		if !ok {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "token不正确"})
			c.Abort()
			return
		}
		//token超时
		if time.Now().Unix() > tokenStruck.ExpiresAt {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "token过期"})
			c.Abort()
			return
		}
		c.Set("tokenUsername", tokenStruck.Username)

		c.Next()
	}
}
