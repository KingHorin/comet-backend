package main

import (
	"comet-server/api"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	userGroup := r.Group("/api/user")
	{
		userGroup.POST("/login", api.LoginJSON())
		userGroup.POST("/register", api.RegisterJSON())
		userGroup.POST("/phone", api.UpdatePhoneJSON())
		userGroup.POST("/email", api.UpdateEmailJSON())
	}
	orderGroup := r.Group("/api/order")
	{
		orderGroup.POST("/addAddressInfo", api.UpdateAddressJSON())
		orderGroup.POST("/postOrderInfo", api.SubmitOrder())
		orderGroup.GET("/getPayStatus", api.GetPayStatus())
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
