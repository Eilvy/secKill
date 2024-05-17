package model

type Register struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type Login struct {
	Id       int64  `form:"id"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
