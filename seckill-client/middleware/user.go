package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/seckill/seckill-client/model"
	"go_code/seckill/seckill-client/resps"
	"go_code/seckill/seckill-client/rpc/kitex_gen/itemcenter"
	"go_code/seckill/seckill-client/services"
	"strings"
)

func CheckUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			resps.ParseErr(c)
			c.Abort()
			return
		}
		parts := strings.SplitN(authorization, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			resps.ParseErr(c)
			c.Abort()
			return
		}
		//parts[1]是token本体
		username, err := services.ParseToken(parts[1], model.Token)
		if err != nil {
			resps.InternalErr(c)
			return
		}
		password, err := services.ParseToken(parts[1], model.Password)
		if err != nil {
			resps.InternalErr(c)
			return
		}
		Id, err := services.ParseToken(parts[1], 5000)
		if err != nil {
			resps.InternalErr(c)
			return
		}
		typ, err := services.ParseToken(parts[1], 6000)
		if err != nil {
			resps.InternalErr(c)
			return
		}
		c.Set("username", username)
		c.Set("password", password)
		c.Set("id", Id)
		c.Set("typ", typ)
		//resp := "welcome back " + username
		//resps.OKWithData(c, resp)
		c.Next()
	}
}

func UserLimit() func(c *gin.Context) {
	return func(c *gin.Context) {
		//CheckUser()
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			resps.ParseErr(c)
			c.Abort()
			return
		}
		parts := strings.SplitN(authorization, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			resps.ParseErr(c)
			c.Abort()
			return
		}
		//parts[1]是token本体
		username, err := services.ParseToken(parts[1], model.Token)
		if err != nil {
			resps.InternalErr(c)
			return
		}
		password, err := services.ParseToken(parts[1], model.Password)
		if err != nil {
			resps.InternalErr(c)
			return
		}
		Id, err := services.ParseToken(parts[1], 5000)
		if err != nil {
			resps.InternalErr(c)
			return
		}
		typ, err := services.ParseToken(parts[1], 6000)
		if err != nil {
			resps.InternalErr(c)
			return
		}
		c.Set("username", username)
		c.Set("password", password)
		c.Set("id", Id)
		c.Set("typ", typ)
		user := &itemcenter.User{
			Username: c.MustGet("username").(string),
			Password: c.MustGet("password").(string),
		}
		if err := services.InRedisMq(user); err != nil {
			resps.ParamErr(c)
			fmt.Println("redisMq error : ", err)
			return
		}
		services.TokenBucket(services.Num, c)
	}
}
