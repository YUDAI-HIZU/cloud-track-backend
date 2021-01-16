package domain

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func (user *User) SetPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to set password: %w", err)
	}
	user.Password = string(hash)
	return nil
}

func (user *User) PasswordVerify(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
