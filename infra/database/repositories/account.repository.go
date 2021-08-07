package repositories

import (
	"store/core/accounts/entities"
	databases "store/infra/database"
)

type AccountRepository struct{}

func (ar AccountRepository) GetAllAccounts() (accounts []entities.Account) {
	db := databases.Connect()
	defer databases.Disconect(db)
	accounts = []entities.Account{}
	db.Raw("SELECT * FROM ACCOUNTS").Scan(&accounts)
	return accounts
}

func (ar AccountRepository) Select(query string) (accounts []entities.Credential) {
	db := databases.Connect()
	defer databases.Disconect(db)
	accounts = []entities.Credential{}
	db.Raw(query).Scan(&accounts)
	return accounts
}

func (ar AccountRepository) Create(entitie *entities.Credential) error {
	db := databases.Connect()
	defer databases.Disconect(db)

	if result := db.Create(&entitie); result.Error != nil {
		return result.Error
	}

	return nil
}

func (ar AccountRepository) TakeCredentialByEmail(email string) entities.Credential {
	db := databases.Connect()
	defer databases.Disconect(db)
	credential := entities.Credential{}
	db.Where(entities.Credential{Email: email}, "email").Find(&credential)
	return credential
}

func (ar AccountRepository) CheckAccountAlreadyExist(email string, identification string) bool {
	db := databases.Connect()
	defer databases.Disconect(db)
	account := entities.Account{}
	result := db.Where("email = ? OR identification = ?", email, identification).First(&account)
	return result.RowsAffected > 0
}

func (ar AccountRepository) SaveAccount(entitie *entities.Account) error {
	db := databases.Connect()
	defer databases.Disconect(db)

	if result := db.Create(&entitie); result.Error != nil {
		return result.Error
	}

	return nil
}
