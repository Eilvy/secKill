package services

import (
	"context"
	"fmt"
	"go_code/seckill/seckill-client/model"
	"go_code/seckill/seckill-client/rpc"
	"go_code/seckill/seckill-client/rpc/kitex_gen/itemcenter"
	"log"
)

func Register(u model.Register) (err error) {
	err = rpc.ItemCenter.Register(context.Background(), &itemcenter.User{ //调用kitex服务端的Register
		Username: u.Username,
		Password: u.Password,
		Status:   false,
	})
	if err != nil {
		log.Println("Register rpc call error:", err)
	}
	return err
}
func Login(u model.Login) (accessToken string, refreshToken string, err error) {
	user, err := rpc.ItemCenter.Login(context.Background(), &itemcenter.User{
		Username: u.Username,
		Password: u.Password,
		Id:       u.Id,
	})
	if err != nil {
		fmt.Println("error : ", err.Error())
		return
	}
	accessToken, err = rpc.ItemCenter.CreateToken(context.Background(), user, model.Token)
	if err != nil {
		fmt.Println("error : ", err.Error())
		return
	}
	refreshToken, err = rpc.ItemCenter.CreateToken(context.Background(), user, model.RefreshToken)
	if err != nil {
		fmt.Println("error : ", err.Error())
		return
	}
	return
}
