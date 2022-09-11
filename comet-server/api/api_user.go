package api

import (
	"comet-server/model"
	"comet-server/pkg"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ValidateUser(username string) (bool, model.User) {
	db := pkg.GetDB()
	user := model.User{}
	err := db.Where("username = ?", username).Take(&user).Error
	return !errors.Is(err, gorm.ErrRecordNotFound), user
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, captcha := c.Query("username"), c.Query("password"), c.Query("captcha")
		db := pkg.GetDB()
		res := model.User{}
		if len(captcha) != 4 {
			c.JSON(400, "验证码错误")
		} else if db.Model(&model.User{}).Where("username = ?", username).Take(&res).RowsAffected == 0 {
			c.JSON(400, "用户名不存在")
		} else if res.Password != password {
			c.JSON(400, "密码错误")
		} else {
			c.JSON(200, "登录成功")
			db.Model(&model.User{}).Where("id = ?", res.Id).Update("online", 1)
		}
	}
}

func LoginJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		type loginReq struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Captcha  string `json:"captcha"`
		}
		r := loginReq{}
		c.BindJSON(&r)
		valid, user := ValidateUser(r.Username)
		db := pkg.GetDB()
		if len(r.Captcha) != 4 {
			c.JSON(400, "验证码错误")
		} else if !valid {
			c.JSON(400, "用户名不存在")
		} else if user.Password != r.Password {
			c.JSON(400, "密码错误")
		} else {
			c.JSON(200, "登录成功")
			db.Model(&model.User{}).Where("id = ?", user.Id).Update("online", 1)
		}
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, captcha := c.Query("username"), c.Query("password"), c.Query("captcha")
		db := pkg.GetDB()
		if len(captcha) != 4 {
			c.JSON(400, "验证码错误")
		} else if db.Model(&model.User{}).Where("username = ?", username).RowsAffected != 0 {
			c.JSON(400, "用户名已存在")
		} else {
			c.JSON(200, "注册成功")
			db.Model(&model.User{}).Create(&model.User{Username: username, Password: password})
		}
	}
}

func RegisterJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		type registerReq struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Captcha  string `json:"captcha"`
		}
		r := registerReq{}
		c.BindJSON(&r)
		db := pkg.GetDB()
		valid, _ := ValidateUser(r.Username)
		if len(r.Captcha) != 4 {
			c.JSON(400, "验证码错误")
		} else if valid {
			c.JSON(400, "用户名已存在")
		} else {
			c.JSON(200, "注册成功")
			db.Create(&model.User{Username: r.Username, Password: r.Password})
		}
	}
}

func UpdatePhone() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password := c.Query("username"), c.Query("password")
		phone, token := c.Query("phone"), c.Query("token")
		token += "1"
		db := pkg.GetDB()
		u := model.User{}
		if db.Model(&model.User{}).Where("username = ?", username).Take(&u).RowsAffected == 0 {
			c.JSON(400, "用户名不存在")
		} else if u.Password != password {
			c.JSON(400, "密码错误")
		} else {
			c.JSON(200, "手机号码绑定成功")
			db.Model(&model.User{}).Where("id = ?", u.Id).Update("phone", phone)
		}
	}
}

func UpdatePhoneJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		type phoneReq struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Phone    string `json:"phone"`
			token    string `json:"token"`
		}
		r := phoneReq{}
		c.BindJSON(&r)
		db := pkg.GetDB()
		valid, user := ValidateUser(r.Username)
		if !valid {
			c.JSON(400, "用户名不存在")
		} else if user.Password != r.Password {
			c.JSON(400, "密码错误")
		} else {
			c.JSON(200, "手机号码绑定成功")
			db.Model(&model.User{}).Where("id = ?", user.Id).Update("phone", r.Phone)
		}
	}
}

func UpdateEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password := c.Query("username"), c.Query("password")
		email, token := c.Query("email"), c.Query("token")
		token += "1"
		db := pkg.GetDB()
		u := model.User{}
		if db.Model(&model.User{}).Where("username = ?", username).Take(&u).RowsAffected == 0 {
			c.JSON(400, "用户名不存在")
		} else if u.Password != password {
			c.JSON(400, "密码错误")
		} else {
			c.JSON(200, "邮箱绑定成功")
			db.Model(&model.User{}).Where("id = ?", u.Id).Update("email", email)
		}
	}
}

func UpdateEmailJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		type emailReq struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Email    string `json:"email"`
			Token    string `json:"token"`
		}
		r := emailReq{}
		c.BindJSON(&r)
		valid, user := ValidateUser(r.Username)
		db := pkg.GetDB()
		if !valid {
			c.JSON(400, "用户名不存在")
		} else if user.Password != r.Password {
			c.JSON(400, "密码错误")
		} else {
			c.JSON(200, "邮箱绑定成功")
			db.Model(&model.User{}).Where("id = ?", user.Id).Update("email", r.Email)
		}
	}
}
