package dao

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"go_code/seckill-service/model"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	DB   *gorm.DB
	DB2  *gorm.DB
	RDB0 *redis.Client
)

func InitDB() {
	username := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	Dbname := "seckill"
	timeout := "10s"
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	dns2 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, "winner", timeout)
	db, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		log.Panicf("Connect DB error: %v  \n", err)
		return
	}
	db2, err := gorm.Open(mysql.Open(dns2))
	if err != nil {
		log.Panicf("Connect DB error: %v  \n", err)
		return
	}
	DB = db
	DB2 = db2
	err = DB.AutoMigrate(&model.User{}) //用model.User初始化数据库user表
	if err != nil {
		fmt.Println("init user table error :", err)
		return
	}
	err = DB2.AutoMigrate(&model.Winner{}) //用model.User初始化数据库user表
	if err != nil {
		fmt.Println("init winner table error :", err)
		return
	}
}

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:        "redis-14520.c299.asia-northeast1-1.gce.cloud.redislabs.com:14520",
		Password:    "rPYdtUeiD5CeJSqcGZdoyHDd6Ou2uApa",
		DB:          0,
		DialTimeout: time.Second * 5,
	})
	RDB0 = client
}
