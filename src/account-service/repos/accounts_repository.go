package repos

import (
	"context"
	"github.com/adigunhammedolalekan/services/src/types"
	"github.com/jinzhu/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db:db}
}

func (repo *AccountRepository) CreateAccount(email, password string) (*types.Account, error) {
	account, err := types.NewAccount(email, password)
	if err != nil {
		return nil, err
	}
	tx := repo.db.Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}
	if err := tx.Create(account).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

}
