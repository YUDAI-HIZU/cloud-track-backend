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
	return &user, nil
}
