package database

import (
	"app/config"
	"app/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("mysql", config.DatabaseURL)
	db.AutoMigrate(&model.User{})

	if err != nil {
		panic(err.Error())
	}
	autoMigration()
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&model.User{})
}
