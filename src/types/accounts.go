package types

import (
	"github.com/adigunhammedolalekan/services/src/fn"
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	Email string
	Password string
}

func NewAccount(email, password string) (*Account, error) {
	if err := fn.ValidateEmail(email); err != nil {
		return nil, err
	}
	hashPassword := fn.HashPassword(password)
	return &Account{
		Email:    email,
		Password: hashPassword,
	}, nil
}
