package interactor

import (
	"app/domain"
	"app/usecase/repository"
)

type UserInteractor struct {
	UserRepository repository.UserRepository
}

func (u *UserInteractor) GetByID(id int) (user domain.User, err error) {
	user, err = u.UserRepository.GetByID(id)
	return user, nil
}
