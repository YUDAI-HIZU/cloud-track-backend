package repository

import (
	"app/domain"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) GetByID(id int) (user domain.User, err error) {
	if err := u.db.Take(&user, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return user, err
		}
		return user, err
	}
	return user, nil
}
