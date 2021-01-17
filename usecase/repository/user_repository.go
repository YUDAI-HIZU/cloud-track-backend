package repository

import "app/domain"

type UserRepository interface {
	GetByID(int) (*domain.User, error)
	Create(user *domain.User) error
	GetByEmail(string) (*domain.User, error)
}
