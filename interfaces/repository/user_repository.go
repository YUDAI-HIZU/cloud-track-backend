package repository

import (
	"app/domain"
	"fmt"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (u *UserRepository) GetByID(id int) (*domain.User, error) {
	var user domain.User
	if err := u.DB.Take(&user, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = err
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) Create(user *domain.User) error {
	result := u.DB.Create(user)
	if err := result.Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (u *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = err
		}
		return nil, err
	}
	return &user, nil
}
