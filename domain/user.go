package domain

import "time"

type User struct {
	ID                int
	Name              string
	Email             string
	Password          string `gorm:"-"`
	EncryptedPassword string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type SignInInput struct {
	Email    string
	Password string
}
