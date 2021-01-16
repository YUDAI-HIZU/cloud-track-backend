package repository

import (
	"app/domain"
	"errors"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (u *UserRepository) GetByID(id int) (*domain.User, error) {
	var user domain.User
	if err := u.DB.Take(&user, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = errors.New("not found")
		}
		return nil, err
	}
	return &user, nil
}
