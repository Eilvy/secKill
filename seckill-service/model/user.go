package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
}
type Winner struct {
	gorm.Model
	Username string
}
