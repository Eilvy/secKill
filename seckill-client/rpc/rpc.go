package rpc

import (
	"github.com/cloudwego/kitex/client"
	"go_code/seckill/seckill-client/rpc/kitex_gen/itemcenter/itemcenter"
	"log"
)

var ItemCenter itemcenter.Client

func InitRPC() {
	item, err := itemcenter.NewClient("itemcenter", client.WithHostPorts("localhost:8888"))
	if err != nil {
		log.Println("连接服务端失败,err:", err)
	}
	ItemCenter = item
}
