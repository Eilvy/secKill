package tool

import (
	"context"
	"fmt"
	"go_code/seckill-service/dao"
	"go_code/seckill-service/rpc/kitex_gen/itemcenter"
)

var (
	ctx      = context.Background()
	consumer = "consumer"
	stream   = "stream"
	group    = "group"
	start    = ">"  //">"指从指定流的最新未消费消息的ID
	count    = 1000 //消息队列最多能容纳的消息数，根据实际情况更改
)

func RedisMq(u *itemcenter.User) error {
	//创建消费者组
	err := dao.RDB0.XGroupCreateMkStream(ctx, stream, group, "0-0").Err()
	if err != nil {
		fmt.Println("create mq error : ", err)
		return err
	}

	return nil
}
