package repository

import "app/domain"

type UserRepository interface {
	GetByID(int) (domain.User, error)
}
