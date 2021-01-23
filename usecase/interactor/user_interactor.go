package interactor

import (
	"app/config"
	"app/domain"
	"app/usecase/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserInteractor struct {
	UserRepository repository.UserRepository
}

func generateEncryptedPassword(password string) (string, error) {
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

func generateToken(id int) (string, error) {
	claims := &jwt.MapClaims{
		"exp":    time.Now().Add(24 * 7 * time.Hour).Unix(),
		"userID": id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	return tokenString, err
}

func (u *UserInteractor) GetByID(id int) (*domain.User, error) {
	user, err := u.UserRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserInteractor) Create(user *domain.User) error {
	encryptedPassword, err := generateEncryptedPassword(user.Password)
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

func (u *UserInteractor) SignIn(input *domain.SignInInput) (*domain.User, string, error) {
	user, err := u.UserRepository.GetByEmail(input.Email)
	if err != nil {
		return nil, "", err
	}
	if err := compareHashAndPassWord(user.EncryptedPassword, input.Password); err != nil {
		return nil, "", err
	}
	token, err := generateToken(user.ID)
	return user, token, nil
}
