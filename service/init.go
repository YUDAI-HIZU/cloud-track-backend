package service

import (
	"app/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	db, err := gorm.Open("mysql", config.DatabaseURL)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	db.LogMode(true)
}
