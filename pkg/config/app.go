package config

import (
	"github.com/jinzhou/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "localhost:localhost/simplerest?charset=utf8&parseTime=tTrue&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
