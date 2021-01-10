package service

import (
	"app/database"
	"app/model"
	"fmt"
	"log"
)

type UserService struct{}

func (UserService) Create(user *model.User) error {
	err := user.SetPassword(user.Password)
	if err != nil {
		log.Println(err.Error())
	}
	db := database.GetDB()
	if err := db.Create(user); err != nil {
		return fmt.Errorf("failed to create user: %w", err.Error)
	}
	return nil
}
