package interfaces

import "store/core/accounts/entities"

type Repository interface {
	Select(query string) (credentials []entities.Credential)
	Create(credential *entities.Credential) error
	TakeCredentialByEmail(email string) entities.Credential
	SaveAccount(account *entities.Account) error
	CheckAccountAlreadyExist(email string, identification string) bool
}
