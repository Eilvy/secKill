package service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go_code/seckill-service/rpc/kitex_gen/itemcenter"
	"strconv"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Typ      int64  `json:"typ"`
	jwt.StandardClaims
}

var (
	JwtKey = []byte("leiyv000")
	exp    time.Time
)

func CreateToken(user *itemcenter.User, typ int64) (token string, err error) {

	if typ == 1000 {
		exp = time.Now().Add(time.Hour * 1)
	} else if typ == 2000 {
		exp = time.Now().Add(time.Hour * 24 * 14)
	} else if typ == 3000 {
		exp = time.Now().Add(time.Hour * 24 * 7)
	}
	claim := jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
		"id":       user.Id,
		"typ":      typ,
		"exp":      exp.Unix(),
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err = tok.SignedString(JwtKey)
	if err != nil {
		fmt.Println("signed token error : ", err.Error())
		return
	}
	return token, err
}

func ParseToken(token string, typ int64) (string, error) {
	Token, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return JwtKey, nil
	})
	if err != nil {
		fmt.Println("parse token error : ", err)
		return "", nil
	}
	if claims, ok := Token.Claims.(*MyClaims); ok && Token.Valid { // 校验token
		if typ == 1000 || typ == 2000 {
			//_, err = dao.FindUserByName(claims.Username)
			//if errors.Is(err, gorm.ErrRecordNotFound) {
			//	err = errors.New("can't find user by username")
			//	return "", err
			//} else if err != nil {
			//	return "", errors.New("find user by username error")
			//}
			return claims.Username, nil
		} else if typ == 4000 {
			//_, err = dao.FindUserByName(claims.Username)
			//if errors.Is(err, gorm.ErrRecordNotFound) {
			//	err = errors.New("can't find user by username")
			//	return "", err
			//} else if err != nil {
			//	return "", errors.New("find user by username error")
			//}
			return claims.Password, nil
		} else if typ == 5000 {
			//_, err = dao.FindUserByName(claims.Username)
			//if errors.Is(err, gorm.ErrRecordNotFound) {
			//	err = errors.New("can't find user by username")
			//	return "", err
			//} else if err != nil {
			//	return "", errors.New("find user by username error")
			//}
			return claims.Id, nil
		} else if typ == 6000 {
			return strconv.FormatInt(claims.Typ, 10), nil
		}

	}
	return "", errors.New("invalid token")
}
