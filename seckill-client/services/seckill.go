package services

import (
	"context"
	"fmt"
	"go_code/seckill/seckill-client/rpc"
	"go_code/seckill/seckill-client/rpc/kitex_gen/itemcenter"
)

func SkStart(user *itemcenter.User) bool {
	OK, err := rpc.ItemCenter.SecKill(context.Background(), user)
	if err != nil {
		fmt.Println("secKill client error : ", err.Error())
		return false
	}
	if OK {
		return OK
	}
	return false
}
