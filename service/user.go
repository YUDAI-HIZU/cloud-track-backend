package service

import (
	"app/database"
	"app/model"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type UserService struct{}

func (UserService) Create(user *model.User) error {
	err := user.SetPassword(user.Password)
	if err != nil {
		log.Print(err.Error())
	}
	db := database.GetDB()
	if err := db.Create(user); err != nil {
		return fmt.Errorf("failed to create user: %w", err.Error)
	}
	return nil
}

func (UserService) GetByEmail(email string) (*model.User, error) {
	var user model.User
	db := database.GetDB()
	if err := db.Where("email = ?", email).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	return &user, nil
}

func (UserService) GetByID(id string) (*model.User, error) {
	var user model.User
	db := database.GetDB()
	if err := db.Where("id = ?", id).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	return &user, nil
}
