package interfaces

import "store/core/accounts/entities"

type Repository interface {
	Select(query string) (credentials []entities.Credential)
	Create(credential *entities.Credential)
	TakeCredentialByEmail(email string) entities.Credential
}
