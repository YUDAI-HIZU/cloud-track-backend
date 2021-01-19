package infrastructure

import (
	"app/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", config.DatabaseURL)
	if err != nil {
		panic(err)
	}
}

func NewDatabase() *gorm.DB {
	return db
}
