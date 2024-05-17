package routers

import (
	"github.com/gin-gonic/gin"
	"go_code/seckill/seckill-client/api"
	"go_code/seckill/seckill-client/middleware"
)

func InitRouters() {
	r := gin.Default()

	u := r.Group("/user")
	{
		u.POST("/register", api.Register)
		u.POST("/login", api.Login)
		u.POST("/refresh", middleware.CheckUser(), api.RefreshToken)
	}
	sk := r.Group("/sk")
	{
		sk.GET("/skstart", middleware.UserLimit(), api.SkStart)
	}

	r.Run()
}
