package database

import (
	"app/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB  *gorm.DB
	err error
)

func Conn() *gorm.DB {
	DB, err = gorm.Open("mysql", config.DatabaseURL)
	if err != nil {
		panic(err)
	}
	return DB
}
