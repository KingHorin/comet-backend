package api

import (
	"comet-server/middleware"
	"comet-server/model"
	"comet-server/pkg"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
)

func ValidateUser(c *gin.Context, username string) bool {
	tokenUsername, _ := c.Get("tokenUsername")
	return utils.ToString(tokenUsername) == username
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		type loginReq struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Captcha  string `json:"captcha"`
		}
		r := loginReq{}
		c.BindJSON(&r)
		db, rd := pkg.GetDB(), pkg.GetRD()
		user := model.User{}
		err := db.Where("username = ?", r.Username).Take(&user).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(200, gin.H{"code": 0, "msg": "用户名不存在"})
		} else if user.Password != r.Password {
			c.JSON(200, gin.H{"code": 0, "msg": "密码错误"})
		} else {
			token, _ := middleware.CreateToken(r.Username)
			c.JSON(200, gin.H{"code": 1, "msg": "登录成功", "token": token})
			rd.HSet("token", r.Username, token)
		}
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		type registerReq struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Captcha  string `json:"captcha"`
		}
		r := registerReq{}
		c.BindJSON(&r)
		db := pkg.GetDB()

		if err := db.Where("username = ?", r.Username).Take(&model.User{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(200, gin.H{"code": 0, "msg": "用户名已存在"})
		} else {
			c.JSON(200, gin.H{"code": 1, "msg": "注册成功"})
			db.Create(&model.User{Username: r.Username, Password: r.Password})
		}
	}
}

func UpdatePhone() gin.HandlerFunc {
	return func(c *gin.Context) {
		type phoneReq struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Phone    string `json:"phone"`
			token    string `json:"token"`
		}
		r := phoneReq{}
		c.BindJSON(&r)
		if !ValidateUser(c, r.Username) {
			c.JSON(200, gin.H{"code": 0, "msg": "用户token不匹配"})
			return
		}
		db := pkg.GetDB()
		user := model.User{}
		db.Where("username = ?", r.Username).Take(&user)

		if user.Password != r.Password {
			c.JSON(200, gin.H{"code": 0, "msg": "密码错误"})
		} else {
			c.JSON(200, gin.H{"code": 1, "msg": "手机号码绑定成功"})
			db.Model(&model.User{}).Where("id = ?", user.Id).Update("phone", r.Phone)
		}
	}
}

func UpdateEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		type emailReq struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Email    string `json:"email"`
			Token    string `json:"token"`
		}
		r := emailReq{}
		c.BindJSON(&r)
		if !ValidateUser(c, r.Username) {
			c.JSON(200, gin.H{"code": 0, "msg": "用户token不匹配"})
			return
		}
		db := pkg.GetDB()
		user := model.User{}
		db.Where("username = ?", r.Username).Take(&user)

		if user.Password != r.Password {
			c.JSON(200, gin.H{"code": 0, "msg": "密码错误"})
		} else {
			c.JSON(200, gin.H{"code": 1, "msg": "邮箱绑定成功"})
			db.Model(&model.User{}).Where("id = ?", user.Id).Update("email", r.Email)
		}
	}
}
