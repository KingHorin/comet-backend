package api

import (
	"comet-server/model"
	"comet-server/pkg"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func UpdateAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, token := c.Query("username"), c.Query("token")
		name, phone := c.Query("name"), c.Query("phone")
		area, address, company := c.Query("area"), c.Query("address"), c.Query("company")
		token += "1"
		db := pkg.GetDB()
		u := model.User{}
		if db.Model(&u).Where("username= ?", username).Take(&u).RowsAffected == 0 {
			c.JSON(400, "用户名不存在")
		} else {
			c.JSON(200, "地址创建成功")
			db.Delete(&model.UserAddress{}, u.Id)
			db.Model(&model.UserAddress{}).Create(&model.UserAddress{
				Id: u.Id, Name: name, Area: area, Address: address, Company: company, Phone: phone})
		}
	}
}

func UpdateAddressJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		type addressReq struct {
			Username string `json:"username"`
			Token    string `json:"token"`
			Name     string `json:"name"`
			Area     string `json:"area"`
			Address  string `json:"address"`
			Company  string `json:"company"`
			Phone    string `json:"phone"`
		}
		r := addressReq{}
		c.BindJSON(&r)
		valid, user := ValidateUser(r.Username)
		db := pkg.GetDB()
		if !valid {
			c.JSON(400, "用户名不存在")
		} else {
			c.JSON(200, "地址创建成功")
			db.Delete(&model.UserAddress{}, user.Id)
			db.Create(&model.UserAddress{Id: user.Id, Name: r.Name, Area: r.Area, Address: r.Address, Company: r.Company, Phone: r.Phone})
		}
	}
}

func SubmitOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		type orderReq struct {
			Username   string          `json:"username"`
			Token      string          `json:"token"`
			Items      json.RawMessage `json:"items"`
			TotalPrice int32           `json:"totalPrice"`
			PayStatus  bool            `json:"payStatus"`
		}
		type itemObj struct {
			Name  string `json:"name"`
			Count int32  `json:"count"`
			Price int32  `json:"price"`
		}
		r := orderReq{}
		c.BindJSON(&r)
		byteArr, _ := r.Items.MarshalJSON()
		var itemArr []*itemObj
		json.Unmarshal(byteArr, &itemArr)
		//for i := 0; i < len(itemArr); i++ {
		//	fmt.Println(itemArr[i].Name)
		//	fmt.Println(itemArr[i].Price)
		//	fmt.Println(itemArr[i].Count)
		//}
		valid, _ := ValidateUser(r.Username)
		if !valid {
			c.JSON(400, "用户名不存在")
		} else {
			c.JSON(200, "提交成功，具体功能开发中")
		}

	}
}

func GetPayStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "请求已接收，具体功能开发中")
	}
}
