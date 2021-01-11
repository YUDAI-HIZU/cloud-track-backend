package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string `binding:"required"`
	Email    string `gorm:"unique_index" binding:"required,email"`
	Password string `binding:"required,min=8"`
}

type UserInfo struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required,min=8"`
}

func (user *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to set password: %w", err)
	}
	user.Password = string(hash)
	return nil
}

func (user *User) PasswordVerify(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
