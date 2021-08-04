package repositories

import (
	"store/core/accounts/entities"

	"gorm.io/gorm"
)

type AccountRepository struct {
	Database *gorm.DB
}

func (repo AccountRepository) SelectAllWithoutPagination() (accounts []entities.Account) {
	accounts = []entities.Account{}
	repo.Database.Raw("SELECT * FROM ACCOUNTS").Scan(&accounts)
	return accounts
}

func (repo AccountRepository) Insert(entitie *entities.Account) {
	repo.Database.Create(&entitie)
}
