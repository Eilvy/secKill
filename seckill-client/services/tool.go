package services

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go_code/seckill/seckill-client/model"
	"go_code/seckill/seckill-client/resps"
	"go_code/seckill/seckill-client/rpc"
	"go_code/seckill/seckill-client/rpc/kitex_gen/itemcenter"
	"golang.org/x/time/rate"
	"time"
)

var (
	Num      int64 = 1000
	RDB0     *redis.Client
	RDB1     *redis.Client
	ctx      = context.Background()
	consumer = "consumer"
	stream   = "secKill"
	group    = "group"
	start    = ">"  //">"指从指定流的最新未消费消息的ID
	count    = 1000 //消息队列最多能容纳的消息数，根据实际情况更改
)

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:        "127.0.0.1:6379",
		DB:          0,
		DialTimeout: time.Second * 60,
	})
	client2 := redis.NewClient(&redis.Options{
		Addr:        "127.0.0.1:6379",
		DB:          2,
		DialTimeout: time.Second * 60,
	})
	RDB0 = client
	RDB1 = client2
}

func ParseToken(token string, typ int64) (string, error) {
	item, err := rpc.ItemCenter.ParseToken(context.Background(), token, typ)
	if err != nil {
		fmt.Println("parse client error : ", err.Error())
		return "", err
	}
	return item, nil
}

func InRedisMq(u *itemcenter.User) error {
	err := RDB0.XGroupCreateMkStream(ctx, stream, group, "0-0").Err()
	if err != nil {
		fmt.Println("create mq error : ", err)
		return err
	}
	values := map[string]interface{}{"username": u.Username, "id": u.Id, "password": u.Password}
	err = RDB0.XAdd(ctx, &redis.XAddArgs{
		Stream: stream,
		Values: values,
	}).Err()
	return nil
}

func TokenBucket(num int64, c *gin.Context) {
	var n int64 = 0
	for n < num {
		messages, err := RDB0.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    group,
			Consumer: consumer,
			Streams:  []string{stream, start},
			Count:    int64(count),
			Block:    0,
			NoAck:    false,
		}).Result()
		if err != nil {
			fmt.Println("redisMq read error : ", err)
			return
		}
		limiter := rate.NewLimiter(rate.Limit(50), 100)
		for _, message := range messages {
			values := message.Messages
			for _, Value := range values {
				value := Value.Values
				if err = limiter.WaitN(ctx, 50); err != nil {
					fmt.Println("limiter error : ", err)
					return
				}
				var user itemcenter.User
				if username, ok := value["username"].(string); ok {
					user.Username = username
				}
				if password, ok := value["password"].(string); ok {
					user.Password = password
				}
				//if id, ok := value["id"].(int64); ok {
				//	user.Id = id
				//}
				skToken, err := rpc.ItemCenter.CreateToken(ctx, &user, model.SKToken)
				if err != nil {
					resps.InternalErr(c)
					c.Abort()
					return
				}
				if err := RDB1.Set(ctx, user.Username, skToken, time.Hour*24*7).Err(); err == nil {
					return
				}
			}
		}
		n++
	}
}
