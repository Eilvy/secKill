package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/seckill/seckill-client/model"
	"go_code/seckill/seckill-client/resps"
	"go_code/seckill/seckill-client/rpc"
	"go_code/seckill/seckill-client/rpc/kitex_gen/itemcenter"
	"go_code/seckill/seckill-client/services"
)

func Register(c *gin.Context) {
	u := model.Register{}
	if err := c.ShouldBind(&u); err != nil {
		resps.ParamErr(c)
		return
	}

	if err := services.Register(u); err != nil {
		resps.InternalErr(c)
		return
	}
	resps.OK(c)
}

func Login(c *gin.Context) {
	u := model.Login{}
	err := c.ShouldBind(&u)
	if err != nil {
		resps.ParamErr(c)
		return
	}
	accessToken, refreshToken, err := services.Login(u)
	if err != nil {
		resps.ReceivedErr(c)
		return
	}
	c.Request.Header.Set("Authorization", "Bearer "+accessToken)
	tokenMap := map[string]string{
		"accessToken":   accessToken,
		"refreshToken":  refreshToken,
		"welcome back ": u.Username,
	}
	resps.OKWithData(c, tokenMap)
}

func RefreshToken(c *gin.Context) {
	if c.MustGet("typ").(int64) != model.RefreshToken {
		resps.ParamErr(c)
		fmt.Println("token wrong to access")
		return
	}
	u := model.Login{
		Username: c.MustGet("username").(string),
		Password: c.MustGet("password").(string),
	}
	accessToken, err := rpc.ItemCenter.CreateToken(context.Background(), &itemcenter.User{
		Username: u.Username,
		Password: u.Password,
	}, model.Token)
	if err != nil {
		fmt.Println("error : ", err.Error())
		return
	}
	refreshToken, err := rpc.ItemCenter.CreateToken(context.Background(), &itemcenter.User{
		Username: u.Username,
		Password: u.Password,
	}, model.RefreshToken)
	if err != nil {
		fmt.Println("error : ", err.Error())
		return
	}
	c.Request.Header.Set("Authorization", "Bearer "+accessToken)
	tokenMap := map[string]string{
		"accessToken":   accessToken,
		"refreshToken":  refreshToken,
		"welcome back ": u.Username,
	}
	resps.OKWithData(c, tokenMap)
}
