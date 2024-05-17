package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/seckill/seckill-client/model"
	"go_code/seckill/seckill-client/resps"
	"go_code/seckill/seckill-client/rpc/kitex_gen/itemcenter"
	"go_code/seckill/seckill-client/services"
	"strconv"
)

func SkStart(c *gin.Context) {
	token, err := services.RDB1.Get(context.Background(), c.MustGet("username").(string)).Result()
	if err != nil {
		resps.ParamErr(c)
		return
	}
	typ, err := services.ParseToken(token, 6000)
	if err != nil {
		resps.InternalErr(c)
		return
	}
	if typ != strconv.FormatInt(model.SKToken, 10) {
		fmt.Println("token is not correct ")
		return
	}
	username, err := services.ParseToken(token, model.Token)
	if err != nil {
		resps.InternalErr(c)
		return
	}
	password, err := services.ParseToken(token, model.Password)
	if err != nil {
		resps.InternalErr(c)
		return
	}
	//Id, err := services.ParseToken(token, 5000)
	//if err != nil {
	//	resps.InternalErr(c)
	//	return
	//}
	OK := services.SkStart(&itemcenter.User{
		Username: username,
		Password: password,
	})
	if OK {
		resps.OKWithData(c, "抢购成功")
	}
}
