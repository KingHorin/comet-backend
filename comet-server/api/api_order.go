package api

import (
	"comet-server/model"
	"comet-server/pkg"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func UpdateAddress() gin.HandlerFunc {
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
		if !ValidateUser(c, r.Username) {
			c.JSON(200, gin.H{"code": 0, "msg": "用户token不匹配"})
		} else {
			c.JSON(200, gin.H{"code": 1, "msg": "地址创建成功"})
			db := pkg.GetDB()
			user := model.User{}
			db.Where("username = ?", r.Username).Take(&user)
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
		if !ValidateUser(c, r.Username) {
			c.JSON(200, gin.H{"code": 0, "msg": "用户token不匹配"})
			return
		}
		c.JSON(200, gin.H{"code": 1, "msg": "提交成功，具体功能开发中"})
	}
}

func GetPayStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 0, "msg": "请求已接收，具体功能开发中"})
	}
}
