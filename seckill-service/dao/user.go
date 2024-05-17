package dao

import (
	"context"
	"errors"
	"fmt"
	"go_code/seckill-service/model"
	"go_code/seckill-service/rpc/kitex_gen/itemcenter"
	"gorm.io/gorm"
	"time"
)

func FindUserByName(username string) (user *itemcenter.User, err error) {
	err = DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("user not found : ", err)
			return nil, err
		}
		fmt.Println("find user by username error :", err)
	}
	return user, err
}

func FindUser(username string, password string) (user *itemcenter.User, err error) {
	if err = DB.Where("username = ? and password =?", username, password).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("find user err : ", err)
	}
	return
}

func CreateUser(u model.User) (err error) {
	if err = DB.Create(&u).Error; err != nil {
		fmt.Println("crate user error :", err)
	}
	if err = RDB0.Set(context.Background(), u.Username, u.Password, time.Duration(time.Hour*24)).Err(); err != nil {
		fmt.Println("user into Redis error : ", err)
		return err
	}
	return nil
}
