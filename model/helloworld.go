package model

import "gorm.io/gorm"

type Helloworld struct {
	gorm.Model
	Msg string `column:"msg"`
}

func (u *Helloworld) TableName() string {
	return "helloworld"
}