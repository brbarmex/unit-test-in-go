package repositories

import (
	"store/core/accounts/entities"
	databases "store/infra/database"
)

type AccountRepository struct{}

func (ar AccountRepository) Select(query string) (accounts []entities.Credential) {
	db := databases.Connect()
	defer databases.Disconect(db)
	accounts = []entities.Credential{}
	db.Raw(query).Scan(&accounts)
	return accounts
}

func (ar AccountRepository) Create(entitie *entities.Credential) {
	db := databases.Connect()
	defer databases.Disconect(db)
	db.Create(&entitie)
}

func (ar AccountRepository) TakeCredentialByEmail(email string) entities.Credential {
	db := databases.Connect()
	defer databases.Disconect(db)
	credential := entities.Credential{}
	db.Where(entities.Credential{Email: email}, "email").Find(&credential)
	return credential
}
