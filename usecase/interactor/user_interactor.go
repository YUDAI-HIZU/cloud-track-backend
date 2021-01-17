package interactor

import (
	"app/domain"
	"app/usecase/repository"

	"golang.org/x/crypto/bcrypt"
)

func genEncryptedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func compareHashAndPassWord(encryptedPassword string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}

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
	encryptedPassword, err := genEncryptedPassword(user.Password)
	if err != nil {
		return err
	}
	user.EncryptedPassword = encryptedPassword
	err = u.UserRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserInteractor) SignIn(input *domain.SignInInput) (*domain.User, error) {
	user, err := u.UserRepository.GetByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if err := compareHashAndPassWord(user.EncryptedPassword, input.Password); err != nil {
		return nil, err
	}
	return user, nil
}
