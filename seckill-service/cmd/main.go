package main

import (
	"go_code/seckill-service/dao"
	"go_code/seckill-service/rpc"
	itemcenter "go_code/seckill-service/rpc/kitex_gen/itemcenter/itemcenter"
	"log"
)

func main() {
	dao.InitDB()
	dao.InitRedis()

	svr := itemcenter.NewServer(new(rpc.ItemcenterImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
