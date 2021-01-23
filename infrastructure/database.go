package infrastructure

import (
	"app/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.DatabaseURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
