package service

import (
	"go_code/seckill-service/dao"
	"go_code/seckill-service/rpc/kitex_gen/itemcenter"
)

func SecKill(u *itemcenter.User) (OK bool, err error) {
	dao.DB2.Create(&u)
	return true, nil
}
