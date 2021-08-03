package repositories

import (
	"store/core/accounts/entities"

	"gorm.io/gorm"
)

func SelectAllWithoutPagination(db *gorm.DB) (accounts []entities.Account) {
	accounts = []entities.Account{}
	db.Raw("SELECT * FROM ACCOUNTS").Scan(&accounts)
	return accounts
}
