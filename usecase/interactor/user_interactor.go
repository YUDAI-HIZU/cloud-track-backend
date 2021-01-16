package interactor

import (
	"app/domain"
	"app/usecase/repository"
)

type UserInteractor struct {
	UserRepository repository.UserRepository
}

func (u *UserInteractor) GetByID(id int) (*domain.User, error) {
	user, err := u.UserRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserInteractor) Create(user *domain.User) error {
	err := u.UserRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}
