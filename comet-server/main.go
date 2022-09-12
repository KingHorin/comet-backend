package main

import (
	"comet-server/api"
	"comet-server/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	userGroup := r.Group("/api/user")
	userGroup.Use()
	{
		userGroup.POST("/login", api.Login())
		userGroup.POST("/register", api.Register())
	}
	userGroup.Use(middleware.JwtMiddleware())
	{
		userGroup.POST("/phone", api.UpdatePhone())
		userGroup.POST("/email", api.UpdateEmail())
	}
	orderGroup := r.Group("/api/order").Use(middleware.JwtMiddleware())
	{
		orderGroup.POST("/addAddressInfo", api.UpdateAddress())
		orderGroup.POST("/postOrderInfo", api.SubmitOrder())
		orderGroup.GET("/getPayStatus", api.GetPayStatus())
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
