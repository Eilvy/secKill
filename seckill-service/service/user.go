package service

import (
	"errors"
	"fmt"
	"go_code/seckill-service/dao"
	"go_code/seckill-service/model"
	"go_code/seckill-service/rpc/kitex_gen/itemcenter"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(u *itemcenter.User) (err error) {
	_, err = dao.FindUserByName(u.Username)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	//加密用户密码
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("password hash error : ", err.Error())
		return err
	}

	err = dao.CreateUser(model.User{
		Username: u.Username,
		Password: string(password),
	})
	return err
}
func Login(u *itemcenter.User) (user *itemcenter.User, err error) {
	u1, err := dao.FindUserByName(u.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u1.Password), []byte(u.Password)); err != nil {
		err = errors.New("password is wrong")
	}

	user = &itemcenter.User{
		Username: u1.Username,
		Password: u1.Password,
		Id:       u1.Id,
	}
	return user, err
}
