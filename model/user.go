package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"type:varchar(100);unique_index;not null"`
	Password string `gorm:"size:255;not null"`
}
