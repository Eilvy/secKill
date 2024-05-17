package main

import (
	"go_code/seckill/seckill-client/routers"
	"go_code/seckill/seckill-client/rpc"
	"go_code/seckill/seckill-client/services"
)

func main() {
	services.InitRedis()
	rpc.InitRPC()
	routers.InitRouters()
}
