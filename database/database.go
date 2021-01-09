package database

import (
	"app/config"
	"app/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GormConnect() {
	db, err := gorm.Open("mysql", config.DatabaseURL)
	db.AutoMigrate(&model.User{})

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	db.LogMode(true)
}
